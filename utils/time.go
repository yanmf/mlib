package utils

import "time"

const (
	HourSeconds = int32(60 * 60)
	DaySeconds  = int64(24 * HourSeconds)
	WeekSeconds = DaySeconds * 7

	TimeLayout = time.DateTime + ".000"
)

func Second() int64 {
	return time.Now().Unix()
}

func Millisecond() int64 {
	return time.Now().UnixMilli()
}

func TimeMilliLayoutFormat(t int64, layout string) string {
	return time.UnixMilli(t).Format(layout)
}

func TimeSecondLayoutFormat(t int64, layout string) string {
	return time.Unix(t, 0).Format(layout)
}
