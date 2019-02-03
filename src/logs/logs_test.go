//AppName: goblog
//Version: V1.0.0
//User: marco
//Date: 2019-02-02

package logs

import (
	"goblog/src/utils/datetime"
	"testing"
	"time"
)

func TestInitLogs(t *testing.T) {
	InitLogs(true)
	for {
		time.Sleep(time.Millisecond * 200)
		Log.Info(datetime.ParseNowTime(datetime.FM_FULL_DATE_TIME_S))
	}

}
