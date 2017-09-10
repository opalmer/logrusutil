package logrusutil

const (
	// DisabledLevel may be passed into a *Config struct
	// to disable logging and discard output.
	DisabledLevel = "disabled"
)

var (
	// DefaultLevel is the default log level used in NewConfig()
	DefaultLevel = "warning"
)

// Config is used to provided configuration information
type Config struct {
	Level string
}

// NewConfig produces a *Config struct with reasonable defaults.
func NewConfig() *Config {
	return &Config{Level: DefaultLevel}
}
