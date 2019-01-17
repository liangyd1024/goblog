package bizerror

var (
	//400类
	BizError400001 = BizError{"400001", "参数错误:"}
	BizError400002 = BizError{"400002", "数据已存在!"}
	BizError400003 = BizError{"400003", "关联数据已存在!"}
	BizError400004 = BizError{"400004", "数据状态错误!"}

	BizError400101 = BizError{"400101", "用户名或密码错误!"}
	BizError400102 = BizError{"400102", "验证码错误!"}

	//404类
	BizError404001 = BizError{"404001", "数据不存在!"}
	BizError404002 = BizError{"404002", "文件不存在!"}

	BizError404101 = BizError{"404101", "用户不存在!"}

	BizError404900 = BizError{"404900", "系统锁不存在!"}

	//500类错误
	BizError500100 = BizError{"500100", "系统繁忙!"}

	BizError500200 = BizError{"500200", "数据库操作异常!"}

)

type BizError struct {
	ErrCode string
	ErrMsg  string
}

func NewError(code, msg string) BizError {
	return BizError{code, msg}
}

func (bizError BizError) PanicError() {
	panic(bizError)
}

func (bizError BizError) PanicErrorMsg(msg string) {
	bizError.ErrMsg = bizError.ErrMsg + msg
	panic(bizError)
}

func (bizError BizError) Error() string {
	return bizError.ErrCode + "-" + bizError.ErrMsg
}

func Check(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func CheckBizError(err error, bizError BizError) {
	if err != nil {
		panic(err.Error() + bizError.Error())
	}
}

func DbCheck(result int64, err error) {
	Check(err)
	if result == 0 {
		panic(BizError500200)
	}
}
