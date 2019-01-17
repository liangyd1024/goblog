package dal

import (
	"github.com/astaxie/beego/orm"
	. "goblog/src/logs"
	"goblog/src/model"
	"goblog/src/utils/bizerror"
	"goblog/src/utils/datetime"
	"strconv"
	"strings"
)

type BowenMapper struct {
	BaseMapper
}

func (bowenMapper *BowenMapper) updateTagCategory(articles *model.Articles, ormer orm.Ormer) {
	//标签信息填充
	tagLength := len(articles.ArticlesTags)
	if tagLength > 0 {
		if articles.ArticlesTags != nil {
			for _, articlesTag := range articles.ArticlesTags {
				articlesTag.ArticlesId = articles.Id
				articlesTag.CreateBy = articles.CreateBy
				//标签文章数更新
				tag := &model.Tag{Id: articlesTag.TagId}
				bizerror.Check(ormer.Read(tag))
				tag.ArticlesNum = tag.ArticlesNum + 1
				bizerror.DbCheck(ormer.Update(tag))
			}
			bizerror.DbCheck(ormer.InsertMulti(tagLength, articles.ArticlesTags))
		}
	}

	//栏目信息填充
	categoryLength := len(articles.ArticlesCategorys)
	if categoryLength > 0 {
		if articles.ArticlesCategorys != nil {
			for _, articlesCategory := range articles.ArticlesCategorys {
				articlesCategory.ArticlesId = articles.Id
				articlesCategory.CreateBy = articles.CreateBy
				//栏目文章数更新
				category := &model.Category{Id: articlesCategory.CategoryId}
				bizerror.Check(ormer.Read(category))
				category.ArticlesNum = category.ArticlesNum + 1
				bizerror.DbCheck(ormer.Update(category))
			}
			bizerror.DbCheck(ormer.InsertMulti(categoryLength, articles.ArticlesCategorys))
		}
	}
}

func (bowenMapper *BowenMapper) Publish(articles *model.Articles) {
	transaction(func(ormer orm.Ormer) {
		bizerror.DbCheck(ormer.Insert(articles))
		articles.ArticlesDetails.Id = articles.Id
		bizerror.DbCheck(ormer.Insert(articles.ArticlesDetails))

		bowenMapper.updateTagCategory(articles, ormer)
	})
}

func (bowenMapper *BowenMapper) Modify(articles *model.Articles) {
	transaction(func(ormer orm.Ormer) {
		bizerror.DbCheck(ormer.Update(articles))
		bizerror.DbCheck(ormer.Update(articles.ArticlesDetails))

		bowenMapper.updateTagCategory(articles, ormer)
	})
}

func (bowenMapper *BowenMapper) GetByCondition(articles *model.Articles) []*model.Articles {
	var articlesList []*model.Articles
	pageSize, offset := articles.Paging.StartPage()
	ormer := getOrmer()
	cond := strings.Builder{}

	if articles.Id != 0 {
		cond.WriteString(" and `id` = ")
		cond.WriteString(strconv.Itoa(articles.Id))
	}
	if articles.Type != "" {
		cond.WriteString(" and `type` = '")
		cond.WriteString(articles.Type)
		cond.WriteString("'")
	}
	if articles.Status != "" {
		cond.WriteString(" and `status` = '")
		cond.WriteString(articles.Status)
		cond.WriteString("'")
	}
	if articles.Title != "" && articles.Desc != "" {
		cond.WriteString(" and (`title` LIKE BINARY '%")
		cond.WriteString(articles.Title)
		cond.WriteString("%' ")
		cond.WriteString(" or `desc` LIKE BINARY '%")
		cond.WriteString(articles.Desc)
		cond.WriteString("%'")
		cond.WriteString(")")
	}
	if !articles.PublishTime.IsZero() {
		cond.WriteString(" and date_format(`publish_time`,'%Y-%m') = '")
		cond.WriteString(datetime.FormatTime(articles.PublishTime, datetime.FM_DATE_MOUNTH))
		cond.WriteString("'")
	}

	total := 0
	err := ormer.Raw("select count(*) from `t_goblog_articles` where 1=1 " + cond.String()).
		QueryRow(&total)
	bizerror.Check(err)
	if total == 0 {
		return articlesList
	}

	articlesList = make([]*model.Articles, 10)
	rows, err := ormer.Raw("select * from `t_goblog_articles` where 1=1 "+cond.String()+" order by `publish_time` desc limit ?,?", offset, pageSize).
		QueryRows(&articlesList)
	bizerror.Check(err)

	//querySeter := getOrmer().QueryTable(articles)
	//
	//cond := orm.NewCondition()
	//if articles.Id != 0 {
	//	cond = cond.And("id", articles.Id)
	//	querySeter = querySeter.SetCond(cond)
	//}
	//if articles.Type != "" {
	//	cond = cond.And("type", articles.Type)
	//	querySeter = querySeter.SetCond(cond)
	//}
	//if articles.Status != "" {
	//	cond = cond.And("status", articles.Status)
	//	querySeter = querySeter.SetCond(cond)
	//}
	//if articles.Title != "" && articles.Desc != "" {
	//	cond = cond.AndCond(orm.NewCondition().And("title__contains", articles.Title).Or("desc__contains", articles.Desc))
	//	querySeter = querySeter.SetCond(cond)
	//}
	//if !articles.PublishTime.IsZero() {
	//	cond = cond.And("publish_time", time.FormatTime(articles.PublishTime, time.FM_DATE_MOUNTH))
	//	querySeter = querySeter.SetCond(cond)
	//}
	//
	//total, err := querySeter.Count()
	//bizerror.Check(err)
	//
	//rows, err := querySeter.
	//	OrderBy("-publish_time").
	//	Limit(pageSize, offset).
	//	All(&articlesList)
	//bizerror.Check(err)

	articles.Paging.CalPages(int64(total))

	Log.Printf("call GetByCondition rows:%v", rows)
	return articlesList
}

func (bowenMapper *BowenMapper) GetArticlesTags(articlesTag *model.ArticlesTag) []*model.ArticlesTag {
	var articlesTags []*model.ArticlesTag
	pageSize, offset := articlesTag.Paging.StartPage()
	querySeter := getOrmer().QueryTable(new(model.ArticlesTag))

	if articlesTag.ArticlesId != 0 {
		querySeter = querySeter.Filter("articles_id", articlesTag.ArticlesId)
	}
	if articlesTag.TagId != 0 {
		querySeter = querySeter.Filter("tag_id", articlesTag.TagId)
	}

	total, err := querySeter.Count()
	bizerror.Check(err)

	rows, err := querySeter.
		OrderBy("-update_at").
		Limit(pageSize, offset).
		All(&articlesTags)
	bizerror.Check(err)

	articlesTag.Paging.CalPages(total)

	Log.Printf("call GetArticlesTags rows:%v", rows)
	return articlesTags
}

func (bowenMapper *BowenMapper) GetTags(articles *model.Articles) ([]*model.Tag, []*model.ArticlesTag) {
	articlesTag := &model.ArticlesTag{ArticlesId: articles.Id}
	articlesTag.InitPaging()
	articlesTags := bowenMapper.GetArticlesTags(articlesTag)

	var tags = make([]*model.Tag, len(articlesTags))
	for index, articlesTag := range articlesTags {
		tag := &model.Tag{Id: articlesTag.TagId}
		bowenMapper.Get(tag)
		tags[index] = tag
	}
	return tags, articlesTags
}

func (bowenMapper *BowenMapper) GetArticlesCategorys(articlesCategory *model.ArticlesCategory) []*model.ArticlesCategory {
	var articlesCategorys []*model.ArticlesCategory
	querySeter := getOrmer().QueryTable(new(model.ArticlesCategory))
	if articlesCategory.ArticlesId != 0 {
		querySeter = querySeter.Filter("articles_id", articlesCategory.ArticlesId)
	}
	if articlesCategory.CategoryId != 0 {
		querySeter = querySeter.Filter("category_id", articlesCategory.CategoryId)
	}
	rows, err := querySeter.
		All(&articlesCategorys)
	bizerror.Check(err)
	Log.Printf("call GetArticlesCategorys rows:%v", rows)
	return articlesCategorys
}

func (bowenMapper *BowenMapper) GetCategorys(articles *model.Articles) ([]*model.Category, []*model.ArticlesCategory) {
	articlesCategorys := bowenMapper.GetArticlesCategorys(&model.ArticlesCategory{ArticlesId: articles.Id})

	var categorys = make([]*model.Category, len(articlesCategorys))
	for index, articlesTag := range articlesCategorys {
		category := &model.Category{Id: articlesTag.CategoryId}
		bowenMapper.Get(category)
		categorys[index] = category
	}
	return categorys, articlesCategorys
}

func (bowenMapper *BowenMapper) GetComments(comment *model.Comment) []*model.Comment {
	var comments []*model.Comment
	pageSize, offset := comment.Paging.StartPage()

	querySeter := getOrmer().QueryTable(comment)
	if comment.ArticlesId != 0 {
		querySeter = querySeter.Filter("articles_id", comment.ArticlesId)
	}
	querySeter = querySeter.Filter("parent_id", comment.ParentId)

	total, err := querySeter.Count()
	bizerror.Check(err)

	rows, err := querySeter.
		OrderBy("-comment_time").
		Limit(pageSize, offset).
		All(&comments)
	bizerror.Check(err)

	comment.Paging.CalPages(total)

	Log.Printf("call GetComments rows:%v", rows)
	return comments
}

func (bowenMapper *BowenMapper) PubComment(comment *model.Comment, articles *model.Articles) {
	transaction(func(ormer orm.Ormer) {
		bizerror.DbCheck(ormer.Insert(comment))
		articles.CommentNum = articles.CommentNum + 1
		bizerror.DbCheck(ormer.Update(articles, "comment_num", "update_at"))
	})
}

func (bowenMapper *BowenMapper) DeleteArticles(articles *model.Articles) {
	transaction(func(ormer orm.Ormer) {
		id := articles.Id
		bizerror.DbCheck(ormer.Delete(articles))
		bizerror.DbCheck(ormer.Delete(&model.ArticlesDetails{Id: id}))
		for _, articlesTag := range articles.ArticlesTags {
			bizerror.DbCheck(ormer.Delete(&model.ArticlesTag{Id: articlesTag.Id}))
			//标签文章数更新
			tag := &model.Tag{Id: articlesTag.TagId}
			bizerror.Check(ormer.Read(tag))
			if tag.ArticlesNum > 0 {
				tag.ArticlesNum = tag.ArticlesNum - 1
				bizerror.DbCheck(ormer.Update(tag))
			}
		}
		for _, articlesCategory := range articles.ArticlesCategorys {
			bizerror.DbCheck(ormer.Delete(&model.ArticlesCategory{Id: articlesCategory.Id}))
			//栏目文章数更新
			category := &model.Category{Id: articlesCategory.CategoryId}
			bizerror.Check(ormer.Read(category))
			if category.ArticlesNum > 0 {
				category.ArticlesNum = category.ArticlesNum - 1
				bizerror.DbCheck(ormer.Update(category))
			}
		}
	})
}

func (bowenMapper *BowenMapper) DeleteArticlesTag(articlesTag *model.ArticlesTag) {
	transaction(func(ormer orm.Ormer) {
		bizerror.DbCheck(ormer.Delete(articlesTag))
		//标签文章数更新
		tag := &model.Tag{Id: articlesTag.TagId}
		bizerror.Check(ormer.Read(tag))
		if tag.ArticlesNum > 0 {
			tag.ArticlesNum = tag.ArticlesNum - 1
			bizerror.DbCheck(ormer.Update(tag))
		}
	})
}

func (bowenMapper *BowenMapper) DeleteArticlesCategory(articlesCategory *model.ArticlesCategory) {
	transaction(func(ormer orm.Ormer) {
		bizerror.DbCheck(ormer.Delete(articlesCategory))
		//栏目文章数更新
		category := &model.Category{Id: articlesCategory.CategoryId}
		bizerror.Check(ormer.Read(category))
		if category.ArticlesNum > 0 {
			category.ArticlesNum = category.ArticlesNum - 1
			bizerror.DbCheck(ormer.Update(category))
		}
	})
}

func (bowenMapper *BowenMapper) CollectGroup(field string) []*model.ArticlesCollect {
	articlesCollects := make([]*model.ArticlesCollect, 1)
	ormer := getOrmer()
	raw := ormer.Raw("select t." + field + ",count(1) as `total` from `t_goblog_articles` t group by t." + field)
	rows, err := raw.QueryRows(&articlesCollects)
	bizerror.Check(err)
	Log.Printf("call CollectGroup rows:%v", rows)
	return articlesCollects
}

func (bowenMapper *BowenMapper) CollectPlaceOfFile() []*model.ArticlesCollect {
	articlesCollects := make([]*model.ArticlesCollect, 1)
	ormer := getOrmer()
	raw := ormer.Raw("select date_format(`publish_time`,'%Y-%m') type,count(1) `total` from `t_goblog_articles` group by date_format(`publish_time`,'%Y-%m')")
	rows, err := raw.QueryRows(&articlesCollects)
	bizerror.Check(err)
	Log.Printf("call CollectPlaceOfFile rows:%v", rows)
	return articlesCollects
}

func (bowenMapper *BowenMapper) ListRecommendArticles(size int) []*model.Articles {
	articlesList := make([]*model.Articles, size)
	ormer := getOrmer()
	raw := ormer.Raw("select * from (select * from `t_goblog_articles` GROUP BY id ORDER BY max(`browse_num`) desc LIMIT 0,?) as t", strconv.Itoa(size))
	rows, err := raw.QueryRows(&articlesList)
	bizerror.Check(err)
	Log.Printf("call ListRecommendArticles rows:%v", rows)
	return articlesList
}
