//AppName: goblog
//Version: V1.0.0
//User: marco
//Date: 2018/12/12

package crypt

import (
	"crypto/md5"
	"encoding/hex"
)

var key = "!@#$%^&*()_GOBLOG"

func GetMd5(data string) string {
	if data == "" {
		return data
	}
	data += "&key=" + key
	return hex.EncodeToString(md5.New().Sum([]byte(data)))
}
