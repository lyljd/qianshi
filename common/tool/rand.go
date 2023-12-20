package tool

import (
	"math"
	"math/rand"
	"strconv"
)

func RandNumStr(length int) string {
	low := int(math.Pow10(length - 1))
	return strconv.Itoa(rand.Intn(int(math.Pow10(length))-low) + low)
}
