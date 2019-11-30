package utils

import (
	"fmt"
)

type Logging struct {
	t *TimeUtil
}

func NewLogging() *Logging {
	return new(Logging)
}

func (l *Logging) log(level, info string) {
	fmt.Println("%s [%s] -> %s", l.t.GetTime(), level, info)
}
