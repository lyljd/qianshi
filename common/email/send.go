package email

import (
	"errors"
	"fmt"
	"github.com/go-gomail/gomail"
	"mime"
	"strings"
)

type Content struct {
	typ   string
	value string
}

func newMessage(sender, from, to, subject string, content *Content) *gomail.Message {
	m := gomail.NewMessage()

	m.SetHeader("From", fmt.Sprintf("%s <%s>", mime.QEncoding.Encode("UTF-8", sender), from))
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody(content.typ, content.value)

	return m
}

func (d *Dialer) SendToOne(to, subject string, content *Content) error {
	return d.dialer.DialAndSend(newMessage(d.sender, d.from, to, subject, content))
}

func (d *Dialer) SendToMany(to []string, subject string, content *Content) error {
	dial, err := d.dialer.Dial()

	if err != nil {
		return err
	}

	defer func(dial gomail.SendCloser) {
		_ = dial.Close()
	}(dial)

	total := len(to)

	ms := make([]*gomail.Message, total)
	for k, v := range to {
		ms[k] = newMessage(d.sender, d.from, v, subject, content)
	}

	var sendErr []string
	var errNum int

	for _, m := range ms {
		if err := gomail.Send(dial, m); err != nil {
			sendErr = append(sendErr, fmt.Sprintf("send to %s fail! err: %s", m.GetHeader("To")[0], err))
			errNum++
		}
	}

	if errNum == 0 {
		return nil
	}

	return errors.New(fmt.Sprintf("total: %d, fail: %d\n%s", total, errNum, strings.Join(sendErr, "\n")))
}

func Text(value string) *Content {
	return &Content{"text/plain", value}
}

func Html(value string) *Content {
	return &Content{"text/html", value}
}
