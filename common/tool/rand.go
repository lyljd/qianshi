package tool

import (
	"math"
	"math/rand"
	"strconv"
	"strings"
)

// RandNumStr 返回给定位数的的随机数（返回值已自动转换为string类型）
func RandNumStr(length int) string {
	low := int(math.Pow10(length - 1))
	return strconv.Itoa(rand.Intn(int(math.Pow10(length))-low) + low)
}

// RandStr 返回给定位数的随机字符串（只包含大小写字母和数字）
func RandStr(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	var builder strings.Builder
	for i := 0; i < length; i++ {
		idx := rand.Intn(62) // 修改了charset后需要同步更新这里的长度
		builder.WriteByte(charset[idx])
	}

	return builder.String()
}
