//AppName: goblog
//Version: V1.0.0
//User: marco
//Date: 2018/11/21

package dal

import "goblog/src/model"

type TagMapper struct {
	BaseMapper
}

func (tagMapper *TagMapper) GetByCondition(tag *model.Tag) []*model.Tag {
	var tags []*model.Tag
	orm := getOrmer()
	querySetter := orm.QueryTable(tag)
	if tag.TagName != "" {
		querySetter = querySetter.Filter("tag_name__contains",tag.TagName)
	}
	querySetter.All(&tags)
	return tags
}
