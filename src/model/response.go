package model

//响应mode
//true:result必定有值;false:Code&msg必定有值
type Response struct {
	Paging  `json:"paging"`              //分页信息
	Success bool        `json:"success"` //是否成功
	Code    string      `json:"code"`    //响应码
	Msg     string      `json:"msg"`     //响应描述
	Result  interface{} `json:"result"`  //响应结果
}

func (resp *Response) Suc(result interface{}) *Response {
	resp.Success = true
	resp.Result = result
	return resp
}

func (resp *Response) SucPage(result interface{}, paging Paging) *Response {
	resp.Success = true
	resp.Result = result
	resp.Paging = paging
	return resp
}

func (resp *Response) Fail(msg string) *Response {
	resp.Success = false
	resp.Msg = msg
	return resp
}

func (resp *Response) FailAll(code, msg string) *Response {
	resp.Success = false
	resp.Code = code
	resp.Msg = msg
	return resp
}
