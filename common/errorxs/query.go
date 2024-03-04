package errorxs

import "errors"

var (
	ErrRecordNotFound = errors.New("记录不存在")
	ErrKeyNotFound    = errors.New("key不存在")
)
