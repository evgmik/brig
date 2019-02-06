package remotesapi

import (
	"fmt"

	"github.com/sahib/brig/catfs"
)

// Mock is for testing purposes whenever a normal RemotesAPI is needed.
// It stores remotes in memory and does not implement really syncing or diffing.
type Mock struct {
	name        string
	fingerprint string
	remotes     map[string]*Remote
	callbacks   []func()
}

// NewMock creates a new Mock.
// `name` and `fingerprint` will be returned
func NewMock(name, fingerprint string) *Mock {
	return &Mock{
		name:        name,
		fingerprint: fingerprint,
		remotes:     make(map[string]*Remote),
	}
}

// List all existing remotes.
func (m *Mock) List() ([]*Remote, error) {
	rmts := []*Remote{}
	for _, rmt := range m.remotes {
		rmts = append(rmts, rmt)
	}

	return rmts, nil
}

// Get a remote by `name`.
func (m *Mock) Get(name string) (*Remote, error) {
	rm, ok := m.remotes[name]
	if !ok {
		return nil, fmt.Errorf("no such remote: %s", name)
	}

	return rm, nil
}

// Set (i.e. add or modify) a remote.
// The mock implementation takes the isOnline, isAuthenticated
// and LastSeen info from the remote, in contrast to the real implementation.
func (m *Mock) Set(rm Remote) error {
	if rm.Name == "" {
		return fmt.Errorf("empty name")
	}

	if rm.Fingerprint == "" {
		return fmt.Errorf("empty fingerprint")
	}

	if rm.Name == m.name {
		return fmt.Errorf("cannot add remote with own name")
	}

	prevRm, ok := m.remotes[rm.Name]
	if ok {
		rm.IsAuthenticated = prevRm.IsAuthenticated
		rm.LastSeen = prevRm.LastSeen
		rm.IsOnline = prevRm.IsOnline
	}

	m.remotes[rm.Name] = &rm
	m.notify()
	return nil
}

// Remove removes a remote by `name`.
func (m *Mock) Remove(name string) error {
	if _, ok := m.remotes[name]; !ok {
		return fmt.Errorf("no such remote: %s", name)
	}

	delete(m.remotes, name)
	m.notify()
	return nil
}

// Self returns the identity of this repository.
func (m *Mock) Self() (Identity, error) {
	return Identity{
		Name:        m.name,
		Fingerprint: m.fingerprint,
	}, nil
}

// Sync synchronizes the latest state of `name` with our latest state.
// The mock implementation does nothing currently.
func (m *Mock) Sync(name string) error {
	if _, ok := m.remotes[name]; !ok {
		return fmt.Errorf("no such remote: %s", name)
	}

	return nil
}

// MakeDiff produces a diff to the remote with `name`.
func (m *Mock) MakeDiff(name string) (*catfs.Diff, error) {
	if _, ok := m.remotes[name]; !ok {
		return nil, fmt.Errorf("no such remote: %s", name)
	}

	// always send an empty diff.
	return &catfs.Diff{
		Added:    []catfs.StatInfo{},
		Removed:  []catfs.StatInfo{},
		Ignored:  []catfs.StatInfo{},
		Missing:  []catfs.StatInfo{},
		Conflict: []catfs.DiffPair{},
		Moved:    []catfs.DiffPair{},
		Merged:   []catfs.DiffPair{},
	}, nil
}

func (m *Mock) notify() {
	for _, fn := range m.callbacks {
		fn()
	}
}

// OnChange register a callback to be called once the remote list changes.
func (m *Mock) OnChange(fn func()) {
	m.callbacks = append(m.callbacks, fn)
}
