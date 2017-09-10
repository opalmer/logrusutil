package logrusutil

import (
	"errors"
	"io/ioutil"

	"github.com/Sirupsen/logrus"
)

var (
	// ErrLevelNotProvided is returned when a level was not provided
	// in the config struct.
	ErrLevelNotProvided = errors.New("level not provided")

	// StandardLogger is an alias for logrus's standard logger
	// struct. This allows it to be replaced at runtime or during
	// tests.
	StandardLogger = logrus.StandardLogger()
)

// ConfigureRoot will use the provided configuration to setup the root
// logrus logger.
func ConfigureRoot(config *Config) error {
	if config.Level == "" {
		return ErrLevelNotProvided
	}

	if config.Level == DisabledLevel {
		StandardLogger.Out = ioutil.Discard

		// We set the level to panic here because even though we're
		// discarding output logrus will still log.
		StandardLogger.SetLevel(logrus.PanicLevel)
		return nil
	}

	level, err := logrus.ParseLevel(config.Level)
	if err != nil {
		return err
	}
	StandardLogger.SetLevel(level)

	return nil
}
