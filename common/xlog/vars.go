package xlog

import (
	"os"
	"qianshi/common/email"
)

var (
	sn          string
	f           *os.File
	ch          chan map[string]any
	done        chan struct{}
	emailDialer *email.Dialer
)
