package jsonhelper

import (
	"52lu/go-helpers/timehelper"
	"database/sql/driver"
	"fmt"
	"github.com/thoas/go-funk"
	"strings"
	"time"
)

type DateTime time.Time
type Date time.Time

/*
 * @Description: 处理时间类型
 * @Author: LiuQHui
 * @Receiver d
 * @Return []byte
 * @Return error
**/
func (d DateTime) MarshalJSON() ([]byte, error) {
	var stamp = fmt.Sprintf(`"%s"`, time.Time(d).Format("2006-01-02 15:04:05"))
	if strings.Contains(stamp, "000") {
		stamp = `""`
	}
	return []byte(stamp), nil
}

/*
 * @Description: 反序列化
 * @Author: LiuQHui
 * @Receiver d
 * @Param b
 * @Return error
**/
func (d *DateTime) UnmarshalJSON(b []byte) error {
	val := string(b)
	var t time.Time
	var err error
	if funk.ContainsString([]string{`""`, `null`, `"null"`}, val) || strings.Contains(val, "000") {
		t, _ = time.Parse(`2006-01-02 15:04:05`, "0000-00-00 00:00:00")
	} else {
		t, err = time.ParseInLocation(`"2006-01-02 15:04:05"`, string(b), time.Local)
	}
	if err != nil {
		return err
	}
	*d = DateTime(t)
	return nil
}

func (d DateTime) String() string {
	dt := time.Time(d)
	if dt.IsZero() {
		return ""
	}
	return dt.Format("2006-01-02 15:04:05")
}

/*
 * @Description:
 * @Author: LiuQHui
 * @Receiver d
 * @Return []byte
 * @Return error
**/
func (d Date) MarshalJSON() ([]byte, error) {
	var stamp = fmt.Sprintf(`"%s"`, time.Time(d).Format("2006-01-02"))
	if strings.Contains(stamp, "000") {
		stamp = `""`
	}
	return []byte(stamp), nil
}

/*
 * @Description: 反序列化
 * @Author: LiuQHui
 * @Receiver d
 * @Param b
 * @Return error
**/
func (d *Date) UnmarshalJSON(b []byte) error {
	val := string(b)
	var t time.Time
	var err error
	if funk.ContainsString([]string{`""`, `null`, `"null"`}, val) || strings.Contains(val, "000") {
		t, _ = time.Parse(`2006-01-02`, "0000-00-00")
	} else {
		t, err = time.ParseInLocation(`"2006-01-02"`, string(b), time.Local)
	}
	if err != nil {
		return err
	}
	*d = Date(t)
	return nil
}

func (d Date) String() string {
	dt := time.Time(d)
	if dt.IsZero() {
		return ""
	}
	return dt.Format("2006-01-02")
}

func (d *Date) Scan(src interface{}) error {
	if t, ok := src.(time.Time); ok {
		*d = Date(t)
		return nil
	}
	// 转成字符串
	date2 := fmt.Sprintf("%s", src)
	if strings.Contains(date2, "000") {
		*d = Date(time.Time{})
		return nil
	}
	location, err := timehelper.ParseDate(date2)
	if err == nil {
		*d = Date(location)
		return nil
	}
	return fmt.Errorf("can not convert %v to wmutil.Date", date2)
}

// 读取数据库时会调用该方法将时间数据转换成自定义时间类型
func (d *DateTime) Scan(src interface{}) error {
	if t, ok := src.(time.Time); ok {
		*d = DateTime(t)
		return nil
	}
	// 转成字符串
	// 转成字符串
	date2 := fmt.Sprintf("%s", src)
	if strings.Contains(date2, "000") {
		*d = DateTime(time.Time{})
		return nil
	}
	location, err := timehelper.ParseDate(date2)
	if err == nil {
		*d = DateTime(location)
		return nil
	}
	return fmt.Errorf("can not convert %v to wmutil.DateTime", date2)
}

func (d Date) Value() (driver.Value, error) {
	timeString := d.String()
	if strings.Contains(timeString, "000") {
		return "", nil
	}
	return timeString, nil
}

func (d DateTime) Value() (driver.Value, error) {
	timeString := d.String()
	if strings.Contains(timeString, "000") {
		return "", nil
	}
	return timeString, nil
}

func (d DateTime) IsZero() bool {
	return time.Time(d).IsZero()
}

func (d Date) IsZero() bool {
	return time.Time(d).IsZero()
}
