package temperature_monitor_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestTemperatureMonitor(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Temperature Monitor Unit Test Suite")
}
