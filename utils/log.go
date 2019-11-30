package utils

import (
	"fmt"
)

type Logger struct {
	t *TimeUtil
}

func NewLogger() *Logger {
	return new(Logger)
}

func (l *Logger) Log(level, info string) {
	fmt.Println("%s [%s] -> %s", l.t.GetTime(), level, info)
}
