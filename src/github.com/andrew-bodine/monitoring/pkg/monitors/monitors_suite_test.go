package monitors_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestMonitors(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Monitors Unit Test Suite")
}
