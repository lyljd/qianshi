package email

import "github.com/go-gomail/gomail"

const (
	host163 = "smtp.163.com"
	port163 = 465
)

type Dialer struct {
	dialer *gomail.Dialer
	sender string
	from   string
}

func NewDialer(host string, port int, sender, username, password string) *Dialer {
	d := gomail.NewDialer(host, port, username, password)

	if dial, err := d.Dial(); err != nil {
		panic("dial email server error! " + err.Error())
	} else {
		_ = dial.Close()
	}

	return &Dialer{d, sender, username}
}

func New163Dialer(sender, username, password string) *Dialer {
	return NewDialer(host163, port163, sender, username, password)
}
