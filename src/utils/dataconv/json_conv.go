//AppName: goblog
//Version: V1.0.0
//User: marco
//Date: 2019/1/2

package dataconv

import (
	"encoding/json"
	"goblog/src/utils/bizerror"
)

func JsonByte2M(data []byte, model interface{}) interface{} {
	err := json.Unmarshal(data, model)
	bizerror.Check(err)
	return model
}

func JsonStr2M(data string, model interface{}) interface{} {
	return JsonByte2M([]byte(data), model)
}

func JsonM2Str(model interface{}) string {
	return string(JsonM2Byte(model))
}

func JsonM2Byte(model interface{}) []byte {
	data, err := json.Marshal(model)
	bizerror.Check(err)
	return data
}
