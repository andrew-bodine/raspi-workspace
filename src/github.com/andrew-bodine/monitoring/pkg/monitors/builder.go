package monitors

import (
	"go.uber.org/zap"
)

// ConfigWithFlagsAndLoggerConstructor declares an interface for concrete
// monitor packages to implement specific config constructors for their needs,
// and provides a consistent way for configs to be created regardless of the
// monitor implementation.
type ConfigWithFlagsAndLoggerConstructor func(logger *zap.Logger) (config interface{}, err error)

// MonitorBuilder declares an interface for concrete monitor packages to
// implement specific monitor constructors for their needs and provides a
// consistent way for monitors to be be created regardless of the monitor
// implementation.
type MonitorBuilder func(config interface{}) (monitor Monitor, err error)

type ConfigAndBuilderWrapper struct {
	Logger            *zap.Logger
	ConfigConstructor ConfigWithFlagsAndLoggerConstructor
	MonitorBuilder    MonitorBuilder
}

func BuildAndRunMonitor(wrapper *ConfigAndBuilderWrapper) error {
	config, err := wrapper.ConfigConstructor(wrapper.Logger)

	if err != nil || config == nil {
		return err
	}

	monitor, err := wrapper.MonitorBuilder(config)

	if err != nil {
		return err
	}

	stopCh := SetupSignalHandler()

	if err = monitor.Run(stopCh); err != nil {
		return err
	}

	return nil
}
