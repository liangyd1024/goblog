//AppName: goblog
//Version: V1.0.0
//User: marco
//Date: 2018/12/15

package controller

type ErrorController struct{
	BaseController
}

func (c *ErrorController) Error404() {
	c.TplName = "404.html"
}

func (c *ErrorController) Error500() {
	c.TplName = "500.html"
}