package temperature_monitor

import (
	"time"

	rpio "github.com/stianeikeland/go-rpio"
)

// http://kookye.com/2017/06/01/how-to-set-up-temperature-and-humidity-sensor-dht11-on-raspberry-pi-2/
// http://osoyoo.com/driver/dht11.py

func (dht11 *DHT11) ReadCurrentTempAndHumidity() *DHT11Result {
	result := &DHT11Result{}

	var protocolState State

	// The default status of the DATA pin is high.
	//
	// When the communication between MCU and DHT11 starts, MCU will pull down
	// the DATA pin for least 18ms. This is called “Start Signal” and it is to
	// ensure DHT11 has detected the signal from MCU.
	dht11.Pin.Output()
	dht11.Pin.PullDown()
	protocolState = StateInitPullDown
	// fmt.Println(">>>> protocolState = StateInitPullDown")
	time.Sleep(time.Millisecond * 18)

	// Then MCU will pull up DATA pin for 20-40us to wait for DHT11’s response.
	dht11.Pin.PullUp()
	protocolState = StateInitPullUp
	// fmt.Println(">>>> protocolState = StateInitPullUp")
	dht11.Pin.Input()
	deadlineTimer := time.NewTimer(time.Microsecond * 40)
	for protocolState != StateDataFirstPullDown {
		select {
		case _ = <-deadlineTimer.C:
			result.Error = ErrorHandshake
			result.Message = "Didn't receive pull down response from DHT11 sensor during initial protocol handshake."
			return result
		default:
			if dht11.Pin.Read() == rpio.Low {
				protocolState = StateDataFirstPullDown
				// fmt.Println(">>>> protocolState = StateDataFirstPullDown")
				continue
			}
			// time.Sleep(time.Microsecond)
		}
	}

	// Once DHT11 detects the start signal, it will pull down the DATA pin as
	// “Response Signal”, which will last 80us.
	deadlineTimer = time.NewTimer(time.Microsecond * 80)
	for protocolState != StateDataPullUp {
		select {
		case _ = <-deadlineTimer.C:
			result.Error = ErrorHandshake
			result.Message = "Didn't receive pull up response from DHT11 sensor during initial protocl handshake."
			return result
		default:
			if dht11.Pin.Read() == rpio.High {
				protocolState = StateDataPullUp
				// fmt.Println(">>>> protocolState = StateDataPullUp")
				continue
			}
			// time.Sleep(time.Microsecond)
		}
	}

	// Then DHT11 will pull up the DATA pin for 80us, and prepare for data sending.

	var mostRecentPullUpTime time.Time
	bits := make([]bool, 40)
	currBit := 0

	unchangedCount := 0
	maxUnchanged := 100

	// During the data transition, every bit of data begins with the 50us
	// low-voltage-level and ends with a high-voltage-level signal. The length
	// of the high-voltage-level signal decides whether the bit is “0” or “1”.
	//
	// Data bit “0” has 26-28us high-voltage length, while data bit “1” has
	// 70us high-voltage length.
	for currBit < 40 {
		if unchangedCount >= maxUnchanged {
			// fmt.Println(">>>> comms process terminated unexpectedly")
			return &DHT11Result{Error: ErrorMissingData, Message: "Comms process terminated unexpectedly."}
			break
		}

		switch dht11.Pin.Read() {
		case rpio.High:
			if protocolState == StateDataPullDown {
				mostRecentPullUpTime = time.Now()
				protocolState = StateDataPullUp
				// fmt.Println(">>>> protocolState = StateDataPullUp")
				unchangedCount = 0
			} else {
				unchangedCount++
			}
		case rpio.Low:
			if protocolState == StateDataPullUp {
				if !mostRecentPullUpTime.IsZero() {
					elapsedTime := time.Now().Sub(mostRecentPullUpTime)
					// if elapsedTime > (time.Microsecond * 28) {
					if elapsedTime > (time.Microsecond * 50) {
						bits[currBit] = true
						// fmt.Println("Got a 1")
						currBit++
					} else {
						bits[currBit] = false
						// fmt.Println("Got a 0")
						currBit++
					}
				}
				protocolState = StateDataPullDown
				// fmt.Println(">>>> protocolState = StateDataPullDown")
				unchangedCount = 0
			} else {
				unchangedCount++
			}
		}

		time.Sleep(time.Microsecond)
	}

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
