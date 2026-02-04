package hanalog

type LogLevel int

const (
	Trace LogLevel = iota
	Debug
	Info
	Warn
	Error
	Fatal
)

func (l LogLevel) String() string {
	return [...]string{
		"TRACE", "DEBUG", "INFO", "WARN", "ERROR", "FATAL",
	}[l]
}
