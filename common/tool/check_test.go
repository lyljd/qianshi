package tool

import (
	"fmt"
	"testing"
)

func TestCheckStrSliDup(t *testing.T) {
	ss := []string{"你好", "nh", "OK", "ok", "18", "OK", "你好"}
	fmt.Println(CheckStrSliDup(ss))
}

func TestCheckDateStr(t *testing.T) {
	ds := []string{"20220305", "20220230", "01240101", "202408311", "abcd", "2001030", "20301101", "20240307"}
	for _, s := range ds {
		fmt.Println(s, CheckDateStr(s))
	}
}

func TestCheckEmail(t *testing.T) {
	ems := []string{"ljd@qianshi.fun", "a@host.b", "1234"}
	for _, e := range ems {
		fmt.Println(e, CheckEmail(e))
	}
}
