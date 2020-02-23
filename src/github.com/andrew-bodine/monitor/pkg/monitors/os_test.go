package monitors_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/andrew-bodine/monitor/pkg/monitors"
)

var _ = Describe("SetupSignalHandler", func() {
	var (
		stopCh <-chan struct{}
	)

	JustBeforeEach(func() {
		stopCh = monitors.SetupSignalHandler()
	})

	It("returns a valid channel of the correct type", func() {
		Expect(stopCh).NotTo(BeNil())
	})
})
