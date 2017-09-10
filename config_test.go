package logrusutil_test

import (
	"testing"

	"github.com/opalmer/logrusutil"
)

func TestNewConfig(t *testing.T) {
	cfg := logrusutil.NewConfig()
	if cfg.Level != logrusutil.DefaultLevel {
		t.Error()
	}
}
