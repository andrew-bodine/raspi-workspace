# temperature-monitor

## Monitor Features
- Read temperature and humidity data from a DHT11 sensor on demand.
- Logs temperature and humidity readings, only when they change.

## DHT11 Sensor Protocol
The GPIO pin that is wired to the DHT11 sensor can either be in LOW or HIGH
state, and is in HIGH state by default. The following sequence of steps
summarizes the transmission protocol for sampling the sensors current readings.
For more information see the [Kookye Tutorial](http://kookye.com/2017/06/01/how-to-set-up-temperature-and-humidity-sensor-dht11-on-raspberry-pi-2)

1. Control process pulls GPIO pin down for 18 milliseconds.
1. Control process pulls GPIO pin up for 20-40 microseconds.
1. DHT11 sensor pulls GPIO pin down for 80 microseconds.
1. DHT11 sensor pulls GPIO pin up for 80 microseconds.
1. Data transmission loop until 40 bits read:
   1. DHT11 sensor pulls GPIO pin down for 50 microseconds.
   1. DHT11 sensor pulls GPIO pin up for either 26-28 microseconds or 70
      microseconds, indicating a 0 or 1 bit respectively.

## Roadmap
- Monitor should be able to record the protocol data for a real device over a specific time period.
- Monitor should be able to calculate/adjust protocol timing variables at runtime.
- Monitor should ship log "events" to cloud storage for processing/visualization.
