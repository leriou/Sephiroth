package utils

import "time"

type TimeUtil struct {
}

/**
 * 获取时间
 */
func (t *TimeUtil) GetTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

/**
 *
 * 获取时间戳
 */
func (t *TimeUtil) GetTimeStamp() int64 {
	return time.Now().UnixNano()
}
