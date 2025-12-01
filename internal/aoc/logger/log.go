package logger

import "fmt"

type Logger struct {
	enabled bool
}

func New(enabled bool) Logger {
	return Logger{enabled: enabled}
}

func (l Logger) Log(a ...any) {
	if !l.enabled {
		return
	}

	fmt.Println(a...)
}

func (l Logger) Logf(format string, args ...any) {
	if !l.enabled {
		return
	}

	fmt.Printf(format, args...)
	fmt.Print("\n")
}
