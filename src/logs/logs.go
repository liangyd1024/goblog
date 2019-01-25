//AppName: goblog
//Version: V1.0.0
//User: marco
//Date: 2018/11/17

package logs

import (
	"fmt"
	"github.com/astaxie/beego"
	"goblog/src/utils/bizerror"
	"goblog/src/utils/datetime"
	"io"
	"log"
	"os"
	"sync"
)

var Log *log.Logger

var lock = new(sync.Mutex)
var me *logs

type logs struct {
	console  bool
	file     bool
	filePath string
	log      *log.Logger
	writer   io.Writer
}

func InitLogs() *logs {
	if me == nil {
		lock.Lock()
		defer lock.Unlock()
		if me == nil {
			cfg := beego.AppConfig
			currentPath, err := os.Getwd()
			bizerror.Check(err)
			fmt.Println("call InitLogs currentPath:" + currentPath)

			me = new(logs)
			me.console = cfg.DefaultBool("consoleMode", true)
			me.file = cfg.DefaultBool("fileMode", true)
			me.filePath = cfg.DefaultString("filePath", currentPath+"/log/")

			_, err = os.Stat(me.filePath)
			if os.IsNotExist(err) {
				err := os.MkdirAll(me.filePath, 0766)
				bizerror.CheckBizError(err, bizerror.BizError404002)
			}
			me.filePath += datetime.ParseNowTime(datetime.FM_DATE) + ".log"

			//TODO
			logFile, err := os.OpenFile(me.filePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
			multiWriter := io.MultiWriter(logFile, os.Stdout)
			me.writer = multiWriter
			Log = log.New(multiWriter, "["+cfg.String("appname")+"]", log.Ldate|log.Ltime|log.Lshortfile)
			Log.SetFlags(log.Ldate | log.Lmicroseconds | log.Lshortfile)
			me.log = Log
			Log.Printf("call logs init me:%+v", me)
		}
	}
	return me
}

func GetLogsWriter() io.Writer {
	return me.writer
}
