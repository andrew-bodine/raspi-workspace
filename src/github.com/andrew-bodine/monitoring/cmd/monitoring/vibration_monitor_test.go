package main_test

import (
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Vibration Monitor Integration Tests", func() {

	var (
		minVibrationPeriod = time.Millisecond * 100
	)

	BeforeEach(func() {

	})

	JustBeforeEach(func() {

	})

	Context("with a vibration period of greater than minVibrationPeriod", func() {

	})

	Context("with a vibration period not greater than minVibrationPeriod", func() {

	})
})
