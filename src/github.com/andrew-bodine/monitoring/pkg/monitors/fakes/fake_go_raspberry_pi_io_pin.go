// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"github.com/andrew-bodine/vibration-monitor/pkg/monitor"
	gorpio "github.com/stianeikeland/go-rpio"
)

type FakeGoRaspberryPiIOPin struct {
	InputStub         func()
	inputMutex        sync.RWMutex
	inputArgsForCall  []struct{}
	OutputStub        func()
	outputMutex       sync.RWMutex
	outputArgsForCall []struct{}
	ClockStub         func()
	clockMutex        sync.RWMutex
	clockArgsForCall  []struct{}
	PwmStub           func()
	pwmMutex          sync.RWMutex
	pwmArgsForCall    []struct{}
	HighStub          func()
	highMutex         sync.RWMutex
	highArgsForCall   []struct{}
	LowStub           func()
	lowMutex          sync.RWMutex
	lowArgsForCall    []struct{}
	ToggleStub        func()
	toggleMutex       sync.RWMutex
	toggleArgsForCall []struct{}
	FreqStub          func(freq int)
	freqMutex         sync.RWMutex
	freqArgsForCall   []struct {
		freq int
	}
	DutyCycleStub        func(dutyLen, cycleLen uint32)
	dutyCycleMutex       sync.RWMutex
	dutyCycleArgsForCall []struct {
		dutyLen  uint32
		cycleLen uint32
	}
	ModeStub        func(mode gorpio.Mode)
	modeMutex       sync.RWMutex
	modeArgsForCall []struct {
		mode gorpio.Mode
	}
	WriteStub        func(state gorpio.State)
	writeMutex       sync.RWMutex
	writeArgsForCall []struct {
		state gorpio.State
	}
	ReadStub        func() gorpio.State
	readMutex       sync.RWMutex
	readArgsForCall []struct{}
	readReturns     struct {
		result1 gorpio.State
	}
	readReturnsOnCall map[int]struct {
		result1 gorpio.State
	}
	PullStub        func(pull gorpio.Pull)
	pullMutex       sync.RWMutex
	pullArgsForCall []struct {
		pull gorpio.Pull
	}
	PullUpStub         func()
	pullUpMutex        sync.RWMutex
	pullUpArgsForCall  []struct{}
	PullOffStub        func()
	pullOffMutex       sync.RWMutex
	pullOffArgsForCall []struct{}
	DetectStub         func(edge gorpio.Edge)
	detectMutex        sync.RWMutex
	detectArgsForCall  []struct {
		edge gorpio.Edge
	}
	EdgeDetectedStub        func() bool
	edgeDetectedMutex       sync.RWMutex
	edgeDetectedArgsForCall []struct{}
	edgeDetectedReturns     struct {
		result1 bool
	}
	edgeDetectedReturnsOnCall map[int]struct {
		result1 bool
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeGoRaspberryPiIOPin) Input() {
	fake.inputMutex.Lock()
	fake.inputArgsForCall = append(fake.inputArgsForCall, struct{}{})
	fake.recordInvocation("Input", []interface{}{})
	fake.inputMutex.Unlock()
	if fake.InputStub != nil {
		fake.InputStub()
	}
}

func (fake *FakeGoRaspberryPiIOPin) InputCallCount() int {
	fake.inputMutex.RLock()
	defer fake.inputMutex.RUnlock()
	return len(fake.inputArgsForCall)
}

func (fake *FakeGoRaspberryPiIOPin) Output() {
	fake.outputMutex.Lock()
	fake.outputArgsForCall = append(fake.outputArgsForCall, struct{}{})
	fake.recordInvocation("Output", []interface{}{})
	fake.outputMutex.Unlock()
	if fake.OutputStub != nil {
		fake.OutputStub()
	}
}

func (fake *FakeGoRaspberryPiIOPin) OutputCallCount() int {
	fake.outputMutex.RLock()
	defer fake.outputMutex.RUnlock()
	return len(fake.outputArgsForCall)
}

func (fake *FakeGoRaspberryPiIOPin) Clock() {
	fake.clockMutex.Lock()
	fake.clockArgsForCall = append(fake.clockArgsForCall, struct{}{})
	fake.recordInvocation("Clock", []interface{}{})
	fake.clockMutex.Unlock()
	if fake.ClockStub != nil {
		fake.ClockStub()
	}
}

func (fake *FakeGoRaspberryPiIOPin) ClockCallCount() int {
	fake.clockMutex.RLock()
	defer fake.clockMutex.RUnlock()
	return len(fake.clockArgsForCall)
}

func (fake *FakeGoRaspberryPiIOPin) Pwm() {
	fake.pwmMutex.Lock()
	fake.pwmArgsForCall = append(fake.pwmArgsForCall, struct{}{})
	fake.recordInvocation("Pwm", []interface{}{})
	fake.pwmMutex.Unlock()
	if fake.PwmStub != nil {
		fake.PwmStub()
	}
}

func (fake *FakeGoRaspberryPiIOPin) PwmCallCount() int {
	fake.pwmMutex.RLock()
	defer fake.pwmMutex.RUnlock()
	return len(fake.pwmArgsForCall)
}

func (fake *FakeGoRaspberryPiIOPin) High() {
	fake.highMutex.Lock()
	fake.highArgsForCall = append(fake.highArgsForCall, struct{}{})
	fake.recordInvocation("High", []interface{}{})
	fake.highMutex.Unlock()
	if fake.HighStub != nil {
		fake.HighStub()
	}
}

func (fake *FakeGoRaspberryPiIOPin) HighCallCount() int {
	fake.highMutex.RLock()
	defer fake.highMutex.RUnlock()
	return len(fake.highArgsForCall)
}

func (fake *FakeGoRaspberryPiIOPin) Low() {
	fake.lowMutex.Lock()
	fake.lowArgsForCall = append(fake.lowArgsForCall, struct{}{})
	fake.recordInvocation("Low", []interface{}{})
	fake.lowMutex.Unlock()
	if fake.LowStub != nil {
		fake.LowStub()
	}
}

func (fake *FakeGoRaspberryPiIOPin) LowCallCount() int {
	fake.lowMutex.RLock()
	defer fake.lowMutex.RUnlock()
	return len(fake.lowArgsForCall)
}

func (fake *FakeGoRaspberryPiIOPin) Toggle() {
	fake.toggleMutex.Lock()
	fake.toggleArgsForCall = append(fake.toggleArgsForCall, struct{}{})
	fake.recordInvocation("Toggle", []interface{}{})
	fake.toggleMutex.Unlock()
	if fake.ToggleStub != nil {
		fake.ToggleStub()
	}
}

func (fake *FakeGoRaspberryPiIOPin) ToggleCallCount() int {
	fake.toggleMutex.RLock()
	defer fake.toggleMutex.RUnlock()
	return len(fake.toggleArgsForCall)
}

func (fake *FakeGoRaspberryPiIOPin) Freq(freq int) {
	fake.freqMutex.Lock()
	fake.freqArgsForCall = append(fake.freqArgsForCall, struct {
		freq int
	}{freq})
	fake.recordInvocation("Freq", []interface{}{freq})
	fake.freqMutex.Unlock()
	if fake.FreqStub != nil {
		fake.FreqStub(freq)
	}
}

func (fake *FakeGoRaspberryPiIOPin) FreqCallCount() int {
	fake.freqMutex.RLock()
	defer fake.freqMutex.RUnlock()
	return len(fake.freqArgsForCall)
}

func (fake *FakeGoRaspberryPiIOPin) FreqArgsForCall(i int) int {
	fake.freqMutex.RLock()
	defer fake.freqMutex.RUnlock()
	return fake.freqArgsForCall[i].freq
}

func (fake *FakeGoRaspberryPiIOPin) DutyCycle(dutyLen uint32, cycleLen uint32) {
	fake.dutyCycleMutex.Lock()
	fake.dutyCycleArgsForCall = append(fake.dutyCycleArgsForCall, struct {
		dutyLen  uint32
		cycleLen uint32
	}{dutyLen, cycleLen})
	fake.recordInvocation("DutyCycle", []interface{}{dutyLen, cycleLen})
	fake.dutyCycleMutex.Unlock()
	if fake.DutyCycleStub != nil {
		fake.DutyCycleStub(dutyLen, cycleLen)
	}
}

func (fake *FakeGoRaspberryPiIOPin) DutyCycleCallCount() int {
	fake.dutyCycleMutex.RLock()
	defer fake.dutyCycleMutex.RUnlock()
	return len(fake.dutyCycleArgsForCall)
}

func (fake *FakeGoRaspberryPiIOPin) DutyCycleArgsForCall(i int) (uint32, uint32) {
	fake.dutyCycleMutex.RLock()
	defer fake.dutyCycleMutex.RUnlock()
	return fake.dutyCycleArgsForCall[i].dutyLen, fake.dutyCycleArgsForCall[i].cycleLen
}

func (fake *FakeGoRaspberryPiIOPin) Mode(mode gorpio.Mode) {
	fake.modeMutex.Lock()
	fake.modeArgsForCall = append(fake.modeArgsForCall, struct {
		mode gorpio.Mode
	}{mode})
	fake.recordInvocation("Mode", []interface{}{mode})
	fake.modeMutex.Unlock()
	if fake.ModeStub != nil {
		fake.ModeStub(mode)
	}
}

func (fake *FakeGoRaspberryPiIOPin) ModeCallCount() int {
	fake.modeMutex.RLock()
	defer fake.modeMutex.RUnlock()
	return len(fake.modeArgsForCall)
}

func (fake *FakeGoRaspberryPiIOPin) ModeArgsForCall(i int) gorpio.Mode {
	fake.modeMutex.RLock()
	defer fake.modeMutex.RUnlock()
	return fake.modeArgsForCall[i].mode
}

func (fake *FakeGoRaspberryPiIOPin) Write(state gorpio.State) {
	fake.writeMutex.Lock()
	fake.writeArgsForCall = append(fake.writeArgsForCall, struct {
		state gorpio.State
	}{state})
	fake.recordInvocation("Write", []interface{}{state})
	fake.writeMutex.Unlock()
	if fake.WriteStub != nil {
		fake.WriteStub(state)
	}
}

func (fake *FakeGoRaspberryPiIOPin) WriteCallCount() int {
	fake.writeMutex.RLock()
	defer fake.writeMutex.RUnlock()
	return len(fake.writeArgsForCall)
}

func (fake *FakeGoRaspberryPiIOPin) WriteArgsForCall(i int) gorpio.State {
	fake.writeMutex.RLock()
	defer fake.writeMutex.RUnlock()
	return fake.writeArgsForCall[i].state
}

func (fake *FakeGoRaspberryPiIOPin) Read() gorpio.State {
	fake.readMutex.Lock()
	ret, specificReturn := fake.readReturnsOnCall[len(fake.readArgsForCall)]
	fake.readArgsForCall = append(fake.readArgsForCall, struct{}{})
	fake.recordInvocation("Read", []interface{}{})
	fake.readMutex.Unlock()
	if fake.ReadStub != nil {
		return fake.ReadStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.readReturns.result1
}

func (fake *FakeGoRaspberryPiIOPin) ReadCallCount() int {
	fake.readMutex.RLock()
	defer fake.readMutex.RUnlock()
	return len(fake.readArgsForCall)
}

func (fake *FakeGoRaspberryPiIOPin) ReadReturns(result1 gorpio.State) {
	fake.ReadStub = nil
	fake.readReturns = struct {
		result1 gorpio.State
	}{result1}
}

func (fake *FakeGoRaspberryPiIOPin) ReadReturnsOnCall(i int, result1 gorpio.State) {
	fake.ReadStub = nil
	if fake.readReturnsOnCall == nil {
		fake.readReturnsOnCall = make(map[int]struct {
			result1 gorpio.State
		})
	}
	fake.readReturnsOnCall[i] = struct {
		result1 gorpio.State
	}{result1}
}

func (fake *FakeGoRaspberryPiIOPin) Pull(pull gorpio.Pull) {
	fake.pullMutex.Lock()
	fake.pullArgsForCall = append(fake.pullArgsForCall, struct {
		pull gorpio.Pull
	}{pull})
	fake.recordInvocation("Pull", []interface{}{pull})
	fake.pullMutex.Unlock()
	if fake.PullStub != nil {
		fake.PullStub(pull)
	}
}

func (fake *FakeGoRaspberryPiIOPin) PullCallCount() int {
	fake.pullMutex.RLock()
	defer fake.pullMutex.RUnlock()
	return len(fake.pullArgsForCall)
}

func (fake *FakeGoRaspberryPiIOPin) PullArgsForCall(i int) gorpio.Pull {
	fake.pullMutex.RLock()
	defer fake.pullMutex.RUnlock()
	return fake.pullArgsForCall[i].pull
}

func (fake *FakeGoRaspberryPiIOPin) PullUp() {
	fake.pullUpMutex.Lock()
	fake.pullUpArgsForCall = append(fake.pullUpArgsForCall, struct{}{})
	fake.recordInvocation("PullUp", []interface{}{})
	fake.pullUpMutex.Unlock()
	if fake.PullUpStub != nil {
		fake.PullUpStub()
	}
}

func (fake *FakeGoRaspberryPiIOPin) PullUpCallCount() int {
	fake.pullUpMutex.RLock()
	defer fake.pullUpMutex.RUnlock()
	return len(fake.pullUpArgsForCall)
}

func (fake *FakeGoRaspberryPiIOPin) PullOff() {
	fake.pullOffMutex.Lock()
	fake.pullOffArgsForCall = append(fake.pullOffArgsForCall, struct{}{})
	fake.recordInvocation("PullOff", []interface{}{})
	fake.pullOffMutex.Unlock()
	if fake.PullOffStub != nil {
		fake.PullOffStub()
	}
}

func (fake *FakeGoRaspberryPiIOPin) PullOffCallCount() int {
	fake.pullOffMutex.RLock()
	defer fake.pullOffMutex.RUnlock()
	return len(fake.pullOffArgsForCall)
}

func (fake *FakeGoRaspberryPiIOPin) Detect(edge gorpio.Edge) {
	fake.detectMutex.Lock()
	fake.detectArgsForCall = append(fake.detectArgsForCall, struct {
		edge gorpio.Edge
	}{edge})
	fake.recordInvocation("Detect", []interface{}{edge})
	fake.detectMutex.Unlock()
	if fake.DetectStub != nil {
		fake.DetectStub(edge)
	}
}

func (fake *FakeGoRaspberryPiIOPin) DetectCallCount() int {
	fake.detectMutex.RLock()
	defer fake.detectMutex.RUnlock()
	return len(fake.detectArgsForCall)
}

func (fake *FakeGoRaspberryPiIOPin) DetectArgsForCall(i int) gorpio.Edge {
	fake.detectMutex.RLock()
	defer fake.detectMutex.RUnlock()
	return fake.detectArgsForCall[i].edge
}

func (fake *FakeGoRaspberryPiIOPin) EdgeDetected() bool {
	fake.edgeDetectedMutex.Lock()
	ret, specificReturn := fake.edgeDetectedReturnsOnCall[len(fake.edgeDetectedArgsForCall)]
	fake.edgeDetectedArgsForCall = append(fake.edgeDetectedArgsForCall, struct{}{})
	fake.recordInvocation("EdgeDetected", []interface{}{})
	fake.edgeDetectedMutex.Unlock()
	if fake.EdgeDetectedStub != nil {
		return fake.EdgeDetectedStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.edgeDetectedReturns.result1
}

func (fake *FakeGoRaspberryPiIOPin) EdgeDetectedCallCount() int {
	fake.edgeDetectedMutex.RLock()
	defer fake.edgeDetectedMutex.RUnlock()
	return len(fake.edgeDetectedArgsForCall)
}

func (fake *FakeGoRaspberryPiIOPin) EdgeDetectedReturns(result1 bool) {
	fake.EdgeDetectedStub = nil
	fake.edgeDetectedReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeGoRaspberryPiIOPin) EdgeDetectedReturnsOnCall(i int, result1 bool) {
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

func (fake *FakeGoRaspberryPiIOPin) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.inputMutex.RLock()
	defer fake.inputMutex.RUnlock()
	fake.outputMutex.RLock()
	defer fake.outputMutex.RUnlock()
	fake.clockMutex.RLock()
	defer fake.clockMutex.RUnlock()
	fake.pwmMutex.RLock()
	defer fake.pwmMutex.RUnlock()
	fake.highMutex.RLock()
	defer fake.highMutex.RUnlock()
	fake.lowMutex.RLock()
	defer fake.lowMutex.RUnlock()
	fake.toggleMutex.RLock()
	defer fake.toggleMutex.RUnlock()
	fake.freqMutex.RLock()
	defer fake.freqMutex.RUnlock()
	fake.dutyCycleMutex.RLock()
	defer fake.dutyCycleMutex.RUnlock()
	fake.modeMutex.RLock()
	defer fake.modeMutex.RUnlock()
	fake.writeMutex.RLock()
	defer fake.writeMutex.RUnlock()
	fake.readMutex.RLock()
	defer fake.readMutex.RUnlock()
	fake.pullMutex.RLock()
	defer fake.pullMutex.RUnlock()
	fake.pullUpMutex.RLock()
	defer fake.pullUpMutex.RUnlock()
	fake.pullOffMutex.RLock()
	defer fake.pullOffMutex.RUnlock()
	fake.detectMutex.RLock()
	defer fake.detectMutex.RUnlock()
	fake.edgeDetectedMutex.RLock()
	defer fake.edgeDetectedMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeGoRaspberryPiIOPin) recordInvocation(key string, args []interface{}) {
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

var _ monitor.GoRaspberryPiIOPin = new(FakeGoRaspberryPiIOPin)