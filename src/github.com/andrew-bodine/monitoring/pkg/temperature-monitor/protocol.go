package temperature_monitor

import (
	"fmt"
	"time"

	rpio "github.com/stianeikeland/go-rpio"

	"github.com/andrew-bodine/monitoring/pkg/monitors"
)

type DHT11 struct {
	Pin monitors.GoRaspberryPiIOPin
}

type DHT11Result struct {
	Error       Error
	Message     string
	Temperature []byte
	Humidity    []byte
}

type Error string

const (
	ErrorNoError     = "NoError"
	ErrorHandshake   = "Handshake"
	ErrorMissingData = "MissingData"
	ErrorChecksum    = "Checksum"
)

type State string

const (
	StateInitPullDown      = "InitPullDown"
	StateInitPullUp        = "InitPullUp"
	StateDataFirstPullDown = "DataFirstPullDown"
	StateDataPullUp        = "DataPullUp"
	StateDataPullDown      = "DataPullDown"
)

// http://kookye.com/2017/06/01/how-to-set-up-temperature-and-humidity-sensor-dht11-on-raspberry-pi-2/
// http://osoyoo.com/driver/dht11.py

func (dht11 *DHT11) Read() *DHT11Result {
	dht11.Pin.Output()

	// Send initial high signal.
	// dht11.SendAndSleep(rpio.High, time.Millisecond*50)
	// dht11.SendAndSleep(rpio.Low, time.Millisecond*18)
	dht11.Pin.PullDown()
	time.Sleep(time.Millisecond * 18)

	// Pull down to low.
	// dht11.SendAndSleep(rpio.Low, time.Millisecond*20)
	// dht11.SendAndSleep(rpio.High, time.Millisecond*40)
	dht11.Pin.PullUp()
	time.Sleep(time.Microsecond * 10)

	// Wait for DHT11 to respond to MCU start signal.
	// time.Sleep(time.Millisecond * 160)

	// Change to input.
	dht11.Pin.Input()

	// Collect data.
	states := dht11.CollectInput()

	// Parse lengths of all data pull up periods.
	pullUpLengths := dht11.ParseDataPullUpLengths(states)

	// If bit count mismatch, return error (4 byte data + 1 byte checksum)
	if len(pullUpLengths) != 40 {
		return &DHT11Result{Error: ErrorMissingData, Message: fmt.Sprintf("pullUpLengths: %d, states: %d", len(pullUpLengths), len(states))}
	}

	// Calculate bits from lengths of the pull up periods.
	bits := dht11.CalculateBits(pullUpLengths)

	// Transform bits to bytes.
	bytes := dht11.BitsToBytes(bits)

	// Calculate and verify the checksum.
	checksum := dht11.CalculateChecksum(bytes)
	if bytes[4] != checksum {
		return &DHT11Result{Error: ErrorChecksum}
	}

	return &DHT11Result{
		Error:       ErrorNoError,
		Temperature: bytes[2:4],
		Humidity:    bytes[0:2],
	}
}

func (dht11 *DHT11) SendAndSleep(state rpio.State, sleepTime time.Duration) {
	dht11.Pin.Write(state)
	time.Sleep(sleepTime)
}

// During the data transition, every bit of data begins with the 50us
// low-voltage-level and ends with a high-voltage-level signal. The length of
// the high-voltage-level signal decides whether the bit is “0” or “1”. Data
// bit “0” has 26-28us high-voltage length, while data bit “1” has 70us
// high-voltage length.
func (dht11 *DHT11) CollectInput() []rpio.State {
	states := []rpio.State{}

	unchangedCount := 0
	maxUnchangedCount := 200

	last := rpio.High

	for unchangedCount <= maxUnchangedCount {
		state := dht11.Pin.Read()
		states = append(states, state)

		if last != state {
			unchangedCount = 0
			last = state
		} else {
			unchangedCount++
		}

		time.Sleep(time.Microsecond * 1)
	}

	return states
}

func (dht11 *DHT11) ParseDataPullUpLengths(states []rpio.State) []int {
	currState := StateInitPullDown

	pullUpPeriodLengths := []int{}
	currLength := 0

	for _, state := range states {
		currLength++

		switch currState {
		case StateInitPullDown:
			if state == rpio.Low {
				currState = StateInitPullUp
			}
			continue
		case StateInitPullUp:
			if state == rpio.High {
				currState = StateDataFirstPullDown
			}
			continue
		case StateDataFirstPullDown:
			if state == rpio.Low {
				currState = StateDataPullUp
			}
			continue
		case StateDataPullUp:
			if state == rpio.High {
				currLength = 0
				currState = StateDataPullDown
			}
			continue
		case StateDataPullDown:
			if state == rpio.Low {
				pullUpPeriodLengths = append(pullUpPeriodLengths, currLength)
				currState = StateDataPullUp
			}
			continue
		}
	}

	return pullUpPeriodLengths
}

func (dht11 *DHT11) CalculateBits(pullUpLengths []int) []bool {
	shortestPullUp := 1000
	longestPullUp := 0

	for _, length := range pullUpLengths {
		if length < shortestPullUp {
			shortestPullUp = length
		}

		if length > longestPullUp {
			longestPullUp = length
		}
	}

	// Use the halfway to determine whether the period is long or short.
	halfway := shortestPullUp + (longestPullUp-shortestPullUp)/2

	bits := []bool{}

	for _, length := range pullUpLengths {
		bit := false

		if length > halfway {
			bit = true
		}

		bits = append(bits, bit)
	}

	return bits
}

func (dht11 *DHT11) BitsToBytes(bits []bool) []byte {
	bytes := []byte{}
	byte := byte(0)

	for idx, bit := range bits {
		byte = byte << 1

		if bit {
			byte = byte | 1
		} else {
			byte = byte | 0
		}

		if (idx+1)%8 == 0 {
			bytes = append(bytes, byte)
			byte = 0
		}
	}

	return bytes
}

func (dht11 *DHT11) CalculateChecksum(payload []byte) byte {
	if len(payload) < 4 {
		return byte(0)
	}

	return payload[0] + payload[1] + payload[2] + payload[3]&255
}
