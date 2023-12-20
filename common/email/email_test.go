package email

import (
	"fmt"
	"strings"
	"testing"
)

const (
	sender   = ""
	username = "@163.com"
	password = ""
	toOne    = ""
	toMany   = ";"
	subject  = ""
	content  = ""
)

func TestOne(t *testing.T) {
	dialer := New163Dialer(sender, username, password)
	if err := dialer.SendToOne(toOne, subject, Text(content)); err != nil {
		fmt.Println(err)
		t.Fail()
	}
}

func TestMany(t *testing.T) {
	dialer := New163Dialer(sender, username, password)
	if err := dialer.SendToMany(strings.Split(toMany, ";"), subject, Text(content)); err != nil {
		fmt.Println(err)
		t.Fail()
	}
}
