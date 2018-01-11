package core

import (
	"errors"
	"fmt"
	"path"
	"strings"
	"time"

	log "github.com/Sirupsen/logrus"
	e "github.com/pkg/errors"
	ie "github.com/sahib/brig/catfs/errors"
	n "github.com/sahib/brig/catfs/nodes"
	h "github.com/sahib/brig/util/hashlib"
)

var (
	ErrIsGhost = errors.New("Is a ghost")
)

// mkdirParents takes the dirname of repoPath and makes sure all intermediate
// directories are created. The last directory will be returned.

// If any directory exist already, it will not be touched.
// You can also think of it as mkdir -p.
func mkdirParents(lkr *Linker, repoPath string) (*n.Directory, error) {
	repoPath = path.Clean(repoPath)

	elems := strings.Split(repoPath, "/")
	for idx := 0; idx < len(elems)-1; idx++ {
		dirname := strings.Join(elems[:idx+1], "/")
		if dirname == "" {
			dirname = "/"
		}

		dir, err := Mkdir(lkr, dirname, false)
		if err != nil {
			return nil, err
		}

		// Return it, if it's the last path component:
		if idx+1 == len(elems)-1 {
			return dir, nil
		}
	}

	return nil, fmt.Errorf("Empty path given")
}

// Mkdir creates the directory at repoPath and any intermediate directories if
// createParents is true. It will fail if there is already a file at `repoPath`
// and it is not a directory.
func Mkdir(lkr *Linker, repoPath string, createParents bool) (dir *n.Directory, err error) {
	dirname, basename := path.Split(repoPath)

	// Take special care of the root node:
	if basename == "" {
		return lkr.Root()
	}

	// Check if the parent exists:
	parent, err := lkr.LookupDirectory(dirname)
	if err != nil && !ie.IsNoSuchFileError(err) {
		return nil, e.Wrap(err, "dirname lookup failed")
	}

	log.Debugf("mkdir: %s", dirname)

	// If it's nil, we might need to create it:
	if parent == nil {
		if !createParents {
			return nil, ie.NoSuchFile(dirname)
		}

		parent, err = mkdirParents(lkr, repoPath)
		if err != nil {
			return nil, err
		}
	}

	child, err := parent.Child(lkr, basename)
	if err != nil {
		return nil, err
	}

	if child != nil {
		if child.Type() != n.NodeTypeDirectory {
			return nil, fmt.Errorf("`%s` exists and is not a directory", repoPath)
		}

		// Nothing to do really. Return the old child.
		return child.(*n.Directory), nil
	}

	// Make sure, NextInode() and StageNode is written in one batch.
	batch := lkr.kv.Batch()
	defer func() {
		if err != nil {
			batch.Rollback()
		} else {
			err = batch.Flush()
		}
	}()

	// Create it then!
	dir, err = n.NewEmptyDirectory(lkr, parent, basename, lkr.NextInode())
	if err != nil {
		return nil, err
	}

	if err := lkr.StageNode(dir); err != nil {
		return nil, err
	}

	return dir, nil
}

// Remove removes a single node from a directory.
// `nd` is the node that shall be removed and may not be root.
// The parent directory is returned.
func Remove(lkr *Linker, nd n.ModNode, createGhost, force bool) (parentDir *n.Directory, ghost *n.Ghost, err error) {
	if !force && nd.Type() == n.NodeTypeGhost {
		return nil, nil, ErrIsGhost
	}

	parentDir, err = n.ParentDirectory(lkr, nd)
	if err != nil {
		return nil, nil, err
	}

	// We shouldn't delete the root directory
	// (only directory with a parent)
	if parentDir == nil {
		return nil, nil, fmt.Errorf("Refusing to delete /")
	}

	if err := parentDir.RemoveChild(lkr, nd); err != nil {
		return nil, nil, fmt.Errorf("Failed to remove child: %v", err)
	}

	lkr.MemIndexPurge(nd)

	// Make sure both StageNode() will be written in one batch:
	batch := lkr.kv.Batch()
	defer func() {
		if err != nil {
			batch.Rollback()
		} else {
			err = batch.Flush()
		}
	}()

	if err := lkr.StageNode(parentDir); err != nil {
		return nil, nil, err
	}

	if createGhost {
		ghost, err := n.MakeGhost(nd, lkr.NextInode())
		if err != nil {
			return nil, nil, err
		}

		parentDir.Add(lkr, ghost)
		if err != nil {
			return nil, nil, err
		}

		if err := lkr.StageNode(ghost); err != nil {
			return nil, nil, err
		}

		return parentDir, ghost, nil
	}

	return parentDir, nil, nil
}

func Move(lkr *Linker, nd n.ModNode, destPath string) (err error) {
	// Forbid moving a node inside of one of it's subdirectories.
	if nd.Path() == destPath {
		return fmt.Errorf("Source and Dest are the same file: %v", destPath)
	}

	if strings.HasPrefix(destPath, nd.Path()) {
		return fmt.Errorf(
			"Cannot move `%s` into it's own subdir `%s`",
			nd.Path(),
			destPath,
		)
	}

	// Check if the destination already exists:
	destNode, err := lkr.LookupModNode(destPath)
	if err != nil && !ie.IsNoSuchFileError(err) {
		return err
	}

	var parentDir *n.Directory

	// Make sure both StageNode() will be written in one batch:
	batch := lkr.kv.Batch()
	defer func() {
		if err != nil {
			batch.Rollback()
		} else {
			err = batch.Flush()
		}
	}()

	if destNode != nil {
		switch destNode.Type() {
		case n.NodeTypeDirectory:
			// Move inside of this directory.
			// Check if there is already a file
			destDir, ok := destNode.(*n.Directory)
			if !ok {
				return ie.ErrBadNode
			}

			child, err := destDir.Child(lkr, nd.Name())
			if err != nil {
				return err
			}

			// Oh, something is in there?
			if child != nil {
				if nd.Type() == n.NodeTypeFile {
					// TODO: more details
					return fmt.Errorf("Cannot overwrite a directory with a file")
				}

				childDir, ok := child.(*n.Directory)
				if !ok {
					return ie.ErrBadNode
				}

				if childDir.Size() > 0 {
					return fmt.Errorf("Cannot move over: %s; directory is not empty!", child.Path())
				}

				// Okay, there is an empty directory. Let's remove it to
				// replace it with our source node.
				log.Warningf("Remove child dir: %v", childDir)
				if _, _, err := Remove(lkr, childDir, false, false); err != nil {
					return err
				}
			}

			parentDir = destDir
		case n.NodeTypeFile:
			log.Warningf("Remove file: %v", destNode.Path())
			parentDir, _, err = Remove(lkr, destNode, false, false)
			if err != nil {
				return err
			}
		case n.NodeTypeGhost:
			// It is already a ghost. Overwrite it and do not create a new one.
			log.Warningf("Remove ghost: %v", destNode.Path())
			parentDir, _, err = Remove(lkr, destNode, false, true)
			if err != nil {
				return err
			}
		default:
			return ie.ErrBadNode
		}
	} else {
		// No node at this place yet, attempt to look it up.
		parentDir, err = lkr.LookupDirectory(path.Dir(destPath))
		if err != nil {
			return err
		}
	}

	oldPath := nd.Path()

	// Remove the old node:
	_, ghost, err := Remove(lkr, nd, true, true)
	if err != nil {
		return err
	}

	// The node needs to be told that it's path changed,
	// since it might need to change it's hash value now.
	if err := nd.NotifyMove(lkr, oldPath, destPath); err != nil {
		return err
	}

	// And add it to the right destination dir:
	if err := parentDir.Add(lkr, nd); err != nil {
		return err
	}

	err = n.Walk(lkr, nd, true, func(child n.Node) error {
		return lkr.StageNode(child)
	})

	if err != nil {
		return err
	}

	if err := lkr.AddMoveMapping(nd, ghost); err != nil {
		return err
	}

	err = lkr.StageNode(nd)
	if err != nil {
		return err
	}

	return err
}

// TODO: This interface sucks.
type NodeUpdate struct {
	Hash   h.Hash
	Size   uint64
	Author string
	Key    []byte
}

func Stage(lkr *Linker, repoPath string, info *NodeUpdate) (file *n.File, err error) {
	node, err := lkr.LookupNode(repoPath)
	if err != nil && !ie.IsNoSuchFileError(err) {
		return nil, err
	}

	batch := lkr.kv.Batch()
	defer func() {
		if err != nil {
			batch.Rollback()
		} else {
			err = batch.Flush()
		}
	}()

	if node != nil && node.Type() == n.NodeTypeGhost {
		ghostParent, err := n.ParentDirectory(lkr, node)
		if err != nil {
			return nil, err
		}

		if ghostParent == nil {
			// TODO: Think about this case. stage() should not be called on directories
			//       anyways (only on files or files to ghosts)
			return nil, fmt.Errorf("The ghost is a root? Something is wrong...")
		}

		if err := ghostParent.RemoveChild(lkr, node); err != nil {
			return nil, err
		}

		// Act like there was no previous node.
		// New node will have a different Inode.
		file = nil
	} else if node != nil {
		var ok bool
		file, ok = node.(*n.File)
		if !ok {
			return nil, ie.ErrBadNode
		}
	}

	needRemove := false

	if file != nil {
		// We know this file already.
		log.WithFields(log.Fields{"file": repoPath}).Info("File exists; modifying.")
		needRemove = true

		if file.Content().Equal(info.Hash) {
			log.Debugf("Hash was not modified. Not doing any update.")
			return file, nil
		}
	} else {
		parent, err := mkdirParents(lkr, repoPath)
		if err != nil {
			return nil, err
		}

		// Create a new file at specified path:
		file, err = n.NewEmptyFile(parent, path.Base(repoPath), lkr.NextInode())
		if err != nil {
			return nil, err
		}
	}

	parentDir, err := n.ParentDirectory(lkr, file)
	if err != nil {
		return nil, err
	}

	if parentDir == nil {
		return nil, fmt.Errorf("%s has no parent yet (BUG)", repoPath)
	}

	if needRemove {
		// Remove the child before changing the hash:
		if err := parentDir.RemoveChild(lkr, file); err != nil {
			return nil, err
		}
	}

	file.SetSize(info.Size)
	file.SetModTime(time.Now())
	file.SetContent(lkr, info.Hash)
	file.SetKey(info.Key)

	// Add it again when the hash was changed.
	if err := parentDir.Add(lkr, file); err != nil {
		return nil, err
	}

	if err := lkr.StageNode(file); err != nil {
		return nil, err
	}

	return file, nil
}

func Log(lkr *Linker, fn func(cmt *n.Commit) error) error {
	curr, err := lkr.Status()
	if err != nil {
		return err
	}

	for curr != nil {
		if err := fn(curr); err != nil {
			return err
		}

		parent, err := curr.Parent(lkr)
		if err != nil {
			return err
		}

		if parent == nil {
			break
		}

		parentCmt, ok := parent.(*n.Commit)
		if !ok {
			return ie.ErrBadNode
		}

		curr = parentCmt
	}

	return nil
}