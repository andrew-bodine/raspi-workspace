// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"context"
	"sync"

	release_raspi_com "github.com/andrew-bodine/release/pkg/protobufs"
)

type FakeReleasesServer struct {
	GetReleasesStub        func(context.Context, *release_raspi_com.GetReleasesRequest) (*release_raspi_com.ReleasesResponse, error)
	getReleasesMutex       sync.RWMutex
	getReleasesArgsForCall []struct {
		arg1 context.Context
		arg2 *release_raspi_com.GetReleasesRequest
	}
	getReleasesReturns struct {
		result1 *release_raspi_com.ReleasesResponse
		result2 error
	}
	getReleasesReturnsOnCall map[int]struct {
		result1 *release_raspi_com.ReleasesResponse
		result2 error
	}
	WatchReleasesStub        func(*release_raspi_com.WatchReleasesRequest, release_raspi_com.Releases_WatchReleasesServer) error
	watchReleasesMutex       sync.RWMutex
	watchReleasesArgsForCall []struct {
		arg1 *release_raspi_com.WatchReleasesRequest
		arg2 release_raspi_com.Releases_WatchReleasesServer
	}
	watchReleasesReturns struct {
		result1 error
	}
	watchReleasesReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeReleasesServer) GetReleases(arg1 context.Context, arg2 *release_raspi_com.GetReleasesRequest) (*release_raspi_com.ReleasesResponse, error) {
	fake.getReleasesMutex.Lock()
	ret, specificReturn := fake.getReleasesReturnsOnCall[len(fake.getReleasesArgsForCall)]
	fake.getReleasesArgsForCall = append(fake.getReleasesArgsForCall, struct {
		arg1 context.Context
		arg2 *release_raspi_com.GetReleasesRequest
	}{arg1, arg2})
	fake.recordInvocation("GetReleases", []interface{}{arg1, arg2})
	fake.getReleasesMutex.Unlock()
	if fake.GetReleasesStub != nil {
		return fake.GetReleasesStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getReleasesReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeReleasesServer) GetReleasesCallCount() int {
	fake.getReleasesMutex.RLock()
	defer fake.getReleasesMutex.RUnlock()
	return len(fake.getReleasesArgsForCall)
}

func (fake *FakeReleasesServer) GetReleasesCalls(stub func(context.Context, *release_raspi_com.GetReleasesRequest) (*release_raspi_com.ReleasesResponse, error)) {
	fake.getReleasesMutex.Lock()
	defer fake.getReleasesMutex.Unlock()
	fake.GetReleasesStub = stub
}

func (fake *FakeReleasesServer) GetReleasesArgsForCall(i int) (context.Context, *release_raspi_com.GetReleasesRequest) {
	fake.getReleasesMutex.RLock()
	defer fake.getReleasesMutex.RUnlock()
	argsForCall := fake.getReleasesArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeReleasesServer) GetReleasesReturns(result1 *release_raspi_com.ReleasesResponse, result2 error) {
	fake.getReleasesMutex.Lock()
	defer fake.getReleasesMutex.Unlock()
	fake.GetReleasesStub = nil
	fake.getReleasesReturns = struct {
		result1 *release_raspi_com.ReleasesResponse
		result2 error
	}{result1, result2}
}

func (fake *FakeReleasesServer) GetReleasesReturnsOnCall(i int, result1 *release_raspi_com.ReleasesResponse, result2 error) {
	fake.getReleasesMutex.Lock()
	defer fake.getReleasesMutex.Unlock()
	fake.GetReleasesStub = nil
	if fake.getReleasesReturnsOnCall == nil {
		fake.getReleasesReturnsOnCall = make(map[int]struct {
			result1 *release_raspi_com.ReleasesResponse
			result2 error
		})
	}
	fake.getReleasesReturnsOnCall[i] = struct {
		result1 *release_raspi_com.ReleasesResponse
		result2 error
	}{result1, result2}
}

func (fake *FakeReleasesServer) WatchReleases(arg1 *release_raspi_com.WatchReleasesRequest, arg2 release_raspi_com.Releases_WatchReleasesServer) error {
	fake.watchReleasesMutex.Lock()
	ret, specificReturn := fake.watchReleasesReturnsOnCall[len(fake.watchReleasesArgsForCall)]
	fake.watchReleasesArgsForCall = append(fake.watchReleasesArgsForCall, struct {
		arg1 *release_raspi_com.WatchReleasesRequest
		arg2 release_raspi_com.Releases_WatchReleasesServer
	}{arg1, arg2})
	fake.recordInvocation("WatchReleases", []interface{}{arg1, arg2})
	fake.watchReleasesMutex.Unlock()
	if fake.WatchReleasesStub != nil {
		return fake.WatchReleasesStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.watchReleasesReturns
	return fakeReturns.result1
}

func (fake *FakeReleasesServer) WatchReleasesCallCount() int {
	fake.watchReleasesMutex.RLock()
	defer fake.watchReleasesMutex.RUnlock()
	return len(fake.watchReleasesArgsForCall)
}

func (fake *FakeReleasesServer) WatchReleasesCalls(stub func(*release_raspi_com.WatchReleasesRequest, release_raspi_com.Releases_WatchReleasesServer) error) {
	fake.watchReleasesMutex.Lock()
	defer fake.watchReleasesMutex.Unlock()
	fake.WatchReleasesStub = stub
}

func (fake *FakeReleasesServer) WatchReleasesArgsForCall(i int) (*release_raspi_com.WatchReleasesRequest, release_raspi_com.Releases_WatchReleasesServer) {
	fake.watchReleasesMutex.RLock()
	defer fake.watchReleasesMutex.RUnlock()
	argsForCall := fake.watchReleasesArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeReleasesServer) WatchReleasesReturns(result1 error) {
	fake.watchReleasesMutex.Lock()
	defer fake.watchReleasesMutex.Unlock()
	fake.WatchReleasesStub = nil
	fake.watchReleasesReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeReleasesServer) WatchReleasesReturnsOnCall(i int, result1 error) {
	fake.watchReleasesMutex.Lock()
	defer fake.watchReleasesMutex.Unlock()
	fake.WatchReleasesStub = nil
	if fake.watchReleasesReturnsOnCall == nil {
		fake.watchReleasesReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.watchReleasesReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeReleasesServer) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getReleasesMutex.RLock()
	defer fake.getReleasesMutex.RUnlock()
	fake.watchReleasesMutex.RLock()
	defer fake.watchReleasesMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeReleasesServer) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ release_raspi_com.ReleasesServer = new(FakeReleasesServer)