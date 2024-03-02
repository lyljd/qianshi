package key

import "strconv"

func GetTokenExp(uid uint) string {
	return "token:exp:" + strconv.FormatUint(uint64(uid), 10)
}
