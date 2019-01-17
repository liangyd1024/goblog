package model

type Paging struct {
	PageSize int `json:"pageSize" orm:"-"` //每页大小
	PageNum  int `json:"pageNum" orm:"-"`  //页码，从1开始
	Total    int `json:"total" orm:"-"`    //总数
	Pages    int `json:"pages" orm:"-"`    //总页数
}

func (paging *Paging) StartPage() (pageSize, offset int) {
	if paging.PageNum < 1 {
		paging.PageNum = 1
	}
	if paging.PageSize == 0 {
		paging.PageSize = 10
	}
	if paging.PageNum > 1 {
		offset = (paging.PageNum - 1) * paging.PageSize
	}
	pageSize = paging.PageSize
	return
}

func (paging *Paging) CalPages(total int64) {
	if paging.PageSize == 0 {
		paging.PageSize = 10
	}
	paging.Total = int(total)
	pages := paging.Total / paging.PageSize
	if paging.Total%paging.PageSize > 0 {
		pages++
	}
	paging.Pages = pages
}
