package helper

import "time"

const TimeLayout = "2006-01-02 15:04:05"
const DateLayout = "2006-01-02"

// GetThisYearStartEnd 获取当前年的开始结束时间
func GetThisYearStartEnd() (time.Time, time.Time) {
	thisYear := time.Now().Year()
	startTime := time.Date(thisYear, 1, 1, 0, 0, 0, 0, time.Local)
	endTime := time.Date(thisYear, 12, 31, 23, 59, 59, 0, time.Local)
	return startTime, endTime
}

// GetNearlyTwelveMonthTime 获取最近12个月
func GetNearlyTwelveMonthTime() (time.Time, time.Time) {
	thisYear := time.Now().Year()
	thisMonth := time.Now().Month()
	firstOfMonth := time.Date(thisYear, thisMonth, 1, 0, 0, 0, 0, time.Local)
	lastOfMonth := firstOfMonth.AddDate(0, 1, -1)
	// 过去12个月
	endTime := time.Date(thisYear, thisMonth, lastOfMonth.Day(), 23, 59, 59, 0, time.Local)
	startTime := time.Date(thisYear, thisMonth-11, 1, 0, 0, 0, 0, time.Local)
	return startTime, endTime
}

// GetYearMonthToDay 查询指定年份指定月份有多少天
func GetYearMonthToDay(year int, month int) int {
	// 有31天的月份
	day31 := map[int]bool{
		1:  true,
		3:  true,
		5:  true,
		7:  true,
		8:  true,
		10: true,
		12: true,
	}
	if day31[month] == true {
		return 31
	}
	// 有30天的月份
	day30 := map[int]bool{
		4:  true,
		6:  true,
		9:  true,
		11: true,
	}
	if day30[month] == true {
		return 30
	}
	// 计算是平年还是闰年
	if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
		// 得出2月的天数
		return 29
	}
	// 得出2月的天数
	return 28
}

// GetYesterdayTime 获取昨日开始结束时间
func GetYesterdayTime() (time.Time, time.Time) {
	nowTime := time.Now()
	yesterdayTime := nowTime.AddDate(0, 0, -1)
	beginTime, _ := time.ParseInLocation(DateLayout, yesterdayTime.Format(DateLayout), time.Local)
	endTime := beginTime.Add(time.Second * (86400 - 1))

	return beginTime, endTime
}

// GetTodayTime 获取今日开始结束时间
func GetTodayTime() (time.Time, time.Time) {
	nowTime := time.Now()
	startTime := time.Date(nowTime.Year(), nowTime.Month(), nowTime.Day(), 0, 0, 0, 0, nowTime.Location())
	endTime := time.Date(nowTime.Year(), nowTime.Month(), nowTime.Day(), 23, 59, 59, 0, nowTime.Location())
	return startTime, endTime
}

// GetMonthStartEnd 获取指定时间所在月的开始 结束时间
func GetMonthStartEnd(t time.Time) (time.Time, time.Time) {
	monthStartDay := t.AddDate(0, 0, -t.Day()+1)
	monthStartTime := time.Date(monthStartDay.Year(), monthStartDay.Month(), monthStartDay.Day(), 0, 0, 0, 0, t.Location())
	monthEndDay := monthStartTime.AddDate(0, 1, -1)
	monthEndTime := time.Date(monthEndDay.Year(), monthEndDay.Month(), monthEndDay.Day(), 23, 59, 59, 0, t.Location())
	return monthStartTime, monthEndTime
}

// GetLastMonthStartEnd 获取上一个月的开始 结束 时间戳
func GetLastMonthStartEnd() (int64, int64) {
	now := time.Now()
	lastMonthFirstDay := now.AddDate(0, -1, -now.Day()+1)
	lastMonthStart := time.Date(lastMonthFirstDay.Year(), lastMonthFirstDay.Month(), lastMonthFirstDay.Day(), 0, 0, 0, 0, now.Location()).Unix()
	lastMonthEndDay := lastMonthFirstDay.AddDate(0, 1, -1)
	lastMonthEnd := time.Date(lastMonthEndDay.Year(), lastMonthEndDay.Month(), lastMonthEndDay.Day(), 23, 59, 59, 0, now.Location()).Unix()
	return lastMonthStart, lastMonthEnd
}

// GetTimestamp Go对时间字符串中包含T和Z的处理
func GetTimestamp(change string) int64 {
	t, _ := time.Parse(time.RFC3339, change)
	timeUint := t.In(time.Local).Unix()
	return timeUint
}

// Time 获取时间戳
func Time() int64 {
	return time.Now().Unix()
}

// StrToTime 格式化时间 StrToTime("02/01/2006 15:04:05", "02/01/2016 15:04:05") == 1451747045
func StrToTime(format, strTime string) (int64, error) {
	t, err := time.Parse(format, strTime)
	if err != nil {
		return 0, err
	}
	return t.Unix(), nil
}

// CheckDate 检查日期
func CheckDate(month, day, year int) bool {
	if month < 1 || month > 12 || day < 1 || day > 31 || year < 1 || year > 32767 {
		return false
	}
	switch month {
	case 4, 6, 9, 11:
		if day > 30 {
			return false
		}
	case 2:
		// leap year
		if year%4 == 0 && (year%100 != 0 || year%400 == 0) {
			if day > 29 {
				return false
			}
		} else if day > 28 {
			return false
		}
	}

	return true
}

// Sleep 睡眠
func Sleep(t int64) {
	time.Sleep(time.Duration(t) * time.Second)
}
