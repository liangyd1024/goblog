package dal

import (
	"github.com/astaxie/beego/orm"
	"goblog/src/utils/bizerror"
)

type Mapper interface {
	Insert(model interface{})

	Update(model interface{})

	Delete(model interface{})

	Get(model interface{}) interface{}

	GetByCondition(model interface{}) []interface{}
}

type BaseMapper struct {
}

func getOrmer() orm.Ormer {
	return orm.NewOrm()
}

//事物处理
func transaction(invoker func(ormer orm.Ormer)) {
	ormer := getOrmer()
	ormer.Begin()

	defer func() {
		if err := recover(); err != nil {
			ormer.Rollback()
			panic(err)
		} else {
			ormer.Commit()
		}
	}()

	invoker(ormer)
}

func (baseMapper *BaseMapper) Insert(base interface{}) {
	bizerror.DbCheck(getOrmer().Insert(base))
}

func (baseMapper *BaseMapper) Update(base interface{}, cols ...string) {
	bizerror.DbCheck(getOrmer().Update(base, cols...))
}

func (baseMapper *BaseMapper) Delete(base interface{}) {
	bizerror.DbCheck(getOrmer().Delete(base))
}

func (baseMapper *BaseMapper) Get(base interface{}, cols ...string) interface{} {
	bizerror.Check(getOrmer().Read(base, cols...))
	return base
}
