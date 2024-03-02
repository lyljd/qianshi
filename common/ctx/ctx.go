package ctx

import (
	"context"
	"strconv"
)

func GetUid(ctx context.Context) uint {
	var uid uint64
	if uidStr, ok := ctx.Value("uid").(string); ok {
		uid, _ = strconv.ParseUint(uidStr, 10, 64)
	}
	return uint(uid)
}
