//AppName: goblog
//Version: V1.0.0
//User: marco
//Date: 2018/12/18

package crypt

import "github.com/mojocn/base64Captcha"

//数字图形码
var configDigit = base64Captcha.ConfigDigit{
	Height:     45,
	Width:      230,
	MaxSkew:    4,
	DotCount:   100,
	CaptchaLen: 4,
}

func GenerateCaptcha() (uid, value string) {
	id, captchaInstance := base64Captcha.GenerateCaptcha("", configDigit)
	base64stringC := base64Captcha.CaptchaWriteToBase64Encoding(captchaInstance)
	uid = id
	value = base64stringC
	return
}

func VerifyCaptcha(uid, value string) bool {
	return base64Captcha.VerifyCaptcha(uid, value)
}
