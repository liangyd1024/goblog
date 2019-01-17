//AppName: goblog
//Version: V1.0.0
//User: marco
//Date: 2018/11/20

package service

import (
	"goblog/src/dal"
	"goblog/src/utils/bizerror"
	"sync"
)

type baseService struct {
	Mutex *sync.Mutex
}

//获取基础mapper
func (base baseService) GetMapper() *dal.BaseMapper {
	return new(dal.BaseMapper)
}

func (base baseService) Lock(){
	if base.Mutex == nil{
		bizerror.BizError404900.PanicError()
	}
	base.Mutex.Lock()
}

func (base baseService) Unlock(){
	if base.Mutex == nil{
		bizerror.BizError404900.PanicError()
	}
	base.Mutex.Unlock()
}