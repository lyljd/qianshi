package key

func GetCaptchaVerify(captchaId string) string {
	return "captcha:verify:" + captchaId
}

func GetVcodeLoginVerify(email string) string {
	return "vcode:verify:login:" + email
}

func GetVcodeChangePasswordVerify(email string) string {
	return "vcode:verify:changePassword:" + email
}

func GetUserChangePasswordVerify(email string) string {
	return "user:verify:changePassword:" + email
}
