package monitors_test

import (
	"errors"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"go.uber.org/zap"

	"github.com/andrew-bodine/raspi/monitor/pkg/monitors"
	monitorFakes "github.com/andrew-bodine/raspi/monitor/pkg/monitors/fakes"
)

var _ = Describe("BuildAndRunMonitor", func() {
	var (
		logger            *zap.Logger
		configConstructor *monitorFakes.FakeConfigWithFlagsAndLoggerConstructor
		monitorBuilder    *monitorFakes.FakeMonitorBuilder
		err               error
	)

	BeforeEach(func() {
		logger, err = zap.NewProduction()
		Expect(err).NotTo(HaveOccurred())

		configConstructor = &monitorFakes.FakeConfigWithFlagsAndLoggerConstructor{}
		monitorBuilder = &monitorFakes.FakeMonitorBuilder{}
	})

	JustBeforeEach(func() {
		err = monitors.BuildAndRunMonitor(&monitors.ConfigAndBuilderWrapper{
			Logger:            logger,
			ConfigConstructor: configConstructor.Spy,
			MonitorBuilder:    monitorBuilder.Spy,
		})
	})

	Context("without any errors", func() {
		BeforeEach(func() {
			configConstructor.Returns(struct{}{}, nil)

			monitor := &monitorFakes.FakeMonitor{}
			monitor.RunReturns(nil)

			monitorBuilder.Returns(monitor, nil)
		})

		It("builds and runs the monitor", func() {
			Expect(err).NotTo(HaveOccurred())
		})
	})

	Context("when config constructor returns an error", func() {
		var configConstructorError error

		BeforeEach(func() {
			configConstructorError = errors.New("config constructor error")
			configConstructor.Returns(nil, configConstructorError)
		})

		It("returns the error", func() {
			Expect(err).To(Equal(configConstructorError))
		})
	})

	Context("when config constructor returns a nil config", func() {
		BeforeEach(func() {
			configConstructor.Returns(nil, nil)
		})

		It("returns nil", func() {
			Expect(err).NotTo(HaveOccurred())
		})
	})

	Context("when monitor builder returns an error", func() {
		var monitorBuilderError error

		BeforeEach(func() {
			configConstructor.Returns(struct{}{}, nil)

			monitorBuilderError = errors.New("monitor builder error")
			monitorBuilder.Returns(nil, monitorBuilderError)
		})

		It("returns the error", func() {
			Expect(err).To(Equal(monitorBuilderError))
		})
	})

	Context("when monitor builder returns a nil monitor", func() {
		BeforeEach(func() {
			configConstructor.Returns(struct{}{}, nil)

			monitorBuilder.Returns(nil, nil)
		})

		It("returns nil", func() {
			Expect(err).NotTo(HaveOccurred())
		})
	})

	Context("when monitor run returns an error", func() {
		var monitorRunError error

		BeforeEach(func() {
			configConstructor.Returns(struct{}{}, nil)

			monitorRunError = errors.New("monitor run error")

			monitor := &monitorFakes.FakeMonitor{}
			monitor.RunReturns(monitorRunError)

			monitorBuilder.Returns(monitor, nil)
		})

		It("returns the error", func() {
			Expect(err).To(Equal(monitorRunError))
		})
	})
})
