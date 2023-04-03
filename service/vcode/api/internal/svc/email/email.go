package email

import (
	"crypto/tls"
	"fmt"
	"net/smtp"
)

const (
	Hostname163 = "smtp.163.com"
	HostnameQQ  = "smtp.qq.com"
	Hostname126 = "smtp.126.com"
)

type Client struct {
	hostname string
	username string
	auth     smtp.Auth
}

type Options struct {
	Hostname string
	Username string
	Password string
}

func NewClient(o *Options) *Client {
	c, err := smtp.Dial(o.Hostname + ":25")
	if err != nil {
		panic("邮件服务器\"" + o.Hostname + "\"连接失败！" + err.Error())
	}
	err = c.StartTLS(&tls.Config{ServerName: o.Hostname})
	if err != nil {
		panic("邮件服务器开启TLS连接失败！" + err.Error())
	}
	auth := smtp.PlainAuth("", o.Username, o.Password, o.Hostname)
	err = c.Auth(auth)
	if err != nil {
		panic("邮件服务器身份验证失败！" + err.Error())
	}
	return &Client{
		hostname: o.Hostname,
		username: o.Username,
		auth:     auth,
	}
}

func (c *Client) SendMany(to []string, title, content string) error {
	return smtp.SendMail(c.hostname+":25", c.auth, c.username, to, []byte(fmt.Sprintf("Subject: %v\r\n\r\n%v\r\n", title, content)))
}

func (c *Client) SendOne(to, title, content string) error {
	return c.SendMany([]string{to}, title, content)
}
