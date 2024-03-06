package key

import "strconv"

func GetTokenExp(uid uint) string {
	return "token:exp:" + strconv.FormatUint(uint64(uid), 10)
}

func GetBindEmail(bindEmail string) string {
	return "bind:email:" + bindEmail
}
