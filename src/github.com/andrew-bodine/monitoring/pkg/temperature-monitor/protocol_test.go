package temperature_monitor_test

import (
	"math/rand"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	monitor "github.com/andrew-bodine/monitoring/pkg/temperature-monitor"
)

var _ = Describe("Temperature Monitor Protocol Unit Tests", func() {
	var (
		dht11 *monitor.DHT11
	)

	BeforeEach(func() {
		dht11 = &monitor.DHT11{}
	})

	Context("Read", func() {
		// TODO:
	})

	Context("SendAndSleep", func() {
		// TODO:
	})

	Context("CollectInput", func() {
		// TODO:
	})

	Context("ParseDataPullUpLengths", func() {
		// TODO:
	})

	Context("CalculateBits", func() {
		var (
			input  []int
			output []bool
		)

		BeforeEach(func() {
			input = []int{}

			rand.Seed(time.Now().Unix())

			// Randomly generate pull up lengths.
			for i := 0; i <= 31; i++ {
				input = append(input, rand.Intn(1000))
			}
		})

		JustBeforeEach(func() {
			output = dht11.CalculateBits(input)
		})

		It("calculates the correct bits from the provided pull up lengths", func() {
			Expect(len(output)).To(Equal(32))
		})
	})

	Context("BitsToBytes", func() {
		var (
			input  []bool
			output []byte
		)

		BeforeEach(func() {
			input = []bool{false, true, true, false, false, true, false, true}
		})

		JustBeforeEach(func() {
			output = dht11.BitsToBytes(input)
		})

		It("converts the bits the to correct bytes", func() {
			Expect(output).To(Equal([]byte{byte(101)}))
		})
	})

	Context("CalculateChecksum", func() {
		var (
			input  []byte
			output byte
		)

		BeforeEach(func() {
			input = []byte("world")
		})

		JustBeforeEach(func() {
			output = dht11.CalculateChecksum(input)
		})

		It("calculates the correct checksum", func() {
			Expect(output).To(Equal(byte(196)))
		})
	})
})
