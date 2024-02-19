// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"github.com/andrew-bodine/raspi/monitor/pkg/monitors"
	"go.uber.org/zap"
)

type FakeConfigWithFlagsAndLoggerConstructor struct {
	Stub        func(*zap.Logger) (interface{}, error)
	mutex       sync.RWMutex
	argsForCall []struct {
		arg1 *zap.Logger
	}
	returns struct {
		result1 interface{}
		result2 error
	}
	returnsOnCall map[int]struct {
		result1 interface{}
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeConfigWithFlagsAndLoggerConstructor) Spy(arg1 *zap.Logger) (interface{}, error) {
	fake.mutex.Lock()
	ret, specificReturn := fake.returnsOnCall[len(fake.argsForCall)]
	fake.argsForCall = append(fake.argsForCall, struct {
		arg1 *zap.Logger
	}{arg1})
	stub := fake.Stub
	returns := fake.returns
	fake.recordInvocation("ConfigWithFlagsAndLoggerConstructor", []interface{}{arg1})
	fake.mutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return returns.result1, returns.result2
}

func (fake *FakeConfigWithFlagsAndLoggerConstructor) CallCount() int {
	fake.mutex.RLock()
	defer fake.mutex.RUnlock()
	return len(fake.argsForCall)
}

func (fake *FakeConfigWithFlagsAndLoggerConstructor) Calls(stub func(*zap.Logger) (interface{}, error)) {
	fake.mutex.Lock()
	defer fake.mutex.Unlock()
	fake.Stub = stub
}

func (fake *FakeConfigWithFlagsAndLoggerConstructor) ArgsForCall(i int) *zap.Logger {
	fake.mutex.RLock()
	defer fake.mutex.RUnlock()
	return fake.argsForCall[i].arg1
}

func (fake *FakeConfigWithFlagsAndLoggerConstructor) Returns(result1 interface{}, result2 error) {
	fake.mutex.Lock()
	defer fake.mutex.Unlock()
	fake.Stub = nil
	fake.returns = struct {
		result1 interface{}
		result2 error
	}{result1, result2}
}

func (fake *FakeConfigWithFlagsAndLoggerConstructor) ReturnsOnCall(i int, result1 interface{}, result2 error) {
	fake.mutex.Lock()
	defer fake.mutex.Unlock()
	fake.Stub = nil
	if fake.returnsOnCall == nil {
		fake.returnsOnCall = make(map[int]struct {
			result1 interface{}
			result2 error
		})
	}
	fake.returnsOnCall[i] = struct {
		result1 interface{}
		result2 error
	}{result1, result2}
}

func (fake *FakeConfigWithFlagsAndLoggerConstructor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.mutex.RLock()
	defer fake.mutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeConfigWithFlagsAndLoggerConstructor) recordInvocation(key string, args []interface{}) {
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

var _ monitors.ConfigWithFlagsAndLoggerConstructor = new(FakeConfigWithFlagsAndLoggerConstructor).Spy
