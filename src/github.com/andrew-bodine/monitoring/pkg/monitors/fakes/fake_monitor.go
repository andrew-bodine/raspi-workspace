// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"github.com/andrew-bodine/monitoring/pkg/monitors"
)

type FakeMonitor struct {
	RunStub        func(stopCh <-chan struct{}) error
	runMutex       sync.RWMutex
	runArgsForCall []struct {
		stopCh <-chan struct{}
	}
	runReturns struct {
		result1 error
	}
	runReturnsOnCall map[int]struct {
		result1 error
	}
	GetStateStub        func() interface{}
	getStateMutex       sync.RWMutex
	getStateArgsForCall []struct{}
	getStateReturns     struct {
		result1 interface{}
	}
	getStateReturnsOnCall map[int]struct {
		result1 interface{}
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeMonitor) Run(stopCh <-chan struct{}) error {
	fake.runMutex.Lock()
	ret, specificReturn := fake.runReturnsOnCall[len(fake.runArgsForCall)]
	fake.runArgsForCall = append(fake.runArgsForCall, struct {
		stopCh <-chan struct{}
	}{stopCh})
	fake.recordInvocation("Run", []interface{}{stopCh})
	fake.runMutex.Unlock()
	if fake.RunStub != nil {
		return fake.RunStub(stopCh)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.runReturns.result1
}

func (fake *FakeMonitor) RunCallCount() int {
	fake.runMutex.RLock()
	defer fake.runMutex.RUnlock()
	return len(fake.runArgsForCall)
}

func (fake *FakeMonitor) RunArgsForCall(i int) <-chan struct{} {
	fake.runMutex.RLock()
	defer fake.runMutex.RUnlock()
	return fake.runArgsForCall[i].stopCh
}

func (fake *FakeMonitor) RunReturns(result1 error) {
	fake.RunStub = nil
	fake.runReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeMonitor) RunReturnsOnCall(i int, result1 error) {
	fake.RunStub = nil
	if fake.runReturnsOnCall == nil {
		fake.runReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.runReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeMonitor) GetState() interface{} {
	fake.getStateMutex.Lock()
	ret, specificReturn := fake.getStateReturnsOnCall[len(fake.getStateArgsForCall)]
	fake.getStateArgsForCall = append(fake.getStateArgsForCall, struct{}{})
	fake.recordInvocation("GetState", []interface{}{})
	fake.getStateMutex.Unlock()
	if fake.GetStateStub != nil {
		return fake.GetStateStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.getStateReturns.result1
}

func (fake *FakeMonitor) GetStateCallCount() int {
	fake.getStateMutex.RLock()
	defer fake.getStateMutex.RUnlock()
	return len(fake.getStateArgsForCall)
}

func (fake *FakeMonitor) GetStateReturns(result1 interface{}) {
	fake.GetStateStub = nil
	fake.getStateReturns = struct {
		result1 interface{}
	}{result1}
}

func (fake *FakeMonitor) GetStateReturnsOnCall(i int, result1 interface{}) {
	fake.GetStateStub = nil
	if fake.getStateReturnsOnCall == nil {
		fake.getStateReturnsOnCall = make(map[int]struct {
			result1 interface{}
		})
	}
	fake.getStateReturnsOnCall[i] = struct {
		result1 interface{}
	}{result1}
}

func (fake *FakeMonitor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.runMutex.RLock()
	defer fake.runMutex.RUnlock()
	fake.getStateMutex.RLock()
	defer fake.getStateMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeMonitor) recordInvocation(key string, args []interface{}) {
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

var _ monitors.Monitor = new(FakeMonitor)
