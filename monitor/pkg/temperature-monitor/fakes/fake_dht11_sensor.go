// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"
	"time"

	temperature_monitor "github.com/andrew-bodine/raspi/monitor/pkg/temperature-monitor"
	rpio "github.com/stianeikeland/go-rpio"
)

type FakeDHT11Sensor struct {
	BitsToBytesStub        func([]bool) []byte
	bitsToBytesMutex       sync.RWMutex
	bitsToBytesArgsForCall []struct {
		arg1 []bool
	}
	bitsToBytesReturns struct {
		result1 []byte
	}
	bitsToBytesReturnsOnCall map[int]struct {
		result1 []byte
	}
	CalculateBitsStub        func([]int) []bool
	calculateBitsMutex       sync.RWMutex
	calculateBitsArgsForCall []struct {
		arg1 []int
	}
	calculateBitsReturns struct {
		result1 []bool
	}
	calculateBitsReturnsOnCall map[int]struct {
		result1 []bool
	}
	CalculateChecksumStub        func([]byte) byte
	calculateChecksumMutex       sync.RWMutex
	calculateChecksumArgsForCall []struct {
		arg1 []byte
	}
	calculateChecksumReturns struct {
		result1 byte
	}
	calculateChecksumReturnsOnCall map[int]struct {
		result1 byte
	}
	CollectInputStub        func() []rpio.State
	collectInputMutex       sync.RWMutex
	collectInputArgsForCall []struct {
	}
	collectInputReturns struct {
		result1 []rpio.State
	}
	collectInputReturnsOnCall map[int]struct {
		result1 []rpio.State
	}
	ParseDataPullUpLengthsStub        func([]rpio.State) []int
	parseDataPullUpLengthsMutex       sync.RWMutex
	parseDataPullUpLengthsArgsForCall []struct {
		arg1 []rpio.State
	}
	parseDataPullUpLengthsReturns struct {
		result1 []int
	}
	parseDataPullUpLengthsReturnsOnCall map[int]struct {
		result1 []int
	}
	ReadStub        func() *temperature_monitor.DHT11Result
	readMutex       sync.RWMutex
	readArgsForCall []struct {
	}
	readReturns struct {
		result1 *temperature_monitor.DHT11Result
	}
	readReturnsOnCall map[int]struct {
		result1 *temperature_monitor.DHT11Result
	}
	SendAndSleepStub        func(rpio.State, time.Duration)
	sendAndSleepMutex       sync.RWMutex
	sendAndSleepArgsForCall []struct {
		arg1 rpio.State
		arg2 time.Duration
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeDHT11Sensor) BitsToBytes(arg1 []bool) []byte {
	var arg1Copy []bool
	if arg1 != nil {
		arg1Copy = make([]bool, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.bitsToBytesMutex.Lock()
	ret, specificReturn := fake.bitsToBytesReturnsOnCall[len(fake.bitsToBytesArgsForCall)]
	fake.bitsToBytesArgsForCall = append(fake.bitsToBytesArgsForCall, struct {
		arg1 []bool
	}{arg1Copy})
	stub := fake.BitsToBytesStub
	fakeReturns := fake.bitsToBytesReturns
	fake.recordInvocation("BitsToBytes", []interface{}{arg1Copy})
	fake.bitsToBytesMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeDHT11Sensor) BitsToBytesCallCount() int {
	fake.bitsToBytesMutex.RLock()
	defer fake.bitsToBytesMutex.RUnlock()
	return len(fake.bitsToBytesArgsForCall)
}

func (fake *FakeDHT11Sensor) BitsToBytesCalls(stub func([]bool) []byte) {
	fake.bitsToBytesMutex.Lock()
	defer fake.bitsToBytesMutex.Unlock()
	fake.BitsToBytesStub = stub
}

func (fake *FakeDHT11Sensor) BitsToBytesArgsForCall(i int) []bool {
	fake.bitsToBytesMutex.RLock()
	defer fake.bitsToBytesMutex.RUnlock()
	argsForCall := fake.bitsToBytesArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeDHT11Sensor) BitsToBytesReturns(result1 []byte) {
	fake.bitsToBytesMutex.Lock()
	defer fake.bitsToBytesMutex.Unlock()
	fake.BitsToBytesStub = nil
	fake.bitsToBytesReturns = struct {
		result1 []byte
	}{result1}
}

func (fake *FakeDHT11Sensor) BitsToBytesReturnsOnCall(i int, result1 []byte) {
	fake.bitsToBytesMutex.Lock()
	defer fake.bitsToBytesMutex.Unlock()
	fake.BitsToBytesStub = nil
	if fake.bitsToBytesReturnsOnCall == nil {
		fake.bitsToBytesReturnsOnCall = make(map[int]struct {
			result1 []byte
		})
	}
	fake.bitsToBytesReturnsOnCall[i] = struct {
		result1 []byte
	}{result1}
}

func (fake *FakeDHT11Sensor) CalculateBits(arg1 []int) []bool {
	var arg1Copy []int
	if arg1 != nil {
		arg1Copy = make([]int, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.calculateBitsMutex.Lock()
	ret, specificReturn := fake.calculateBitsReturnsOnCall[len(fake.calculateBitsArgsForCall)]
	fake.calculateBitsArgsForCall = append(fake.calculateBitsArgsForCall, struct {
		arg1 []int
	}{arg1Copy})
	stub := fake.CalculateBitsStub
	fakeReturns := fake.calculateBitsReturns
	fake.recordInvocation("CalculateBits", []interface{}{arg1Copy})
	fake.calculateBitsMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeDHT11Sensor) CalculateBitsCallCount() int {
	fake.calculateBitsMutex.RLock()
	defer fake.calculateBitsMutex.RUnlock()
	return len(fake.calculateBitsArgsForCall)
}

func (fake *FakeDHT11Sensor) CalculateBitsCalls(stub func([]int) []bool) {
	fake.calculateBitsMutex.Lock()
	defer fake.calculateBitsMutex.Unlock()
	fake.CalculateBitsStub = stub
}

func (fake *FakeDHT11Sensor) CalculateBitsArgsForCall(i int) []int {
	fake.calculateBitsMutex.RLock()
	defer fake.calculateBitsMutex.RUnlock()
	argsForCall := fake.calculateBitsArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeDHT11Sensor) CalculateBitsReturns(result1 []bool) {
	fake.calculateBitsMutex.Lock()
	defer fake.calculateBitsMutex.Unlock()
	fake.CalculateBitsStub = nil
	fake.calculateBitsReturns = struct {
		result1 []bool
	}{result1}
}

func (fake *FakeDHT11Sensor) CalculateBitsReturnsOnCall(i int, result1 []bool) {
	fake.calculateBitsMutex.Lock()
	defer fake.calculateBitsMutex.Unlock()
	fake.CalculateBitsStub = nil
	if fake.calculateBitsReturnsOnCall == nil {
		fake.calculateBitsReturnsOnCall = make(map[int]struct {
			result1 []bool
		})
	}
	fake.calculateBitsReturnsOnCall[i] = struct {
		result1 []bool
	}{result1}
}

func (fake *FakeDHT11Sensor) CalculateChecksum(arg1 []byte) byte {
	var arg1Copy []byte
	if arg1 != nil {
		arg1Copy = make([]byte, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.calculateChecksumMutex.Lock()
	ret, specificReturn := fake.calculateChecksumReturnsOnCall[len(fake.calculateChecksumArgsForCall)]
	fake.calculateChecksumArgsForCall = append(fake.calculateChecksumArgsForCall, struct {
		arg1 []byte
	}{arg1Copy})
	stub := fake.CalculateChecksumStub
	fakeReturns := fake.calculateChecksumReturns
	fake.recordInvocation("CalculateChecksum", []interface{}{arg1Copy})
	fake.calculateChecksumMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeDHT11Sensor) CalculateChecksumCallCount() int {
	fake.calculateChecksumMutex.RLock()
	defer fake.calculateChecksumMutex.RUnlock()
	return len(fake.calculateChecksumArgsForCall)
}

func (fake *FakeDHT11Sensor) CalculateChecksumCalls(stub func([]byte) byte) {
	fake.calculateChecksumMutex.Lock()
	defer fake.calculateChecksumMutex.Unlock()
	fake.CalculateChecksumStub = stub
}

func (fake *FakeDHT11Sensor) CalculateChecksumArgsForCall(i int) []byte {
	fake.calculateChecksumMutex.RLock()
	defer fake.calculateChecksumMutex.RUnlock()
	argsForCall := fake.calculateChecksumArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeDHT11Sensor) CalculateChecksumReturns(result1 byte) {
	fake.calculateChecksumMutex.Lock()
	defer fake.calculateChecksumMutex.Unlock()
	fake.CalculateChecksumStub = nil
	fake.calculateChecksumReturns = struct {
		result1 byte
	}{result1}
}

func (fake *FakeDHT11Sensor) CalculateChecksumReturnsOnCall(i int, result1 byte) {
	fake.calculateChecksumMutex.Lock()
	defer fake.calculateChecksumMutex.Unlock()
	fake.CalculateChecksumStub = nil
	if fake.calculateChecksumReturnsOnCall == nil {
		fake.calculateChecksumReturnsOnCall = make(map[int]struct {
			result1 byte
		})
	}
	fake.calculateChecksumReturnsOnCall[i] = struct {
		result1 byte
	}{result1}
}

func (fake *FakeDHT11Sensor) CollectInput() []rpio.State {
	fake.collectInputMutex.Lock()
	ret, specificReturn := fake.collectInputReturnsOnCall[len(fake.collectInputArgsForCall)]
	fake.collectInputArgsForCall = append(fake.collectInputArgsForCall, struct {
	}{})
	stub := fake.CollectInputStub
	fakeReturns := fake.collectInputReturns
	fake.recordInvocation("CollectInput", []interface{}{})
	fake.collectInputMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeDHT11Sensor) CollectInputCallCount() int {
	fake.collectInputMutex.RLock()
	defer fake.collectInputMutex.RUnlock()
	return len(fake.collectInputArgsForCall)
}

func (fake *FakeDHT11Sensor) CollectInputCalls(stub func() []rpio.State) {
	fake.collectInputMutex.Lock()
	defer fake.collectInputMutex.Unlock()
	fake.CollectInputStub = stub
}

func (fake *FakeDHT11Sensor) CollectInputReturns(result1 []rpio.State) {
	fake.collectInputMutex.Lock()
	defer fake.collectInputMutex.Unlock()
	fake.CollectInputStub = nil
	fake.collectInputReturns = struct {
		result1 []rpio.State
	}{result1}
}

func (fake *FakeDHT11Sensor) CollectInputReturnsOnCall(i int, result1 []rpio.State) {
	fake.collectInputMutex.Lock()
	defer fake.collectInputMutex.Unlock()
	fake.CollectInputStub = nil
	if fake.collectInputReturnsOnCall == nil {
		fake.collectInputReturnsOnCall = make(map[int]struct {
			result1 []rpio.State
		})
	}
	fake.collectInputReturnsOnCall[i] = struct {
		result1 []rpio.State
	}{result1}
}

func (fake *FakeDHT11Sensor) ParseDataPullUpLengths(arg1 []rpio.State) []int {
	var arg1Copy []rpio.State
	if arg1 != nil {
		arg1Copy = make([]rpio.State, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.parseDataPullUpLengthsMutex.Lock()
	ret, specificReturn := fake.parseDataPullUpLengthsReturnsOnCall[len(fake.parseDataPullUpLengthsArgsForCall)]
	fake.parseDataPullUpLengthsArgsForCall = append(fake.parseDataPullUpLengthsArgsForCall, struct {
		arg1 []rpio.State
	}{arg1Copy})
	stub := fake.ParseDataPullUpLengthsStub
	fakeReturns := fake.parseDataPullUpLengthsReturns
	fake.recordInvocation("ParseDataPullUpLengths", []interface{}{arg1Copy})
	fake.parseDataPullUpLengthsMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeDHT11Sensor) ParseDataPullUpLengthsCallCount() int {
	fake.parseDataPullUpLengthsMutex.RLock()
	defer fake.parseDataPullUpLengthsMutex.RUnlock()
	return len(fake.parseDataPullUpLengthsArgsForCall)
}

func (fake *FakeDHT11Sensor) ParseDataPullUpLengthsCalls(stub func([]rpio.State) []int) {
	fake.parseDataPullUpLengthsMutex.Lock()
	defer fake.parseDataPullUpLengthsMutex.Unlock()
	fake.ParseDataPullUpLengthsStub = stub
}

func (fake *FakeDHT11Sensor) ParseDataPullUpLengthsArgsForCall(i int) []rpio.State {
	fake.parseDataPullUpLengthsMutex.RLock()
	defer fake.parseDataPullUpLengthsMutex.RUnlock()
	argsForCall := fake.parseDataPullUpLengthsArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeDHT11Sensor) ParseDataPullUpLengthsReturns(result1 []int) {
	fake.parseDataPullUpLengthsMutex.Lock()
	defer fake.parseDataPullUpLengthsMutex.Unlock()
	fake.ParseDataPullUpLengthsStub = nil
	fake.parseDataPullUpLengthsReturns = struct {
		result1 []int
	}{result1}
}

func (fake *FakeDHT11Sensor) ParseDataPullUpLengthsReturnsOnCall(i int, result1 []int) {
	fake.parseDataPullUpLengthsMutex.Lock()
	defer fake.parseDataPullUpLengthsMutex.Unlock()
	fake.ParseDataPullUpLengthsStub = nil
	if fake.parseDataPullUpLengthsReturnsOnCall == nil {
		fake.parseDataPullUpLengthsReturnsOnCall = make(map[int]struct {
			result1 []int
		})
	}
	fake.parseDataPullUpLengthsReturnsOnCall[i] = struct {
		result1 []int
	}{result1}
}

func (fake *FakeDHT11Sensor) Read() *temperature_monitor.DHT11Result {
	fake.readMutex.Lock()
	ret, specificReturn := fake.readReturnsOnCall[len(fake.readArgsForCall)]
	fake.readArgsForCall = append(fake.readArgsForCall, struct {
	}{})
	stub := fake.ReadStub
	fakeReturns := fake.readReturns
	fake.recordInvocation("Read", []interface{}{})
	fake.readMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeDHT11Sensor) ReadCallCount() int {
	fake.readMutex.RLock()
	defer fake.readMutex.RUnlock()
	return len(fake.readArgsForCall)
}

func (fake *FakeDHT11Sensor) ReadCalls(stub func() *temperature_monitor.DHT11Result) {
	fake.readMutex.Lock()
	defer fake.readMutex.Unlock()
	fake.ReadStub = stub
}

func (fake *FakeDHT11Sensor) ReadReturns(result1 *temperature_monitor.DHT11Result) {
	fake.readMutex.Lock()
	defer fake.readMutex.Unlock()
	fake.ReadStub = nil
	fake.readReturns = struct {
		result1 *temperature_monitor.DHT11Result
	}{result1}
}

func (fake *FakeDHT11Sensor) ReadReturnsOnCall(i int, result1 *temperature_monitor.DHT11Result) {
	fake.readMutex.Lock()
	defer fake.readMutex.Unlock()
	fake.ReadStub = nil
	if fake.readReturnsOnCall == nil {
		fake.readReturnsOnCall = make(map[int]struct {
			result1 *temperature_monitor.DHT11Result
		})
	}
	fake.readReturnsOnCall[i] = struct {
		result1 *temperature_monitor.DHT11Result
	}{result1}
}

func (fake *FakeDHT11Sensor) SendAndSleep(arg1 rpio.State, arg2 time.Duration) {
	fake.sendAndSleepMutex.Lock()
	fake.sendAndSleepArgsForCall = append(fake.sendAndSleepArgsForCall, struct {
		arg1 rpio.State
		arg2 time.Duration
	}{arg1, arg2})
	stub := fake.SendAndSleepStub
	fake.recordInvocation("SendAndSleep", []interface{}{arg1, arg2})
	fake.sendAndSleepMutex.Unlock()
	if stub != nil {
		fake.SendAndSleepStub(arg1, arg2)
	}
}

func (fake *FakeDHT11Sensor) SendAndSleepCallCount() int {
	fake.sendAndSleepMutex.RLock()
	defer fake.sendAndSleepMutex.RUnlock()
	return len(fake.sendAndSleepArgsForCall)
}

func (fake *FakeDHT11Sensor) SendAndSleepCalls(stub func(rpio.State, time.Duration)) {
	fake.sendAndSleepMutex.Lock()
	defer fake.sendAndSleepMutex.Unlock()
	fake.SendAndSleepStub = stub
}

func (fake *FakeDHT11Sensor) SendAndSleepArgsForCall(i int) (rpio.State, time.Duration) {
	fake.sendAndSleepMutex.RLock()
	defer fake.sendAndSleepMutex.RUnlock()
	argsForCall := fake.sendAndSleepArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeDHT11Sensor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.bitsToBytesMutex.RLock()
	defer fake.bitsToBytesMutex.RUnlock()
	fake.calculateBitsMutex.RLock()
	defer fake.calculateBitsMutex.RUnlock()
	fake.calculateChecksumMutex.RLock()
	defer fake.calculateChecksumMutex.RUnlock()
	fake.collectInputMutex.RLock()
	defer fake.collectInputMutex.RUnlock()
	fake.parseDataPullUpLengthsMutex.RLock()
	defer fake.parseDataPullUpLengthsMutex.RUnlock()
	fake.readMutex.RLock()
	defer fake.readMutex.RUnlock()
	fake.sendAndSleepMutex.RLock()
	defer fake.sendAndSleepMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeDHT11Sensor) recordInvocation(key string, args []interface{}) {
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

var _ temperature_monitor.DHT11Sensor = new(FakeDHT11Sensor)
