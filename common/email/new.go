package email

import (
	"context"
	"github.com/go-gomail/gomail"
	"time"
)

const (
	host163    = "smtp.163.com"
	port163    = 465
	hostExmail = "smtp.exmail.qq.com"
	portExmail = 465
)

type Dialer struct {
	dialer *gomail.Dialer
	sender string
	from   string
}

func NewDialer(host string, port int, sender, username, password string) *Dialer {
	ctx, cancelCtx := context.WithTimeout(context.Background(), time.Second*10)

	go func() {
		<-ctx.Done()
		if ctx.Err() == context.DeadlineExceeded {
			panic("dial email server timeout! Please ensure network connectivity and correct parameter information.")
		}
	}()

	d := gomail.NewDialer(host, port, username, password)

	if dial, err := d.Dial(); err != nil {
		panic("dial email server error! " + err.Error())
	} else {
		_ = dial.Close()
		cancelCtx()
	}

	return &Dialer{d, sender, username}
}

func New163Dialer(sender, username, password string) *Dialer {
	return NewDialer(host163, port163, sender, username, password)
}

func NewExmailDialer(sender, username, password string) *Dialer {
	return NewDialer(hostExmail, portExmail, sender, username, password)
}
