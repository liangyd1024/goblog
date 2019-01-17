//AppName: goblog
//Version: V1.0.0
//User: marco
//Date: 2018/11/16

package check

import (
	"github.com/astaxie/beego/validation"
	"goblog/src/utils/bizerror"
)

func CheckParams(params interface{}) {
	if params != nil {
		valid := validation.Validation{}
		result, err := valid.Valid(params)
		bizerror.Check(err)
		if !result {
			for _, err := range valid.Errors {
				bizerror.BizError400001.PanicErrorMsg(err.Field+" "+err.Message)
			}
		}
	}
}
