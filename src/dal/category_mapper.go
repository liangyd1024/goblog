//AppName: goblog
//Version: V1.0.0
//User: marco
//Date: 2018/11/21

package dal

import (
	"goblog/src/model"
)

type CategoryMapper struct {
	BaseMapper
}

func (categoryMapper *CategoryMapper) GetByCondition(tag *model.Category) []*model.Category {
	var categories []*model.Category
	orm := getOrmer()
	querySetter := orm.QueryTable(tag)
	if tag.CategoryName != "" {
		querySetter = querySetter.Filter("category_name__contains", tag.CategoryName)
	}
	querySetter.All(&categories)
	return categories
}
