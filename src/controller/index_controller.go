package controller

import "fmt"

type IndexController struct {
	BaseController
}

func (index *IndexController) Test() {
	fmt.Println("index.Ctx.Input.Protocol:", index.Ctx.Input.Protocol())
	fmt.Println("index.Ctx.Input.URI:", index.Ctx.Input.URI())
	fmt.Println("index.Ctx.Input.URL:", index.Ctx.Input.URL())
	fmt.Println("index.Ctx.Input.Site:", index.Ctx.Input.Site())
	fmt.Println("index.Ctx.Input.Scheme:", index.Ctx.Input.Scheme())
	fmt.Println("index.Ctx.Input.Domain:", index.Ctx.Input.Domain())
	fmt.Println("index.Ctx.Input.SubDomains:", index.Ctx.Input.SubDomains())
	fmt.Println("index.Ctx.Input.Method:", index.Ctx.Input.Method())
	fmt.Println("index.Ctx.Input.IP:", index.Ctx.Input.IP())
	fmt.Println("index.Ctx.Input.Proxy:", index.Ctx.Input.Proxy())
	fmt.Println("index.Ctx.Input.Refer:", index.Ctx.Input.Refer())
	fmt.Println("index.Ctx.Input.Referer:", index.Ctx.Input.Referer())
	fmt.Println("index.Ctx.Input.Port:", index.Ctx.Input.Port())
	fmt.Println("index.Ctx.Input.UserAgent:", index.Ctx.Input.UserAgent())

	test := index.Ctx.Input.Cookie("test")
	fmt.Println("index.Ctx.Input.Cookie:", test)
	if test == "" {
		fmt.Println("index.Ctx.Input.Cookie set")
		index.Ctx.Output.Cookie("test", "213123213213")
	}

	v := index.GetSession("asta")
	if v == nil {
		index.SetSession("asta", int(1))
		index.Data["num"] = 0
	} else {
		index.SetSession("asta", v.(int)+1)
		index.Data["num"] = v.(int)
	}

	index.Data["ctxPath"] = index.Ctx.Input.Host()
	index.Data["msg"] = "hello world"
	index.TplName = "test.html"
}

func (index *IndexController) Index() {
	index.TplName = "index.html"
}

func (index *IndexController) Details() {
	index.TplName = "details.html"
}

func (index *IndexController) About() {
	index.TplName = "about.html"
}

func (index *IndexController) Contact() {
	index.TplName = "contact.html"
}
