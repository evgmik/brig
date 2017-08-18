// Code generated by capnpc-go. DO NOT EDIT.

package capnp

import (
	strconv "strconv"
	capnp "zombiezen.com/go/capnproto2"
	text "zombiezen.com/go/capnproto2/encoding/text"
	schemas "zombiezen.com/go/capnproto2/schemas"
)

// Person might be any brig user
type Person struct{ capnp.Struct }

// Person_TypeID is the unique identifier for the type Person.
const Person_TypeID = 0xf736dd278ea58545

func NewPerson(s *capnp.Segment) (Person, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return Person{st}, err
}

func NewRootPerson(s *capnp.Segment) (Person, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return Person{st}, err
}

func ReadRootPerson(msg *capnp.Message) (Person, error) {
	root, err := msg.RootPtr()
	return Person{root.Struct()}, err
}

func (s Person) String() string {
	str, _ := text.Marshal(0xf736dd278ea58545, s.Struct)
	return str
}

func (s Person) Ident() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s Person) HasIdent() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Person) IdentBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s Person) SetIdent(v string) error {
	return s.Struct.SetText(0, v)
}

func (s Person) Hash() ([]byte, error) {
	p, err := s.Struct.Ptr(1)
	return []byte(p.Data()), err
}

func (s Person) HasHash() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s Person) SetHash(v []byte) error {
	return s.Struct.SetData(1, v)
}

// Person_List is a list of Person.
type Person_List struct{ capnp.List }

// NewPerson creates a new list of Person.
func NewPerson_List(s *capnp.Segment, sz int32) (Person_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2}, sz)
	return Person_List{l}, err
}

func (s Person_List) At(i int) Person { return Person{s.List.Struct(i)} }

func (s Person_List) Set(i int, v Person) error { return s.List.SetStruct(i, v.Struct) }

func (s Person_List) String() string {
	str, _ := text.MarshalList(0xf736dd278ea58545, s.List)
	return str
}

// Person_Promise is a wrapper for a Person promised by a client call.
type Person_Promise struct{ *capnp.Pipeline }

func (p Person_Promise) Struct() (Person, error) {
	s, err := p.Pipeline.Struct()
	return Person{s}, err
}

// Commit is a set of changes to nodes
type Commit struct{ capnp.Struct }
type Commit_merge Commit

// Commit_TypeID is the unique identifier for the type Commit.
const Commit_TypeID = 0x8da013c66e545daf

func NewCommit(s *capnp.Segment) (Commit, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 6})
	return Commit{st}, err
}

func NewRootCommit(s *capnp.Segment) (Commit, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 6})
	return Commit{st}, err
}

func ReadRootCommit(msg *capnp.Message) (Commit, error) {
	root, err := msg.RootPtr()
	return Commit{root.Struct()}, err
}

func (s Commit) String() string {
	str, _ := text.Marshal(0x8da013c66e545daf, s.Struct)
	return str
}

func (s Commit) Message() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s Commit) HasMessage() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Commit) MessageBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s Commit) SetMessage(v string) error {
	return s.Struct.SetText(0, v)
}

func (s Commit) Author() (Person, error) {
	p, err := s.Struct.Ptr(1)
	return Person{Struct: p.Struct()}, err
}

func (s Commit) HasAuthor() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s Commit) SetAuthor(v Person) error {
	return s.Struct.SetPtr(1, v.Struct.ToPtr())
}

// NewAuthor sets the author field to a newly
// allocated Person struct, preferring placement in s's segment.
func (s Commit) NewAuthor() (Person, error) {
	ss, err := NewPerson(s.Struct.Segment())
	if err != nil {
		return Person{}, err
	}
	err = s.Struct.SetPtr(1, ss.Struct.ToPtr())
	return ss, err
}

func (s Commit) Parent() ([]byte, error) {
	p, err := s.Struct.Ptr(2)
	return []byte(p.Data()), err
}

func (s Commit) HasParent() bool {
	p, err := s.Struct.Ptr(2)
	return p.IsValid() || err != nil
}

func (s Commit) SetParent(v []byte) error {
	return s.Struct.SetData(2, v)
}

func (s Commit) Root() ([]byte, error) {
	p, err := s.Struct.Ptr(3)
	return []byte(p.Data()), err
}

func (s Commit) HasRoot() bool {
	p, err := s.Struct.Ptr(3)
	return p.IsValid() || err != nil
}

func (s Commit) SetRoot(v []byte) error {
	return s.Struct.SetData(3, v)
}

func (s Commit) Merge() Commit_merge { return Commit_merge(s) }

func (s Commit_merge) IsMerge() bool {
	return s.Struct.Bit(0)
}

func (s Commit_merge) SetIsMerge(v bool) {
	s.Struct.SetBit(0, v)
}

func (s Commit_merge) With() (Person, error) {
	p, err := s.Struct.Ptr(4)
	return Person{Struct: p.Struct()}, err
}

func (s Commit_merge) HasWith() bool {
	p, err := s.Struct.Ptr(4)
	return p.IsValid() || err != nil
}

func (s Commit_merge) SetWith(v Person) error {
	return s.Struct.SetPtr(4, v.Struct.ToPtr())
}

// NewWith sets the with field to a newly
// allocated Person struct, preferring placement in s's segment.
func (s Commit_merge) NewWith() (Person, error) {
	ss, err := NewPerson(s.Struct.Segment())
	if err != nil {
		return Person{}, err
	}
	err = s.Struct.SetPtr(4, ss.Struct.ToPtr())
	return ss, err
}

func (s Commit_merge) Hash() ([]byte, error) {
	p, err := s.Struct.Ptr(5)
	return []byte(p.Data()), err
}

func (s Commit_merge) HasHash() bool {
	p, err := s.Struct.Ptr(5)
	return p.IsValid() || err != nil
}

func (s Commit_merge) SetHash(v []byte) error {
	return s.Struct.SetData(5, v)
}

// Commit_List is a list of Commit.
type Commit_List struct{ capnp.List }

// NewCommit creates a new list of Commit.
func NewCommit_List(s *capnp.Segment, sz int32) (Commit_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 6}, sz)
	return Commit_List{l}, err
}

func (s Commit_List) At(i int) Commit { return Commit{s.List.Struct(i)} }

func (s Commit_List) Set(i int, v Commit) error { return s.List.SetStruct(i, v.Struct) }

func (s Commit_List) String() string {
	str, _ := text.MarshalList(0x8da013c66e545daf, s.List)
	return str
}

// Commit_Promise is a wrapper for a Commit promised by a client call.
type Commit_Promise struct{ *capnp.Pipeline }

func (p Commit_Promise) Struct() (Commit, error) {
	s, err := p.Pipeline.Struct()
	return Commit{s}, err
}

func (p Commit_Promise) Author() Person_Promise {
	return Person_Promise{Pipeline: p.Pipeline.GetPipeline(1)}
}

func (p Commit_Promise) Merge() Commit_merge_Promise { return Commit_merge_Promise{p.Pipeline} }

// Commit_merge_Promise is a wrapper for a Commit_merge promised by a client call.
type Commit_merge_Promise struct{ *capnp.Pipeline }

func (p Commit_merge_Promise) Struct() (Commit_merge, error) {
	s, err := p.Pipeline.Struct()
	return Commit_merge{s}, err
}

func (p Commit_merge_Promise) With() Person_Promise {
	return Person_Promise{Pipeline: p.Pipeline.GetPipeline(4)}
}

// A single directory entry
type DirEntry struct{ capnp.Struct }

// DirEntry_TypeID is the unique identifier for the type DirEntry.
const DirEntry_TypeID = 0x8b15ee76774b1f9d

func NewDirEntry(s *capnp.Segment) (DirEntry, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return DirEntry{st}, err
}

func NewRootDirEntry(s *capnp.Segment) (DirEntry, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return DirEntry{st}, err
}

func ReadRootDirEntry(msg *capnp.Message) (DirEntry, error) {
	root, err := msg.RootPtr()
	return DirEntry{root.Struct()}, err
}

func (s DirEntry) String() string {
	str, _ := text.Marshal(0x8b15ee76774b1f9d, s.Struct)
	return str
}

func (s DirEntry) Name() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s DirEntry) HasName() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s DirEntry) NameBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s DirEntry) SetName(v string) error {
	return s.Struct.SetText(0, v)
}

func (s DirEntry) Hash() ([]byte, error) {
	p, err := s.Struct.Ptr(1)
	return []byte(p.Data()), err
}

func (s DirEntry) HasHash() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s DirEntry) SetHash(v []byte) error {
	return s.Struct.SetData(1, v)
}

// DirEntry_List is a list of DirEntry.
type DirEntry_List struct{ capnp.List }

// NewDirEntry creates a new list of DirEntry.
func NewDirEntry_List(s *capnp.Segment, sz int32) (DirEntry_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2}, sz)
	return DirEntry_List{l}, err
}

func (s DirEntry_List) At(i int) DirEntry { return DirEntry{s.List.Struct(i)} }

func (s DirEntry_List) Set(i int, v DirEntry) error { return s.List.SetStruct(i, v.Struct) }

func (s DirEntry_List) String() string {
	str, _ := text.MarshalList(0x8b15ee76774b1f9d, s.List)
	return str
}

// DirEntry_Promise is a wrapper for a DirEntry promised by a client call.
type DirEntry_Promise struct{ *capnp.Pipeline }

func (p DirEntry_Promise) Struct() (DirEntry, error) {
	s, err := p.Pipeline.Struct()
	return DirEntry{s}, err
}

// Directory contains one or more directories or files
type Directory struct{ capnp.Struct }

// Directory_TypeID is the unique identifier for the type Directory.
const Directory_TypeID = 0xe24c59306c829c01

func NewDirectory(s *capnp.Segment) (Directory, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 3})
	return Directory{st}, err
}

func NewRootDirectory(s *capnp.Segment) (Directory, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 3})
	return Directory{st}, err
}

func ReadRootDirectory(msg *capnp.Message) (Directory, error) {
	root, err := msg.RootPtr()
	return Directory{root.Struct()}, err
}

func (s Directory) String() string {
	str, _ := text.Marshal(0xe24c59306c829c01, s.Struct)
	return str
}

func (s Directory) Size() uint64 {
	return s.Struct.Uint64(0)
}

func (s Directory) SetSize(v uint64) {
	s.Struct.SetUint64(0, v)
}

func (s Directory) Parent() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s Directory) HasParent() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Directory) ParentBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s Directory) SetParent(v string) error {
	return s.Struct.SetText(0, v)
}

func (s Directory) Children() (DirEntry_List, error) {
	p, err := s.Struct.Ptr(1)
	return DirEntry_List{List: p.List()}, err
}

func (s Directory) HasChildren() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s Directory) SetChildren(v DirEntry_List) error {
	return s.Struct.SetPtr(1, v.List.ToPtr())
}

// NewChildren sets the children field to a newly
// allocated DirEntry_List, preferring placement in s's segment.
func (s Directory) NewChildren(n int32) (DirEntry_List, error) {
	l, err := NewDirEntry_List(s.Struct.Segment(), n)
	if err != nil {
		return DirEntry_List{}, err
	}
	err = s.Struct.SetPtr(1, l.List.ToPtr())
	return l, err
}

func (s Directory) GhostRef() ([]byte, error) {
	p, err := s.Struct.Ptr(2)
	return []byte(p.Data()), err
}

func (s Directory) HasGhostRef() bool {
	p, err := s.Struct.Ptr(2)
	return p.IsValid() || err != nil
}

func (s Directory) SetGhostRef(v []byte) error {
	return s.Struct.SetData(2, v)
}

// Directory_List is a list of Directory.
type Directory_List struct{ capnp.List }

// NewDirectory creates a new list of Directory.
func NewDirectory_List(s *capnp.Segment, sz int32) (Directory_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 3}, sz)
	return Directory_List{l}, err
}

func (s Directory_List) At(i int) Directory { return Directory{s.List.Struct(i)} }

func (s Directory_List) Set(i int, v Directory) error { return s.List.SetStruct(i, v.Struct) }

func (s Directory_List) String() string {
	str, _ := text.MarshalList(0xe24c59306c829c01, s.List)
	return str
}

// Directory_Promise is a wrapper for a Directory promised by a client call.
type Directory_Promise struct{ *capnp.Pipeline }

func (p Directory_Promise) Struct() (Directory, error) {
	s, err := p.Pipeline.Struct()
	return Directory{s}, err
}

// A leaf node in the MDAG
type File struct{ capnp.Struct }

// File_TypeID is the unique identifier for the type File.
const File_TypeID = 0x8ea7393d37893155

func NewFile(s *capnp.Segment) (File, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 4})
	return File{st}, err
}

func NewRootFile(s *capnp.Segment) (File, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 4})
	return File{st}, err
}

func ReadRootFile(msg *capnp.Message) (File, error) {
	root, err := msg.RootPtr()
	return File{root.Struct()}, err
}

func (s File) String() string {
	str, _ := text.Marshal(0x8ea7393d37893155, s.Struct)
	return str
}

func (s File) Size() uint64 {
	return s.Struct.Uint64(0)
}

func (s File) SetSize(v uint64) {
	s.Struct.SetUint64(0, v)
}

func (s File) Parent() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s File) HasParent() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s File) ParentBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s File) SetParent(v string) error {
	return s.Struct.SetText(0, v)
}

func (s File) Key() ([]byte, error) {
	p, err := s.Struct.Ptr(1)
	return []byte(p.Data()), err
}

func (s File) HasKey() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s File) SetKey(v []byte) error {
	return s.Struct.SetData(1, v)
}

func (s File) GhostRef() ([]byte, error) {
	p, err := s.Struct.Ptr(2)
	return []byte(p.Data()), err
}

func (s File) HasGhostRef() bool {
	p, err := s.Struct.Ptr(2)
	return p.IsValid() || err != nil
}

func (s File) SetGhostRef(v []byte) error {
	return s.Struct.SetData(2, v)
}

func (s File) Content() ([]byte, error) {
	p, err := s.Struct.Ptr(3)
	return []byte(p.Data()), err
}

func (s File) HasContent() bool {
	p, err := s.Struct.Ptr(3)
	return p.IsValid() || err != nil
}

func (s File) SetContent(v []byte) error {
	return s.Struct.SetData(3, v)
}

// File_List is a list of File.
type File_List struct{ capnp.List }

// NewFile creates a new list of File.
func NewFile_List(s *capnp.Segment, sz int32) (File_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 4}, sz)
	return File_List{l}, err
}

func (s File_List) At(i int) File { return File{s.List.Struct(i)} }

func (s File_List) Set(i int, v File) error { return s.List.SetStruct(i, v.Struct) }

func (s File_List) String() string {
	str, _ := text.MarshalList(0x8ea7393d37893155, s.List)
	return str
}

// File_Promise is a wrapper for a File promised by a client call.
type File_Promise struct{ *capnp.Pipeline }

func (p File_Promise) Struct() (File, error) {
	s, err := p.Pipeline.Struct()
	return File{s}, err
}

// Ghost indicates that a certain node was at this path once
type Ghost struct{ capnp.Struct }
type Ghost_Which uint16

const (
	Ghost_Which_commit    Ghost_Which = 0
	Ghost_Which_directory Ghost_Which = 1
	Ghost_Which_file      Ghost_Which = 2
)

func (w Ghost_Which) String() string {
	const s = "commitdirectoryfile"
	switch w {
	case Ghost_Which_commit:
		return s[0:6]
	case Ghost_Which_directory:
		return s[6:15]
	case Ghost_Which_file:
		return s[15:19]

	}
	return "Ghost_Which(" + strconv.FormatUint(uint64(w), 10) + ")"
}

// Ghost_TypeID is the unique identifier for the type Ghost.
const Ghost_TypeID = 0x80c828d7e89c12ea

func NewGhost(s *capnp.Segment) (Ghost, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	return Ghost{st}, err
}

func NewRootGhost(s *capnp.Segment) (Ghost, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	return Ghost{st}, err
}

func ReadRootGhost(msg *capnp.Message) (Ghost, error) {
	root, err := msg.RootPtr()
	return Ghost{root.Struct()}, err
}

func (s Ghost) String() string {
	str, _ := text.Marshal(0x80c828d7e89c12ea, s.Struct)
	return str
}

func (s Ghost) Which() Ghost_Which {
	return Ghost_Which(s.Struct.Uint16(0))
}
func (s Ghost) Commit() (Commit, error) {
	p, err := s.Struct.Ptr(0)
	return Commit{Struct: p.Struct()}, err
}

func (s Ghost) HasCommit() bool {
	if s.Struct.Uint16(0) != 0 {
		return false
	}
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Ghost) SetCommit(v Commit) error {
	s.Struct.SetUint16(0, 0)
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewCommit sets the commit field to a newly
// allocated Commit struct, preferring placement in s's segment.
func (s Ghost) NewCommit() (Commit, error) {
	s.Struct.SetUint16(0, 0)
	ss, err := NewCommit(s.Struct.Segment())
	if err != nil {
		return Commit{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

func (s Ghost) Directory() (Directory, error) {
	p, err := s.Struct.Ptr(0)
	return Directory{Struct: p.Struct()}, err
}

func (s Ghost) HasDirectory() bool {
	if s.Struct.Uint16(0) != 1 {
		return false
	}
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Ghost) SetDirectory(v Directory) error {
	s.Struct.SetUint16(0, 1)
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewDirectory sets the directory field to a newly
// allocated Directory struct, preferring placement in s's segment.
func (s Ghost) NewDirectory() (Directory, error) {
	s.Struct.SetUint16(0, 1)
	ss, err := NewDirectory(s.Struct.Segment())
	if err != nil {
		return Directory{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

func (s Ghost) File() (File, error) {
	p, err := s.Struct.Ptr(0)
	return File{Struct: p.Struct()}, err
}

func (s Ghost) HasFile() bool {
	if s.Struct.Uint16(0) != 2 {
		return false
	}
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Ghost) SetFile(v File) error {
	s.Struct.SetUint16(0, 2)
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewFile sets the file field to a newly
// allocated File struct, preferring placement in s's segment.
func (s Ghost) NewFile() (File, error) {
	s.Struct.SetUint16(0, 2)
	ss, err := NewFile(s.Struct.Segment())
	if err != nil {
		return File{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

// Ghost_List is a list of Ghost.
type Ghost_List struct{ capnp.List }

// NewGhost creates a new list of Ghost.
func NewGhost_List(s *capnp.Segment, sz int32) (Ghost_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1}, sz)
	return Ghost_List{l}, err
}

func (s Ghost_List) At(i int) Ghost { return Ghost{s.List.Struct(i)} }

func (s Ghost_List) Set(i int, v Ghost) error { return s.List.SetStruct(i, v.Struct) }

func (s Ghost_List) String() string {
	str, _ := text.MarshalList(0x80c828d7e89c12ea, s.List)
	return str
}

// Ghost_Promise is a wrapper for a Ghost promised by a client call.
type Ghost_Promise struct{ *capnp.Pipeline }

func (p Ghost_Promise) Struct() (Ghost, error) {
	s, err := p.Pipeline.Struct()
	return Ghost{s}, err
}

func (p Ghost_Promise) Commit() Commit_Promise {
	return Commit_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

func (p Ghost_Promise) Directory() Directory_Promise {
	return Directory_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

func (p Ghost_Promise) File() File_Promise {
	return File_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

// Node is a node in the merkle dag of brig
type Node struct{ capnp.Struct }
type Node_Which uint16

const (
	Node_Which_commit    Node_Which = 0
	Node_Which_directory Node_Which = 1
	Node_Which_file      Node_Which = 2
	Node_Which_ghost     Node_Which = 3
)

func (w Node_Which) String() string {
	const s = "commitdirectoryfileghost"
	switch w {
	case Node_Which_commit:
		return s[0:6]
	case Node_Which_directory:
		return s[6:15]
	case Node_Which_file:
		return s[15:19]
	case Node_Which_ghost:
		return s[19:24]

	}
	return "Node_Which(" + strconv.FormatUint(uint64(w), 10) + ")"
}

// Node_TypeID is the unique identifier for the type Node.
const Node_TypeID = 0xa629eb7f7066fae3

func NewNode(s *capnp.Segment) (Node, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 16, PointerCount: 4})
	return Node{st}, err
}

func NewRootNode(s *capnp.Segment) (Node, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 16, PointerCount: 4})
	return Node{st}, err
}

func ReadRootNode(msg *capnp.Message) (Node, error) {
	root, err := msg.RootPtr()
	return Node{root.Struct()}, err
}

func (s Node) String() string {
	str, _ := text.Marshal(0xa629eb7f7066fae3, s.Struct)
	return str
}

func (s Node) Which() Node_Which {
	return Node_Which(s.Struct.Uint16(8))
}
func (s Node) Name() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s Node) HasName() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Node) NameBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s Node) SetName(v string) error {
	return s.Struct.SetText(0, v)
}

func (s Node) Hash() ([]byte, error) {
	p, err := s.Struct.Ptr(1)
	return []byte(p.Data()), err
}

func (s Node) HasHash() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s Node) SetHash(v []byte) error {
	return s.Struct.SetData(1, v)
}

func (s Node) ModTime() (string, error) {
	p, err := s.Struct.Ptr(2)
	return p.Text(), err
}

func (s Node) HasModTime() bool {
	p, err := s.Struct.Ptr(2)
	return p.IsValid() || err != nil
}

func (s Node) ModTimeBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(2)
	return p.TextBytes(), err
}

func (s Node) SetModTime(v string) error {
	return s.Struct.SetText(2, v)
}

func (s Node) Inode() uint64 {
	return s.Struct.Uint64(0)
}

func (s Node) SetInode(v uint64) {
	s.Struct.SetUint64(0, v)
}

func (s Node) Commit() (Commit, error) {
	p, err := s.Struct.Ptr(3)
	return Commit{Struct: p.Struct()}, err
}

func (s Node) HasCommit() bool {
	if s.Struct.Uint16(8) != 0 {
		return false
	}
	p, err := s.Struct.Ptr(3)
	return p.IsValid() || err != nil
}

func (s Node) SetCommit(v Commit) error {
	s.Struct.SetUint16(8, 0)
	return s.Struct.SetPtr(3, v.Struct.ToPtr())
}

// NewCommit sets the commit field to a newly
// allocated Commit struct, preferring placement in s's segment.
func (s Node) NewCommit() (Commit, error) {
	s.Struct.SetUint16(8, 0)
	ss, err := NewCommit(s.Struct.Segment())
	if err != nil {
		return Commit{}, err
	}
	err = s.Struct.SetPtr(3, ss.Struct.ToPtr())
	return ss, err
}

func (s Node) Directory() (Directory, error) {
	p, err := s.Struct.Ptr(3)
	return Directory{Struct: p.Struct()}, err
}

func (s Node) HasDirectory() bool {
	if s.Struct.Uint16(8) != 1 {
		return false
	}
	p, err := s.Struct.Ptr(3)
	return p.IsValid() || err != nil
}

func (s Node) SetDirectory(v Directory) error {
	s.Struct.SetUint16(8, 1)
	return s.Struct.SetPtr(3, v.Struct.ToPtr())
}

// NewDirectory sets the directory field to a newly
// allocated Directory struct, preferring placement in s's segment.
func (s Node) NewDirectory() (Directory, error) {
	s.Struct.SetUint16(8, 1)
	ss, err := NewDirectory(s.Struct.Segment())
	if err != nil {
		return Directory{}, err
	}
	err = s.Struct.SetPtr(3, ss.Struct.ToPtr())
	return ss, err
}

func (s Node) File() (File, error) {
	p, err := s.Struct.Ptr(3)
	return File{Struct: p.Struct()}, err
}

func (s Node) HasFile() bool {
	if s.Struct.Uint16(8) != 2 {
		return false
	}
	p, err := s.Struct.Ptr(3)
	return p.IsValid() || err != nil
}

func (s Node) SetFile(v File) error {
	s.Struct.SetUint16(8, 2)
	return s.Struct.SetPtr(3, v.Struct.ToPtr())
}

// NewFile sets the file field to a newly
// allocated File struct, preferring placement in s's segment.
func (s Node) NewFile() (File, error) {
	s.Struct.SetUint16(8, 2)
	ss, err := NewFile(s.Struct.Segment())
	if err != nil {
		return File{}, err
	}
	err = s.Struct.SetPtr(3, ss.Struct.ToPtr())
	return ss, err
}

func (s Node) Ghost() (Ghost, error) {
	p, err := s.Struct.Ptr(3)
	return Ghost{Struct: p.Struct()}, err
}

func (s Node) HasGhost() bool {
	if s.Struct.Uint16(8) != 3 {
		return false
	}
	p, err := s.Struct.Ptr(3)
	return p.IsValid() || err != nil
}

func (s Node) SetGhost(v Ghost) error {
	s.Struct.SetUint16(8, 3)
	return s.Struct.SetPtr(3, v.Struct.ToPtr())
}

// NewGhost sets the ghost field to a newly
// allocated Ghost struct, preferring placement in s's segment.
func (s Node) NewGhost() (Ghost, error) {
	s.Struct.SetUint16(8, 3)
	ss, err := NewGhost(s.Struct.Segment())
	if err != nil {
		return Ghost{}, err
	}
	err = s.Struct.SetPtr(3, ss.Struct.ToPtr())
	return ss, err
}

// Node_List is a list of Node.
type Node_List struct{ capnp.List }

// NewNode creates a new list of Node.
func NewNode_List(s *capnp.Segment, sz int32) (Node_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 16, PointerCount: 4}, sz)
	return Node_List{l}, err
}

func (s Node_List) At(i int) Node { return Node{s.List.Struct(i)} }

func (s Node_List) Set(i int, v Node) error { return s.List.SetStruct(i, v.Struct) }

func (s Node_List) String() string {
	str, _ := text.MarshalList(0xa629eb7f7066fae3, s.List)
	return str
}

// Node_Promise is a wrapper for a Node promised by a client call.
type Node_Promise struct{ *capnp.Pipeline }

func (p Node_Promise) Struct() (Node, error) {
	s, err := p.Pipeline.Struct()
	return Node{s}, err
}

func (p Node_Promise) Commit() Commit_Promise {
	return Commit_Promise{Pipeline: p.Pipeline.GetPipeline(3)}
}

func (p Node_Promise) Directory() Directory_Promise {
	return Directory_Promise{Pipeline: p.Pipeline.GetPipeline(3)}
}

func (p Node_Promise) File() File_Promise {
	return File_Promise{Pipeline: p.Pipeline.GetPipeline(3)}
}

func (p Node_Promise) Ghost() Ghost_Promise {
	return Ghost_Promise{Pipeline: p.Pipeline.GetPipeline(3)}
}

const schema_9195d073cb5c5953 = "x\xda\xac\x95m\x88\x15U\x18\xc7\x9f\xff9s\xef\xec" +
	"\x85\xdd\xf6\x8e\xb3\x0b\x05-s\x10#\x954\xdf\xe8E" +
	"\x90um7\xd3\\\xd9c\x0a\x09\x16\x8d\xf7\x9e\xbd3" +
	"x\xef\xcc23\xb6\xdd(VCAK%\xa9\xa0\xc0" +
	"H\x02\x0d\xfa\xe2w\xa1/Bo\xd6\x07\x8d,\x0d\x85" +
	"2#\xe9\x8d\x8a\xa0\x92t\xe2\xdc{\xf7\xde\xab\xdd\xb5" +
	">\xf4q\x9eyf\xcey~\xff\xff\xf9\x9fE\xd7\xd8" +
	"J\xb683e\x10\xc9e\x99l\xfa\xdd\xacC\x97?" +
	"\x9f\xfb\xc1\x0e\x92\xb3\x80\xf4\x91\xcd[>\x8aO\xbdr" +
	"\x90F`r\"\xfb\x0ev\xc2^\xc0L{\x01sl" +
	"\x9f\x1d#\xa4\xaf;\x0fO>\xf9S\xff\x0bd\xcdj" +
	"k\xcf0\x93\xc8\xce\xf0/l\x8b\x9b\xb6\xc5\x1d{\x05" +
	"\x9f$\xa4\xc7\x1e\xdb\x18\xbcg\x1f\xde\x7f\xc3\xdf3Y" +
	"\xdd\xfe\x1a?i\x1f\xe1\xa6}\x84;K\xcfp\x07\x84" +
	"t\xd3\xe2\xbd\xf7\xae\xb8\xff\xad\x037\xf6\x1b\xba\xffC" +
	"\xe3\xb8}\xda0\xed\xd3\x86\xb3\xf4O\xa3\xd6\xff\xf5\x95" +
	"\xf1\x89\xa9\xef\xe7\x1d\xd5\xfd\xacm\xf7\x86i\xc0\xb0{" +
	"\xb2\xc7\xed\xfe\xaci\xf7g\x9d\xa5k\xb2\xef\xeb\x0fp" +
	"\xe8\xb9\xf2\xa2\xcd\xeb.\xde\xb8\x00\xd7\x0bX]\x17\xed" +
	"\x81.\xd3\x1e\xe8r\xecM]\xdf\x12\xd2\x91\xddG\x0e" +
	"\xdcy\xe1\x9e\xdf;M;\x94;i\x8f\xe6L{4" +
	"\xe7\xd8\xd5\x9c\x9e\xf6p\xee\xed\xa3\x97VDWH\xde" +
	"\x86\xb6\xd9\xfb\xb3&\x88\xec3\xb9+\x04\xfb\\\xee\x18" +
	"\xdd\x97\x16\xdc\x89`\xe2\xeeJ\xc8\x8a\xaa\xbc\xb0\xf6\xb0" +
	"|\xb5\x17\xc6\x09\x8d\x01\xd2\x00K\x1f\x7f\xe9\x0d\xf9\xce" +
	"g\xcf\xbfK\xd2`\x18\xba\x0b\xe8&Z\x8cO\x90\xd6" +
	"\xda\x84\x1fd\x8b~\xc1MT,\x12\xcfM\x84+\x0a" +
	"*J\\?\x10AXTb\xd2\x8d\x85\x9b\x88\xc4\xf3" +
	"c1\xe1&\x9e\x08\x83\x02\x14\x91\xec\xe6Fw\x9a\x1a" +
	" \xb2F\x96\x13\xc9\x95\x1cr\x1dC\x0f\xae\xa5}\xd0" +
	"\xd55\x1b\x88\xe4C\x1cr#C\x0f\xbb\x9a\xf6\x81\x11" +
	"Yr>\x91\\\xc7!\x1fe\x18,\x84\x95\x8a\x9f " +
	"\xdf\x1a\x90\x80<!-\xfa\x91*$aD\xa8\"\xdf" +
	"\"]\x7f\xdb;\xee\x97\x15\xf2-\x85\x1b\x1fu 1" +
	"\xecG#A\xc2\xa3jg\x18\xb7\xd7`X8\x99\x0e" +
	"\x89\xd8\x0fJe\xc5\xc4\xf4\xd2U\xa1\x82$\xaa\x12d" +
	"\x177\x88j\x83\xce\xd3\x9b\x9f\xc3!\x171X@}" +
	"\xce\x05\xba8\x97C.c\xe8\x0d\xdc\x8aB71t" +
	"\x13z=7\xf6\xd0C\x0c=\x9dw\xf7@m|\x9a" +
	"A)\xd1Pj6\xd2z\xa3\xf0y,\\\x11\xabD" +
	"\x84\xe3\xa2\xe0\xb9AI\x8b\x16\x8a 4\x8b*&\x92" +
	"}\xcd\x9d>\xbb\x8aH>\xc5!w\xb5\xedt\xa7\xd6" +
	"\xe9\x19\x0e\xb9\x87\xc1b\xac.\xc8n]\xdc\xc1!\xf7" +
	"1X\x9c\xf7\x81\x13Y{\xf5L\xbb8\xe4\x8b\x0c0" +
	"\xd0\xe6Gk\xff\x12bS\x15\x15\xc7n\xa99\xe9\xa0" +
	"\xbb=\xf1\xc2\x08\xf9\x96\xcb\xeb\x9a\x0cN\xb8\x91\x0a\x92" +
	"i\x08\xbdQ\x186\x1f\x9c\x8a\x8aJ\xaa\xc9\x05\xd3\\" +
	"\x06'\x96?\xe8\x97Ug(\xb76\x14;\x91\x0e\x89" +
	"\xb2r\xc7E\xc0\xb4K\xfd@$\x9e\x12\xa3\xc3C\xab" +
	"\xe9z\x0e\xf3[\x1c:c@\x03\xc3\xecv\x0c\xac\x81" +
	"a-\x91\xdc\xc3!_f\xb0\x0c\xde\x07\x83\xc8:\xa8" +
	"\xd1\xee\xe3\x90\xaf2\xf4\xc6\xfe\xd3\x0a9b\xc8\xb5\x86" +
	"m@1\xb7\xa9jS\xfd\x92>k\x1b\xd48\x11M" +
	"\xd7\xa6\x0aa\x90\xb4\xc1\xe9Hb}X\x9c\x81\xc4\x9c" +
	"\x86=\xd6\"]_C\x10\x0b\xc3\xad\x9f\xd9\x06\x8d\x8a" +
	"\x8a\xb6\x95\x95(\xba%\xed\x97\xad\x91_\"H\xd1D" +
	"sZ\xa3\xf9\x98C\x9em\xb3\xc8\x19]<\xc5!\xcf" +
	"\xb7Y\xe4\x9c\x9e\xf8S\x0e\xf9%\x03\x1a\x0e\xb9\xb0\x84" +
	"H\x9e\xe5\x90\x97\x18\x06\x8c4m\xc0\xf9J\xb3=\xcf" +
	"!/3\x0cd\xae\xe9r\x86\xc8\xfaFg\xc1%\x0e" +
	"\xf93\xc3@\xf6\xaa.g\x89\xac\x1f\xf5j\x979\xe4" +
	"o\x0c\x03\xe6_\xbal\x12Y\xbf\xe8\x7f\xff\xc0!\xff" +
	"\xb8\xd9\x89\x9a\xaa\x84\xc5\x8d~\xeb\xa5\xe3\xeb\xe1\x9bb" +
	"\xfc\xaf\xd9\xe2\xd4\xf4C\xbeu\xc7\xdd4sT!1" +
	"\xc3\x99BgnC\xb87\x91\x0e7\xb6\x92\xa9\x0am" +
	"\x06\xd7\x0fb\x11\x06J\x84\x91\xa8\x84\x91jf\x91\xaf" +
	"b]\x1b\xf7\xcdr\xed\x9c\xe7\x9b\"\xba\x9a\xe0\x16\x0e" +
	"\xe9\xb5\xfc\xad\xb4\x06Op\xc8r\x9b\xbf}me\xaf" +
	"\x11\x08\xd3\xfe\xde\xb9\xb6e\xfa\x9bY9-x~\xb9" +
	"\x18\xa9@\xdb\xf7\x16\xc2\x18\x07\xf2\xad\xfb\x9b\xa0\x8b\x9d" +
	"<\xde\x89\xcf\x98\x8a\xe20\x98)\xf5\xa6#\xf9\xd7\xb4" +
	"\xde'*\xcc/y\x89\xd8\xaa\x84\x1bTk.v\xc4" +
	"\xf6XED\xed\xc1\xbc\xe4_\x82\xd9\xf1\x8bm\xf3\xfc" +
	"\xc7d\x1e\\X\x8b(}\xd5\x11\xd5\xfd=\xb2\xaau" +
	"\xd5Y0\xea\xee^\xa3\x17\x1a\xe6\x90c\x1ax\xa6\xee" +
	"\xed\xd1\xf9\xad\xebo\xca\x8fG\xf5\x9f\x00b\x00\xa1w" +
	"\xd2O\xbc\x7f\xa6\xe5u\xdb\xfa;\x00\x00\xff\xffV\xbe" +
	"R\xda"

func init() {
	schemas.Register(schema_9195d073cb5c5953,
		0x80c828d7e89c12ea,
		0x8b15ee76774b1f9d,
		0x8da013c66e545daf,
		0x8ea7393d37893155,
		0xa629eb7f7066fae3,
		0xe24c59306c829c01,
		0xf736dd278ea58545,
		0xfa723de4a6aa09a0)
}
