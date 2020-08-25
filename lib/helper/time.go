package helper

import (
    "errors"
    "time"
)

var TimeZone, _ = time.LoadLocation("Asia/Shanghai")

var InitDate string = "2006-01-02"
var InitDateTime string = "2006-01-02 15:04:05"
var InitDateMinute string = "2006-01-02 15:04"

// FormatTime 格式化时间
func FormatTime(time time.Time) string {
    return time.Format("2006-01-02 15:04:05")
}

// UnixTime 将时间转化为毫秒数
func UnixTime(t time.Time) int64 {
    return t.UnixNano() / 1000000
}

func GetDateWithWeekdayBetween(format, from, to string) (map[string]int, error) {
    if from > to {
        return nil, errors.New("from time greater than to time")
    }
    fromTime, err := time.ParseInLocation(format, from, TimeZone)
    if err != nil {
        return nil, err
    }

    toTime, err := time.ParseInLocation(format, to, TimeZone)
    if err != nil {
        return nil, err
    }

    res := make(map[string]int)
    for begin := fromTime; toTime.Sub(begin).Nanoseconds() >= 0; begin = begin.AddDate(0, 0, 1) {
        res[begin.Format(format)] = int(begin.Weekday())
    }
    return res, nil
}

// 验证时间格式是否符合规则
func CheckTime(inputTime string, format string) bool {
    unixtime, err := time.Parse(format, inputTime)
    if err != nil {
        return false
    }
    formatTime := unixtime.Format(format)
    if formatTime != inputTime {
        return false
    }
    return true
}
