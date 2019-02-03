//AppName: goblog
//Version: V1.0.0
//User: marco
//Date: 2018/12/10

package datetime

import (
	"goblog/src/utils/bizerror"
	"time"
)

const (
	FM_DATE_TIME_S = "2006-01-02 15:04:05:000"
	FM_DATE_TIME   = "2006-01-02 15:04:05"
	FM_DATE_MOUNTH = "2006-01"
	FM_DATE        = "2006-01-02"
	FM_TIME        = "15:04:05"

	FM_SPRIT_DATE_TIME = "2006/01/02 15:04:05"
	FM_SPRIT_DATE      = "2006/01/02"
	FM_SPRIT_TIME      = "15:04:05"

	FM_FULL_DATE_TIME_S = "20060102150405.000"
	FM_FULL_DATE_TIME   = "20060102150405"
	FM_FULL_DATE        = "20060102"
	FM_FULL_TIME        = "150405"
)

//根据格式获取当前时间
func ParseNowTime(fm string) string {
	return time.Now().Format(fm)
}

func ParseTime(fm, dt string) time.Time {
	fTime, err := time.Parse(fm, dt)
	bizerror.Check(err)
	return fTime
}

func FormatTime(dt time.Time, fm string) string {
	return dt.Format(fm)
}
