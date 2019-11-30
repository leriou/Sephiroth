package utils

import "fmt"

type Logger struct {
	t *TimeUtil
}

func NewLogger() *Logger {
	return new(Logger)
}

func (l *Logger) Log(level, info string) {
	fmt.Println(l.t.GetTime() + " [" + level + "] -> " + info)
}

func (l *Logger) Info(msg string) {
	l.Log("info", msg)
}

func (l *Logger) Error(msg string) {
	l.Log("error", msg)
}

func (l *Logger) Debug(msg string) {
	l.Log("debug", msg)
}
