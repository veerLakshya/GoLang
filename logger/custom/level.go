package mylogger

type Level int8

const (
	DebugLevel Level = iota - 1 //-1
	InfoLevel                   //0
	WarnLevel
	ErrorLevel
	DPanicLevel
	PanicLevel
	FatalLevel

	_minLevel = DebugLevel
	_maxLevel = FatalLevel

	InvalidLevel = _maxLevel + 1
)

func ParseLevel(text string) (Level, error) {
	var level Level
	err := level.UnmarshalText
}

func (l *Level) UnmarshalText(text []byte) bool {
	switch string(text) {
	case "debug":
		*l = DebugLevel
	case "info":
		*l = InfoLevel
	case "warn", "warning":
		*l = WarnLevel
	case "dpanic":

	}
}
