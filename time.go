package main

import "time"

const DATETIMEFORMAT = "2006-01-02 15:04:05" //必须是2006-01-02 15:04:05这个时间
const DATEFORMAT = "2006-01-02"
const DURATIONEND = "23h59m59s"

func main() {

}

//获取当前时间的时间戳
func GetTimeUnix() int64 {
	return time.Now().Unix()
}

func GetTimeDate() string {
	return time.Now().Format(DATETIMEFORMAT)
}

//将unix转换为DATETIME格式的时间
func UnixToDateTime(u int64) string {
	return time.Unix(u, 0).Format(DATETIMEFORMAT)
}

//将DATETIME格式的时间转换为时间戳
func DateTimeToUnix(datetime string) int64 {
	parse, e := time.ParseInLocation(DATETIMEFORMAT, datetime, time.Local)
	if e != nil {
		return 0
	}
	return parse.Unix()
}

// 获取当天开始时间戳 通过datetime字符串
func GetDayStartUnix(datetime string) int64 {
	return time.Unix(DateTimeToUnix(datetime+" 00:00:00"), 0).Unix()
}

// 获取当天结束时间戳 通过datetime字符串
func GetDayEndUnix(datetime string) int64 {
	return time.Unix(DateTimeToUnix(datetime+" 23:59:59"), 0).Unix()
}

// 获取当天开始时间戳 通过unix时间戳
func GetDayStartByUnix(u int64) int64 {
	t := time.Unix(u, 0).Format(DATEFORMAT)
	ts, e := time.ParseInLocation(DATEFORMAT, t, time.Local)
	if e != nil {
		panic(e)
		return 0
	}
	return ts.Unix()
}

//获取当天结束时间戳 通过unix时间戳
func GetDayEndByUnix(u int64) int64 {
	duration, e := time.ParseDuration(DURATIONEND)
	if e != nil {
		panic(e)
		return 0
	}
	return GetDayStartByUnix(u) + int64(duration.Seconds())
}

//获取两个时间点相差的秒数
func GetTimeSub(t1 time.Time, t2 time.Time) float64 {
	var d time.Duration
	var sub float64
	switch {
	case t1.After(t2):
		d = t1.Sub(t2)
	case t1.Before(t2):
		d = t2.Sub(t1)
	case t1.Equal(t2):
		return 0
	}
	sub = d.Seconds()
	return sub
}

//获取t时间点的相差years年months月days天的日期的字符串
func GetDateTimeByN(t time.Time, years int, months int, days int) string {
	return t.AddDate(years, months, days).Format(DATEFORMAT)
}
