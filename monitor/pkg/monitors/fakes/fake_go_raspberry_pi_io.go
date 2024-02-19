// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"github.com/andrew-bodine/raspi/monitor/pkg/monitors"
	rpio "github.com/stianeikeland/go-rpio"
)

type FakeGoRaspberryPiIO struct {
	CloseStub        func() error
	closeMutex       sync.RWMutex
	closeArgsForCall []struct {
	}
	closeReturns struct {
		result1 error
	}
	closeReturnsOnCall map[int]struct {
		result1 error
	}
	DetectEdgeStub        func(monitors.GoRaspberryPiIOPin, rpio.Edge)
	detectEdgeMutex       sync.RWMutex
	detectEdgeArgsForCall []struct {
		arg1 monitors.GoRaspberryPiIOPin
		arg2 rpio.Edge
	}
	DisableIRQsStub        func(uint64)
	disableIRQsMutex       sync.RWMutex
	disableIRQsArgsForCall []struct {
		arg1 uint64
	}
	EdgeDetectedStub        func(monitors.GoRaspberryPiIOPin) bool
	edgeDetectedMutex       sync.RWMutex
	edgeDetectedArgsForCall []struct {
		arg1 monitors.GoRaspberryPiIOPin
	}
	edgeDetectedReturns struct {
		result1 bool
	}
	edgeDetectedReturnsOnCall map[int]struct {
		result1 bool
	}
	EnableIRQsStub        func(uint64)
	enableIRQsMutex       sync.RWMutex
	enableIRQsArgsForCall []struct {
		arg1 uint64
	}
	OpenStub        func() error
	openMutex       sync.RWMutex
	openArgsForCall []struct {
	}
	openReturns struct {
		result1 error
	}
	openReturnsOnCall map[int]struct {
		result1 error
	}
	PinModeStub        func(monitors.GoRaspberryPiIOPin, rpio.Mode)
	pinModeMutex       sync.RWMutex
	pinModeArgsForCall []struct {
		arg1 monitors.GoRaspberryPiIOPin
		arg2 rpio.Mode
	}
	PullModeStub        func(monitors.GoRaspberryPiIOPin, rpio.Pull)
	pullModeMutex       sync.RWMutex
	pullModeArgsForCall []struct {
		arg1 monitors.GoRaspberryPiIOPin
		arg2 rpio.Pull
	}
	ReadPinStub        func(monitors.GoRaspberryPiIOPin) rpio.State
	readPinMutex       sync.RWMutex
	readPinArgsForCall []struct {
		arg1 monitors.GoRaspberryPiIOPin
	}
	readPinReturns struct {
		result1 rpio.State
	}
	readPinReturnsOnCall map[int]struct {
		result1 rpio.State
	}
	SetDutyCycleStub        func(monitors.GoRaspberryPiIOPin, uint32, uint32)
	setDutyCycleMutex       sync.RWMutex
	setDutyCycleArgsForCall []struct {
		arg1 monitors.GoRaspberryPiIOPin
		arg2 uint32
		arg3 uint32
	}
	SetFreqStub        func(monitors.GoRaspberryPiIOPin, int)
	setFreqMutex       sync.RWMutex
	setFreqArgsForCall []struct {
		arg1 monitors.GoRaspberryPiIOPin
		arg2 int
	}
	StartPwmStub        func()
	startPwmMutex       sync.RWMutex
	startPwmArgsForCall []struct {
	}
	StopPwmStub        func()
	stopPwmMutex       sync.RWMutex
	stopPwmArgsForCall []struct {
	}
	TogglePinStub        func(monitors.GoRaspberryPiIOPin)
	togglePinMutex       sync.RWMutex
	togglePinArgsForCall []struct {
		arg1 monitors.GoRaspberryPiIOPin
	}
	WritePinStub        func(monitors.GoRaspberryPiIOPin, rpio.State)
	writePinMutex       sync.RWMutex
	writePinArgsForCall []struct {
		arg1 monitors.GoRaspberryPiIOPin
		arg2 rpio.State
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeGoRaspberryPiIO) Close() error {
	fake.closeMutex.Lock()
	ret, specificReturn := fake.closeReturnsOnCall[len(fake.closeArgsForCall)]
	fake.closeArgsForCall = append(fake.closeArgsForCall, struct {
	}{})
	stub := fake.CloseStub
	fakeReturns := fake.closeReturns
	fake.recordInvocation("Close", []interface{}{})
	fake.closeMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeGoRaspberryPiIO) CloseCallCount() int {
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	return len(fake.closeArgsForCall)
}

func (fake *FakeGoRaspberryPiIO) CloseCalls(stub func() error) {
	fake.closeMutex.Lock()
	defer fake.closeMutex.Unlock()
	fake.CloseStub = stub
}

func (fake *FakeGoRaspberryPiIO) CloseReturns(result1 error) {
	fake.closeMutex.Lock()
	defer fake.closeMutex.Unlock()
	fake.CloseStub = nil
	fake.closeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeGoRaspberryPiIO) CloseReturnsOnCall(i int, result1 error) {
	fake.closeMutex.Lock()
	defer fake.closeMutex.Unlock()
	fake.CloseStub = nil
	if fake.closeReturnsOnCall == nil {
		fake.closeReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.closeReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeGoRaspberryPiIO) DetectEdge(arg1 monitors.GoRaspberryPiIOPin, arg2 rpio.Edge) {
	fake.detectEdgeMutex.Lock()
	fake.detectEdgeArgsForCall = append(fake.detectEdgeArgsForCall, struct {
		arg1 monitors.GoRaspberryPiIOPin
		arg2 rpio.Edge
	}{arg1, arg2})
	stub := fake.DetectEdgeStub
	fake.recordInvocation("DetectEdge", []interface{}{arg1, arg2})
	fake.detectEdgeMutex.Unlock()
	if stub != nil {
		fake.DetectEdgeStub(arg1, arg2)
	}
}

func (fake *FakeGoRaspberryPiIO) DetectEdgeCallCount() int {
	fake.detectEdgeMutex.RLock()
	defer fake.detectEdgeMutex.RUnlock()
	return len(fake.detectEdgeArgsForCall)
}

func (fake *FakeGoRaspberryPiIO) DetectEdgeCalls(stub func(monitors.GoRaspberryPiIOPin, rpio.Edge)) {
	fake.detectEdgeMutex.Lock()
	defer fake.detectEdgeMutex.Unlock()
	fake.DetectEdgeStub = stub
}

func (fake *FakeGoRaspberryPiIO) DetectEdgeArgsForCall(i int) (monitors.GoRaspberryPiIOPin, rpio.Edge) {
	fake.detectEdgeMutex.RLock()
	defer fake.detectEdgeMutex.RUnlock()
	argsForCall := fake.detectEdgeArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeGoRaspberryPiIO) DisableIRQs(arg1 uint64) {
	fake.disableIRQsMutex.Lock()
	fake.disableIRQsArgsForCall = append(fake.disableIRQsArgsForCall, struct {
		arg1 uint64
	}{arg1})
	stub := fake.DisableIRQsStub
	fake.recordInvocation("DisableIRQs", []interface{}{arg1})
	fake.disableIRQsMutex.Unlock()
	if stub != nil {
		fake.DisableIRQsStub(arg1)
	}
}

func (fake *FakeGoRaspberryPiIO) DisableIRQsCallCount() int {
	fake.disableIRQsMutex.RLock()
	defer fake.disableIRQsMutex.RUnlock()
	return len(fake.disableIRQsArgsForCall)
}

func (fake *FakeGoRaspberryPiIO) DisableIRQsCalls(stub func(uint64)) {
	fake.disableIRQsMutex.Lock()
	defer fake.disableIRQsMutex.Unlock()
	fake.DisableIRQsStub = stub
}

func (fake *FakeGoRaspberryPiIO) DisableIRQsArgsForCall(i int) uint64 {
	fake.disableIRQsMutex.RLock()
	defer fake.disableIRQsMutex.RUnlock()
	argsForCall := fake.disableIRQsArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeGoRaspberryPiIO) EdgeDetected(arg1 monitors.GoRaspberryPiIOPin) bool {
	fake.edgeDetectedMutex.Lock()
	ret, specificReturn := fake.edgeDetectedReturnsOnCall[len(fake.edgeDetectedArgsForCall)]
	fake.edgeDetectedArgsForCall = append(fake.edgeDetectedArgsForCall, struct {
		arg1 monitors.GoRaspberryPiIOPin
	}{arg1})
	stub := fake.EdgeDetectedStub
	fakeReturns := fake.edgeDetectedReturns
	fake.recordInvocation("EdgeDetected", []interface{}{arg1})
	fake.edgeDetectedMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeGoRaspberryPiIO) EdgeDetectedCallCount() int {
	fake.edgeDetectedMutex.RLock()
	defer fake.edgeDetectedMutex.RUnlock()
	return len(fake.edgeDetectedArgsForCall)
}

func (fake *FakeGoRaspberryPiIO) EdgeDetectedCalls(stub func(monitors.GoRaspberryPiIOPin) bool) {
	fake.edgeDetectedMutex.Lock()
	defer fake.edgeDetectedMutex.Unlock()
	fake.EdgeDetectedStub = stub
}

func (fake *FakeGoRaspberryPiIO) EdgeDetectedArgsForCall(i int) monitors.GoRaspberryPiIOPin {
	fake.edgeDetectedMutex.RLock()
	defer fake.edgeDetectedMutex.RUnlock()
	argsForCall := fake.edgeDetectedArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeGoRaspberryPiIO) EdgeDetectedReturns(result1 bool) {
	fake.edgeDetectedMutex.Lock()
	defer fake.edgeDetectedMutex.Unlock()
	fake.EdgeDetectedStub = nil
	fake.edgeDetectedReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeGoRaspberryPiIO) EdgeDetectedReturnsOnCall(i int, result1 bool) {
	fake.edgeDetectedMutex.Lock()
	defer fake.edgeDetectedMutex.Unlock()
	fake.EdgeDetectedStub = nil
	if fake.edgeDetectedReturnsOnCall == nil {
		fake.edgeDetectedReturnsOnCall = make(map[int]struct {
			result1 bool
		})
	}
	fake.edgeDetectedReturnsOnCall[i] = struct {
		result1 bool
	}{result1}
}

func (fake *FakeGoRaspberryPiIO) EnableIRQs(arg1 uint64) {
	fake.enableIRQsMutex.Lock()
	fake.enableIRQsArgsForCall = append(fake.enableIRQsArgsForCall, struct {
		arg1 uint64
	}{arg1})
	stub := fake.EnableIRQsStub
	fake.recordInvocation("EnableIRQs", []interface{}{arg1})
	fake.enableIRQsMutex.Unlock()
	if stub != nil {
		fake.EnableIRQsStub(arg1)
	}
}

func (fake *FakeGoRaspberryPiIO) EnableIRQsCallCount() int {
	fake.enableIRQsMutex.RLock()
	defer fake.enableIRQsMutex.RUnlock()
	return len(fake.enableIRQsArgsForCall)
}

func (fake *FakeGoRaspberryPiIO) EnableIRQsCalls(stub func(uint64)) {
	fake.enableIRQsMutex.Lock()
	defer fake.enableIRQsMutex.Unlock()
	fake.EnableIRQsStub = stub
}

func (fake *FakeGoRaspberryPiIO) EnableIRQsArgsForCall(i int) uint64 {
	fake.enableIRQsMutex.RLock()
	defer fake.enableIRQsMutex.RUnlock()
	argsForCall := fake.enableIRQsArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeGoRaspberryPiIO) Open() error {
	fake.openMutex.Lock()
	ret, specificReturn := fake.openReturnsOnCall[len(fake.openArgsForCall)]
	fake.openArgsForCall = append(fake.openArgsForCall, struct {
	}{})
	stub := fake.OpenStub
	fakeReturns := fake.openReturns
	fake.recordInvocation("Open", []interface{}{})
	fake.openMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeGoRaspberryPiIO) OpenCallCount() int {
	fake.openMutex.RLock()
	defer fake.openMutex.RUnlock()
	return len(fake.openArgsForCall)
}

func (fake *FakeGoRaspberryPiIO) OpenCalls(stub func() error) {
	fake.openMutex.Lock()
	defer fake.openMutex.Unlock()
	fake.OpenStub = stub
}

func (fake *FakeGoRaspberryPiIO) OpenReturns(result1 error) {
	fake.openMutex.Lock()
	defer fake.openMutex.Unlock()
	fake.OpenStub = nil
	fake.openReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeGoRaspberryPiIO) OpenReturnsOnCall(i int, result1 error) {
	fake.openMutex.Lock()
	defer fake.openMutex.Unlock()
	fake.OpenStub = nil
	if fake.openReturnsOnCall == nil {
		fake.openReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.openReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeGoRaspberryPiIO) PinMode(arg1 monitors.GoRaspberryPiIOPin, arg2 rpio.Mode) {
	fake.pinModeMutex.Lock()
	fake.pinModeArgsForCall = append(fake.pinModeArgsForCall, struct {
		arg1 monitors.GoRaspberryPiIOPin
		arg2 rpio.Mode
	}{arg1, arg2})
	stub := fake.PinModeStub
	fake.recordInvocation("PinMode", []interface{}{arg1, arg2})
	fake.pinModeMutex.Unlock()
	if stub != nil {
		fake.PinModeStub(arg1, arg2)
	}
}

func (fake *FakeGoRaspberryPiIO) PinModeCallCount() int {
	fake.pinModeMutex.RLock()
	defer fake.pinModeMutex.RUnlock()
	return len(fake.pinModeArgsForCall)
}

func (fake *FakeGoRaspberryPiIO) PinModeCalls(stub func(monitors.GoRaspberryPiIOPin, rpio.Mode)) {
	fake.pinModeMutex.Lock()
	defer fake.pinModeMutex.Unlock()
	fake.PinModeStub = stub
}

func (fake *FakeGoRaspberryPiIO) PinModeArgsForCall(i int) (monitors.GoRaspberryPiIOPin, rpio.Mode) {
	fake.pinModeMutex.RLock()
	defer fake.pinModeMutex.RUnlock()
	argsForCall := fake.pinModeArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeGoRaspberryPiIO) PullMode(arg1 monitors.GoRaspberryPiIOPin, arg2 rpio.Pull) {
	fake.pullModeMutex.Lock()
	fake.pullModeArgsForCall = append(fake.pullModeArgsForCall, struct {
		arg1 monitors.GoRaspberryPiIOPin
		arg2 rpio.Pull
	}{arg1, arg2})
	stub := fake.PullModeStub
	fake.recordInvocation("PullMode", []interface{}{arg1, arg2})
	fake.pullModeMutex.Unlock()
	if stub != nil {
		fake.PullModeStub(arg1, arg2)
	}
}

func (fake *FakeGoRaspberryPiIO) PullModeCallCount() int {
	fake.pullModeMutex.RLock()
	defer fake.pullModeMutex.RUnlock()
	return len(fake.pullModeArgsForCall)
}

func (fake *FakeGoRaspberryPiIO) PullModeCalls(stub func(monitors.GoRaspberryPiIOPin, rpio.Pull)) {
	fake.pullModeMutex.Lock()
	defer fake.pullModeMutex.Unlock()
	fake.PullModeStub = stub
}

func (fake *FakeGoRaspberryPiIO) PullModeArgsForCall(i int) (monitors.GoRaspberryPiIOPin, rpio.Pull) {
	fake.pullModeMutex.RLock()
	defer fake.pullModeMutex.RUnlock()
	argsForCall := fake.pullModeArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeGoRaspberryPiIO) ReadPin(arg1 monitors.GoRaspberryPiIOPin) rpio.State {
	fake.readPinMutex.Lock()
	ret, specificReturn := fake.readPinReturnsOnCall[len(fake.readPinArgsForCall)]
	fake.readPinArgsForCall = append(fake.readPinArgsForCall, struct {
		arg1 monitors.GoRaspberryPiIOPin
	}{arg1})
	stub := fake.ReadPinStub
	fakeReturns := fake.readPinReturns
	fake.recordInvocation("ReadPin", []interface{}{arg1})
	fake.readPinMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeGoRaspberryPiIO) ReadPinCallCount() int {
	fake.readPinMutex.RLock()
	defer fake.readPinMutex.RUnlock()
	return len(fake.readPinArgsForCall)
}

func (fake *FakeGoRaspberryPiIO) ReadPinCalls(stub func(monitors.GoRaspberryPiIOPin) rpio.State) {
	fake.readPinMutex.Lock()
	defer fake.readPinMutex.Unlock()
	fake.ReadPinStub = stub
}

func (fake *FakeGoRaspberryPiIO) ReadPinArgsForCall(i int) monitors.GoRaspberryPiIOPin {
	fake.readPinMutex.RLock()
	defer fake.readPinMutex.RUnlock()
	argsForCall := fake.readPinArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeGoRaspberryPiIO) ReadPinReturns(result1 rpio.State) {
	fake.readPinMutex.Lock()
	defer fake.readPinMutex.Unlock()
	fake.ReadPinStub = nil
	fake.readPinReturns = struct {
		result1 rpio.State
	}{result1}
}

func (fake *FakeGoRaspberryPiIO) ReadPinReturnsOnCall(i int, result1 rpio.State) {
	fake.readPinMutex.Lock()
	defer fake.readPinMutex.Unlock()
	fake.ReadPinStub = nil
	if fake.readPinReturnsOnCall == nil {
		fake.readPinReturnsOnCall = make(map[int]struct {
			result1 rpio.State
		})
	}
	fake.readPinReturnsOnCall[i] = struct {
		result1 rpio.State
	}{result1}
}

func (fake *FakeGoRaspberryPiIO) SetDutyCycle(arg1 monitors.GoRaspberryPiIOPin, arg2 uint32, arg3 uint32) {
	fake.setDutyCycleMutex.Lock()
	fake.setDutyCycleArgsForCall = append(fake.setDutyCycleArgsForCall, struct {
		arg1 monitors.GoRaspberryPiIOPin
		arg2 uint32
		arg3 uint32
	}{arg1, arg2, arg3})
	stub := fake.SetDutyCycleStub
	fake.recordInvocation("SetDutyCycle", []interface{}{arg1, arg2, arg3})
	fake.setDutyCycleMutex.Unlock()
	if stub != nil {
		fake.SetDutyCycleStub(arg1, arg2, arg3)
	}
}

func (fake *FakeGoRaspberryPiIO) SetDutyCycleCallCount() int {
	fake.setDutyCycleMutex.RLock()
	defer fake.setDutyCycleMutex.RUnlock()
	return len(fake.setDutyCycleArgsForCall)
}

func (fake *FakeGoRaspberryPiIO) SetDutyCycleCalls(stub func(monitors.GoRaspberryPiIOPin, uint32, uint32)) {
	fake.setDutyCycleMutex.Lock()
	defer fake.setDutyCycleMutex.Unlock()
	fake.SetDutyCycleStub = stub
}

func (fake *FakeGoRaspberryPiIO) SetDutyCycleArgsForCall(i int) (monitors.GoRaspberryPiIOPin, uint32, uint32) {
	fake.setDutyCycleMutex.RLock()
	defer fake.setDutyCycleMutex.RUnlock()
	argsForCall := fake.setDutyCycleArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeGoRaspberryPiIO) SetFreq(arg1 monitors.GoRaspberryPiIOPin, arg2 int) {
	fake.setFreqMutex.Lock()
	fake.setFreqArgsForCall = append(fake.setFreqArgsForCall, struct {
		arg1 monitors.GoRaspberryPiIOPin
		arg2 int
	}{arg1, arg2})
	stub := fake.SetFreqStub
	fake.recordInvocation("SetFreq", []interface{}{arg1, arg2})
	fake.setFreqMutex.Unlock()
	if stub != nil {
		fake.SetFreqStub(arg1, arg2)
	}
}

func (fake *FakeGoRaspberryPiIO) SetFreqCallCount() int {
	fake.setFreqMutex.RLock()
	defer fake.setFreqMutex.RUnlock()
	return len(fake.setFreqArgsForCall)
}

func (fake *FakeGoRaspberryPiIO) SetFreqCalls(stub func(monitors.GoRaspberryPiIOPin, int)) {
	fake.setFreqMutex.Lock()
	defer fake.setFreqMutex.Unlock()
	fake.SetFreqStub = stub
}

func (fake *FakeGoRaspberryPiIO) SetFreqArgsForCall(i int) (monitors.GoRaspberryPiIOPin, int) {
	fake.setFreqMutex.RLock()
	defer fake.setFreqMutex.RUnlock()
	argsForCall := fake.setFreqArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeGoRaspberryPiIO) StartPwm() {
	fake.startPwmMutex.Lock()
	fake.startPwmArgsForCall = append(fake.startPwmArgsForCall, struct {
	}{})
	stub := fake.StartPwmStub
	fake.recordInvocation("StartPwm", []interface{}{})
	fake.startPwmMutex.Unlock()
	if stub != nil {
		fake.StartPwmStub()
	}
}

func (fake *FakeGoRaspberryPiIO) StartPwmCallCount() int {
	fake.startPwmMutex.RLock()
	defer fake.startPwmMutex.RUnlock()
	return len(fake.startPwmArgsForCall)
}

func (fake *FakeGoRaspberryPiIO) StartPwmCalls(stub func()) {
	fake.startPwmMutex.Lock()
	defer fake.startPwmMutex.Unlock()
	fake.StartPwmStub = stub
}

func (fake *FakeGoRaspberryPiIO) StopPwm() {
	fake.stopPwmMutex.Lock()
	fake.stopPwmArgsForCall = append(fake.stopPwmArgsForCall, struct {
	}{})
	stub := fake.StopPwmStub
	fake.recordInvocation("StopPwm", []interface{}{})
	fake.stopPwmMutex.Unlock()
	if stub != nil {
		fake.StopPwmStub()
	}
}

func (fake *FakeGoRaspberryPiIO) StopPwmCallCount() int {
	fake.stopPwmMutex.RLock()
	defer fake.stopPwmMutex.RUnlock()
	return len(fake.stopPwmArgsForCall)
}

func (fake *FakeGoRaspberryPiIO) StopPwmCalls(stub func()) {
	fake.stopPwmMutex.Lock()
	defer fake.stopPwmMutex.Unlock()
	fake.StopPwmStub = stub
}

func (fake *FakeGoRaspberryPiIO) TogglePin(arg1 monitors.GoRaspberryPiIOPin) {
	fake.togglePinMutex.Lock()
	fake.togglePinArgsForCall = append(fake.togglePinArgsForCall, struct {
		arg1 monitors.GoRaspberryPiIOPin
	}{arg1})
	stub := fake.TogglePinStub
	fake.recordInvocation("TogglePin", []interface{}{arg1})
	fake.togglePinMutex.Unlock()
	if stub != nil {
		fake.TogglePinStub(arg1)
	}
}

func (fake *FakeGoRaspberryPiIO) TogglePinCallCount() int {
	fake.togglePinMutex.RLock()
	defer fake.togglePinMutex.RUnlock()
	return len(fake.togglePinArgsForCall)
}

func (fake *FakeGoRaspberryPiIO) TogglePinCalls(stub func(monitors.GoRaspberryPiIOPin)) {
	fake.togglePinMutex.Lock()
	defer fake.togglePinMutex.Unlock()
	fake.TogglePinStub = stub
}

func (fake *FakeGoRaspberryPiIO) TogglePinArgsForCall(i int) monitors.GoRaspberryPiIOPin {
	fake.togglePinMutex.RLock()
	defer fake.togglePinMutex.RUnlock()
	argsForCall := fake.togglePinArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeGoRaspberryPiIO) WritePin(arg1 monitors.GoRaspberryPiIOPin, arg2 rpio.State) {
	fake.writePinMutex.Lock()
	fake.writePinArgsForCall = append(fake.writePinArgsForCall, struct {
		arg1 monitors.GoRaspberryPiIOPin
		arg2 rpio.State
	}{arg1, arg2})
	stub := fake.WritePinStub
	fake.recordInvocation("WritePin", []interface{}{arg1, arg2})
	fake.writePinMutex.Unlock()
	if stub != nil {
		fake.WritePinStub(arg1, arg2)
	}
}

func (fake *FakeGoRaspberryPiIO) WritePinCallCount() int {
	fake.writePinMutex.RLock()
	defer fake.writePinMutex.RUnlock()
	return len(fake.writePinArgsForCall)
}

func (fake *FakeGoRaspberryPiIO) WritePinCalls(stub func(monitors.GoRaspberryPiIOPin, rpio.State)) {
	fake.writePinMutex.Lock()
	defer fake.writePinMutex.Unlock()
	fake.WritePinStub = stub
}

func (fake *FakeGoRaspberryPiIO) WritePinArgsForCall(i int) (monitors.GoRaspberryPiIOPin, rpio.State) {
	fake.writePinMutex.RLock()
	defer fake.writePinMutex.RUnlock()
	argsForCall := fake.writePinArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeGoRaspberryPiIO) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	fake.detectEdgeMutex.RLock()
	defer fake.detectEdgeMutex.RUnlock()
	fake.disableIRQsMutex.RLock()
	defer fake.disableIRQsMutex.RUnlock()
	fake.edgeDetectedMutex.RLock()
	defer fake.edgeDetectedMutex.RUnlock()
	fake.enableIRQsMutex.RLock()
	defer fake.enableIRQsMutex.RUnlock()
	fake.openMutex.RLock()
	defer fake.openMutex.RUnlock()
	fake.pinModeMutex.RLock()
	defer fake.pinModeMutex.RUnlock()
	fake.pullModeMutex.RLock()
	defer fake.pullModeMutex.RUnlock()
	fake.readPinMutex.RLock()
	defer fake.readPinMutex.RUnlock()
	fake.setDutyCycleMutex.RLock()
	defer fake.setDutyCycleMutex.RUnlock()
	fake.setFreqMutex.RLock()
	defer fake.setFreqMutex.RUnlock()
	fake.startPwmMutex.RLock()
	defer fake.startPwmMutex.RUnlock()
	fake.stopPwmMutex.RLock()
	defer fake.stopPwmMutex.RUnlock()
	fake.togglePinMutex.RLock()
	defer fake.togglePinMutex.RUnlock()
	fake.writePinMutex.RLock()
	defer fake.writePinMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeGoRaspberryPiIO) recordInvocation(key string, args []interface{}) {
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

var _ monitors.GoRaspberryPiIO = new(FakeGoRaspberryPiIO)
