package key

func GetCaptchaVerify(captchaId string) string {
	return "captcha:verify:" + captchaId
}

func GetVcodeLoginVerify(email string) string {
	return "vcode:verify:login" + email
}
