package ctx

import (
	"context"
	"encoding/json"
)

func GetUid(ctx context.Context) int64 {
	var uid int64
	if jsonUid, ok := ctx.Value("uid").(json.Number); ok {
		uid, _ = jsonUid.Int64()
	}
	return uid
}
