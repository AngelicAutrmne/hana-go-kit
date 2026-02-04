package log

import "fmt"

func format(level, msg string, fields ...Field) string {
	s := fmt.Sprintf("[%s] %s", level, msg)
	for _, f := range fields {
		s += fmt.Sprintf(" %s=%v", f.Key, f.Value)
	}
	return s
}
