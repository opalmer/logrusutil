package logrusutil_test

import (
	"io/ioutil"
	"testing"

	"github.com/Sirupsen/logrus"
	"github.com/opalmer/logrusutil"
)

func setup() {
	logrusutil.Logger = logrus.New()
}

func teardown() {
	logrusutil.Logger = logrus.StandardLogger()
}

func TestConfigureRoot_ErrLevelNotProvided_Level(t *testing.T) {
	setup()
	defer teardown()

	cfg := logrusutil.NewConfig()
	cfg.Level = ""
	if logrusutil.ConfigureRoot(cfg) != logrusutil.ErrLevelNotProvided {
		t.Error()
	}
}

func TestConfigureRoot_ErrLevelNotProvided_HookLevel(t *testing.T) {
	setup()
	defer teardown()

	cfg := logrusutil.NewConfig()
	cfg.HookLevel = ""
	if logrusutil.ConfigureRoot(cfg) != logrusutil.ErrLevelNotProvided {
		t.Error()
	}
}

func TestConfigureRoot_DisabledLevel(t *testing.T) {
	setup()
	defer teardown()

	cfg := logrusutil.NewConfig()
	cfg.Level = logrusutil.DisabledLevel
	if err := logrusutil.ConfigureRoot(cfg); err != nil {
		t.Error(err)
	}

	if logrusutil.Logger.Level != logrus.PanicLevel {
		t.Error()
	}

	if logrusutil.Logger.Out != ioutil.Discard {
		t.Error()
	}
}

func TestConfigureRoot_HookDisabled(t *testing.T) {
	setup()
	defer teardown()

	cfg := logrusutil.NewConfig()
	cfg.Level = logrusutil.DisabledLevel
	cfg.HookLevel = logrusutil.DisabledLevel
	if err := logrusutil.ConfigureRoot(cfg); err != nil {
		t.Error(err)
	}

	if logrusutil.Logger.Level != logrus.PanicLevel {
		t.Error()
	}

	if logrusutil.Logger.Out != ioutil.Discard {
		t.Error()
	}
}

func TestConfigureRoot_Level_ParseError(t *testing.T) {
	setup()
	defer teardown()

	cfg := logrusutil.NewConfig()
	cfg.Level = "foobar"
	if err := logrusutil.ConfigureRoot(cfg); err == nil {
		t.Error()
	}
}

func TestConfigureRoot_HookLevel_BadLevel(t *testing.T) {
	setup()
	defer teardown()

	cfg := logrusutil.NewConfig()
	cfg.HookLevel = "foobar"
	if err := logrusutil.ConfigureRoot(cfg); err == nil {
		t.Error()
	}
}

func TestConfigureRoot_SetLevel(t *testing.T) {
	setup()
	defer teardown()

	for _, level := range logrus.AllLevels {
		cfg := logrusutil.NewConfig()
		cfg.Level = level.String()
		if err := logrusutil.ConfigureRoot(cfg); err != nil {
			t.Error(err)
		}
		if logrusutil.Logger.Level != level {
			t.Errorf("%s != %s", logrusutil.Logger.Level, level)
		}
	}
}

func TestConfigureRoot(t *testing.T) {
	setup()
	defer teardown()
	cfg := logrusutil.NewConfig()
	cfg.Level = "warning"
	cfg.HookLevel = "warning"

	for _, hooks := range logrusutil.Logger.Hooks {
		if len(hooks) != 0 {
			t.Error()
		}
	}

	if err := logrusutil.ConfigureRoot(cfg); err != nil {
		t.Error(err)
	}

	for _, hooks := range logrusutil.Logger.Hooks {
		if len(hooks) != 1 {
			t.Error()
		}
	}
}
