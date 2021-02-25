package log

type Level struct {
	Value int
	Text  string
	Color Color
}

var (
	LevelDebug = Level{ 0, "DEBUG", white}
	LevelInfo  = Level{ 1, "INFO", green}
	LevelWarn  = Level{ 2, "WARN", yellow}
	LevelError = Level{ 3, "ERROR", red}

	LevelInfoAlwaysOn = Level{ 999999, "INFO", green}
)

