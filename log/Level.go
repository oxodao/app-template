package log

type Level struct {
	Value int
	Text  string
	Color Color
}

var (
	LevelDebug = Level{ 0, "DEBUG", white}
	LevelError = Level{ 1, "ERROR", red}
	LevelWarn  = Level{ 2, "WARN", yellow}
	LevelInfo  = Level{ 3, "INFO", green}

	LevelFatal = Level{ 999999, "DEBUG", red}
	LevelInfoAlwaysOn = Level{ 999999, "INFO", green}
)

