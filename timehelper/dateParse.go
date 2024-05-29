package timehelper

import (
	"fmt"
	"github.com/cockroachdb/errors"
	"gitlab.weimiaocaishang.com/components/go-utils/constant"
	"gitlab.weimiaocaishang.com/components/go-utils/errutil"
	"strconv"
	"strings"
	"time"
	"unicode"
)

type CompareDateResult struct {
	Status      bool // true 为赋值成功
	Day         int
	Hour        int
	Minute      int
	Second      float64
	Desc        string
	TotalSecond float64
}

/*
 * getDateFormat
 * @Description: 获取实际格式
 * @Author: LiuQHui
 * @Param date
 * @Return string
**/
func GetDateFormat(date string) string {
	if _, err := time.Parse("2006-01-02", date); err == nil {
		return "2006-01-02"
	} else if _, err := time.Parse("2006-01-02 15:04:05", date); err == nil {
		return "2006-01-02 15:04:05"
	} else if _, err := time.Parse("2006-01-02 15:04:05.000", date); err == nil {
		return "2006-01-02 15:04:05.000"
	}
	return ""
}

/*
 * DateSub
 * @Description: 时间相减
 * @Author: LiuQHui
 * @Param maxDate
 * @Param minDate
 * @Param format
 * @Return CompareDateResult
 * @Return error
**/
func DateSub(maxTime, minTime string) (CompareDateResult, error) {
	var result CompareDateResult
	maxDate, minDate, err := parseCompareTime(maxTime, minTime)
	if err != nil {
		return result, err
	}
	// diff = 55h0m2.45s
	result.TotalSecond = maxDate.Sub(minDate).Seconds()
	diff := maxDate.Sub(minDate).String()
	fieldsFunc := strings.FieldsFunc(diff, func(r rune) bool {
		return unicode.IsLetter(r)
	})
	result.Desc = diff
	if len(fieldsFunc) == 3 {
		result.Status = true
		// 总小时
		totalHours, _ := strconv.Atoi(fieldsFunc[0])
		day := totalHours / 24
		// 天
		result.Day = day
		// 小时
		result.Hour = totalHours - (day * 24)
		// 分钟
		result.Minute, _ = strconv.Atoi(fieldsFunc[1])
		// 秒
		second, _ := strconv.ParseFloat(fieldsFunc[2], 64)
		result.Second = second
	}
	return result, nil
}

/*
 * IsDiffDay
 * @Description: 判断时间差是否满足 >= day
 * @Author: LiuQHui
 * @Param max
 * @Param min
 * @Param day
 * @Return bool
 * @Return error
**/
func IsDiffDay(max string, min string, day int) (bool, error) {
	maxDate, minDate, err := parseCompareTime(max, min)
	if err != nil {
		return false, err
	}
	hours := maxDate.Sub(minDate).Hours()
	if hours/float64(24) >= float64(day) {
		return true, nil
	}
	return false, nil
}

/*
 * parseTime
 * @Description: 解析时间
 * @Author: LiuQHui
 * @Param max
 * @Param min
 * @Return time.Time
 * @Return time.Time
 * @Return error
**/
func parseCompareTime(max string, min string) (time.Time, time.Time, error) {
	if len(max) != len(min) {
		return time.Time{}, time.Time{}, errors.New("两个时间格式需要一致")
	}
	format := GetDateFormat(max)
	minDate, err1 := time.ParseInLocation(format, min, time.Local)
	maxDate, err2 := time.ParseInLocation(format, max, time.Local)
	if err1 != nil || err2 != nil {
		return time.Time{}, time.Time{}, errors.New("时间解析失败")
	}
	if maxDate.After(minDate) {
		return maxDate, minDate, nil
	}
	return minDate, maxDate, nil
}

/*
* @Description: 字符串时间格式化
* @Author: LiuQHui
* @Param date
* @Param format
* @Date 2024-05-29 13:59:29
 */
func FormatLocalStrDate(date string, format string) (string, error) {
	dateFormat := GetDateFormat(date)
	location, err := time.ParseInLocation(dateFormat, date, time.Local)
	if err != nil {
		return "", err
	}
	retDate := location.Format(format)
	return retDate, nil
}

/*
* @Description: 获取今天是当月的第几天
* @Author: LiuQHui
* @Return int
* @Date 2023-05-12 17:28:01
 */
func GetDaysOfMonth(date ...string) (int, error) {
	if len(date) == 0 {
		return time.Now().Day(), nil
	}
	// 解析时间
	parseDate, err := ParseDate(date[0])
	if err != nil {
		return 0, err
	}
	return parseDate.Day(), nil
}

/*
* @Description: 获取指定月的第一天和最后一天
* @Author: LiuQHui
* @Param date
* @Return []*string
* @Return error
* @Date 2023-05-12 17:53:47
 */
func GetMonthBeginAndEndDate(date ...string) ([]string, error) {
	var dateTmp string
	if len(date) == 0 {
		dateTmp = time.Now().Format(constant.YYYYMMDD)
	} else {
		dateTmp = date[0]
	}
	// 解析时间
	parseDate, err := ParseDate(dateTmp)
	if err != nil {
		return nil, err
	}
	var result []string
	// 获取第一天
	firstDayTime := time.Date(parseDate.Year(), parseDate.Month(), 1, 0, 0, 0, 0, parseDate.Location())
	result = append(result, firstDayTime.Format(constant.YYYYMMDD))
	// 获取月内总天数
	_, _, daysInMonth := time.Date(parseDate.Year(), parseDate.Month()+1, 1, 0, 0, 0, 0, parseDate.Location()).Add(-24 * time.Hour).Date()
	// 获取最后一天

	lastDayTime := time.Date(parseDate.Year(), parseDate.Month(), daysInMonth, 0, 0, 0, 0, parseDate.Location())
	result = append(result, lastDayTime.Format(constant.YYYYMMDD))
	return result, nil
}

/*
* @Description: 解析时间
* @Author: LiuQHui
* @Param date
* @Return time.Time
* @Return error
* @Date 2023-05-12 17:47:34
 */
func ParseDate(date string) (time.Time, error) {
	format := GetDateFormat(date)
	if format == "" {
		return time.Time{}, errutil.ThrowErrorMsg("日期格式解析错误~")
	}
	return time.ParseInLocation(format, date, time.Local)
}

/*
* @Description: 获取上个月信息，返回格式 2023-05
* @Author: LiuQHui
* @Return string
* @Date 2023-05-19 18:41:17
 */
func GetLastMonth() string {
	return time.Now().AddDate(0, -1, 0).Format(constant.YYYYMM)
}

/*
* @Description: 比较时间返回多久前
* @Author: LiuQHui
* @Param t1
* @Param t2
* @Return string
* @Date 2023-05-23 11:36:44
 */
func CompareTime(t1 time.Time, t2 time.Time) string {
	diff := t1.Sub(t2)
	if diff.Seconds() < 60 {
		return fmt.Sprintf("%d秒前", int(diff.Seconds()))
	} else if diff.Minutes() < 60 {
		return fmt.Sprintf("%d分钟前", int(diff.Minutes()))
	} else if diff.Hours() < 24 {
		return fmt.Sprintf("%d小时前", int(diff.Hours()))
	} else {
		return "超过一天前"
	}
}

/*
* @Description: 获取指定日期所属月的第一天和最后一天
* @Author: LiuQHui
* @Param date
* @Return string
* @Return string
* @Return error
* @Date 2023-05-26 15:58:35
 */
func GetMonthBeginAndEndDateTime(date ...string) (string, string, error) {
	dates, err := GetMonthBeginAndEndDate(date...)
	if err != nil {
		return "", "", err
	}
	return dates[0] + " 00:00:00", dates[1] + " 23:59:59", nil
}
