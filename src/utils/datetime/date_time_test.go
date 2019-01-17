//AppName: goblog
//Version: V1.0.0
//User: marco
//Date: 2018/12/10

package datetime

import (
	"fmt"
	"testing"
	"time"
)

func TestGetNowFm(t *testing.T) {
	fmt.Println(ParseNowTime(FM_FULL_DATE_TIME))
	fmt.Println(ParseNowTime(FM_DATE))
	fmt.Println(ParseNowTime(FM_TIME))
	fmt.Println(ParseNowTime(FM_DATE_TIME))
}

func TestFormatTime(t *testing.T) {
	fmt.Println(FormatTime(time.Now(),FM_DATE_MOUNTH))
}

func TestParseTime(t *testing.T) {
	fmt.Println(ParseTime(FM_DATE_MOUNTH,"2018-12"))
}