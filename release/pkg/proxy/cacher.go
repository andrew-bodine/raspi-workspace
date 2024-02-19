package proxy

import (
	"sync"

	pb "github.com/andrew-bodine/raspi/release/pkg/protobufs"
)

// Cacher declares the interface which the raspi release proxy leverages to
// delegates the task of release cache management to a separate goroutine, and
// bind the cacher goroutine's lifecycle to the lifecycle of the proxy.
//
//go:generate counterfeiter -o fakes/fake_cacher.go --fake-name FakeCacher . Cacher
type Cacher interface {
	Cache() ReleaseCache
	Run() error
	Stop() error
	State() CacherState
}

type ReleaseListOptions struct{}

//go:generate counterfeiter -o fakes/fake_release_lister.go --fake-name FakeReleaseLister . ReleaseLister
type ReleaseLister func(opts *ReleaseListOptions) []pb.Release

// https://developer.github.com/v3/repos/releases/#list-releases-for-a-repository

type ReleaseCache struct {
	releases map[string]pb.Release
}

type CacherState string

const (
	CacherStateStopped CacherState = "Stopped"
	CacherStateRunning CacherState = "Running"
)

// Default release cacher implementation.
type cacher struct {
	cache    ReleaseCache
	cacheMux sync.Mutex

	lister ReleaseLister

	stopCh chan struct{}

	state    CacherState
	stateMux sync.Mutex
}

func NewCacher(lister ReleaseLister) Cacher {
	cm := &cacher{
		lister: lister,
		stopCh: make(chan struct{}),
		state:  CacherStateStopped,
	}

	return cm
}

// Implement the Cacher interface.
func (cm *cacher) Cache() ReleaseCache {
	cm.cacheMux.Lock()
	defer cm.cacheMux.Unlock()
	return cm.cache
}

func (cm *cacher) Run() error {
	if cm.State() == CacherStateRunning {
		return nil
	}

	cm.setState(CacherStateRunning)
	defer cm.setState(CacherStateStopped)

	_ = cm.lister(nil)

	for {
		select {
		case <-cm.stopCh:
			return nil
		}
	}

	return nil
}

func (cm *cacher) Stop() error {
	if cm.State() == CacherStateRunning {
		close(cm.stopCh)
	}

	return nil
}

func (cm *cacher) State() CacherState {
	cm.stateMux.Lock()
	defer cm.stateMux.Unlock()
	return cm.state
}

func (cm *cacher) setState(to CacherState) {
	cm.stateMux.Lock()
	defer cm.stateMux.Unlock()
	cm.state = to
}
