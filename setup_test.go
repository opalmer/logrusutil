package logrusutil_test

import (
	"io/ioutil"
	"testing"

	"github.com/Sirupsen/logrus"
	"github.com/opalmer/logrusutil"
)

func setup() {
	logrusutil.StandardLogger = logrus.New()
}

func teardown() {
	logrusutil.StandardLogger = logrus.StandardLogger()
}

func TestConfigureRoot_ErrLevelNotProvided(t *testing.T) {
	setup()
	defer teardown()

	cfg := logrusutil.NewConfig()
	cfg.Level = ""
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

	if logrusutil.StandardLogger.Level != logrus.PanicLevel {
		t.Error()
	}

	if logrusutil.StandardLogger.Out != ioutil.Discard {
		t.Error()
	}
}

func TestConfigureRoot_ParseError(t *testing.T) {
	setup()
	defer teardown()

	cfg := logrusutil.NewConfig()
	cfg.Level = "foobar"
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
		if logrusutil.StandardLogger.Level != level {
			t.Errorf("%s != %s", logrusutil.StandardLogger.Level, level)
		}
	}
}
