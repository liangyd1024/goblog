package model

import "time"

type Base struct {
	CreateBy string    `json:"createBy" orm:"size(32)"`                     //创建人
	CreateAt time.Time `json:"createAt" orm:"auto_now_add;type(datetime)"`  //创建时间
	UpdateBy string    `json:"updateBy" orm:"size(32);null"`                //更新人
	UpdateAt time.Time `json:"updateAt" orm:"auto_now;type(datetime);null"` //更新时间
}
