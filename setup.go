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

	// Logger is an alias for logrus's standard logger
	// struct. This allows it to be replaced at runtime or during
	// tests.
	Logger = logrus.StandardLogger()
)

// ConfigureRoot will use the provided configuration to setup the root
// logrus logger.
func ConfigureRoot(config *Config) error {
	if config.Level == "" || config.HookLevel == "" {
		return ErrLevelNotProvided
	}

	if config.Level == DisabledLevel {
		Logger.Out = ioutil.Discard

		// We set the level to panic here because even though we're
		// discarding output logrus will still log.
		Logger.SetLevel(logrus.PanicLevel)
	}

	if config.Level != DisabledLevel {
		// Setup the logger's root level.
		level, err := logrus.ParseLevel(config.Level)
		if err != nil {
			return err
		}
		Logger.SetLevel(level)
	}

	if config.HookLevel != DisabledLevel {
		level, err := logrus.ParseLevel(config.HookLevel)
		if err != nil {
			return err
		}
		Logger.Hooks.Add(
			NewCallerHook(
				false, config.HookStackLevel,
				config.CallerHookField, level))
		return nil
	}

	return nil
}
