package randx

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

// NewN return vcode with a length of n
func NewN(n int) string {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	return fmt.Sprintf("%06v", rnd.Int31n(int32(math.Pow10(n))))
}

// New equal NewN(6)
func New() string {
	return NewN(6)
}
