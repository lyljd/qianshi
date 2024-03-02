package tool

import (
	"fmt"
	"testing"
)

func TestBase64(t *testing.T) {
	s := "testString"
	b := []byte(s)
	encode := Base64Encode(b)
	decode, err := Base64Decode(encode)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	if ds := string(decode); s != ds {
		fmt.Println(s, ds, decode)
		t.Fail()
	}
}

func TestSha256(t *testing.T) {
	s := "testString"
	res, err := Sha256(s, "pass", "hello")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	fmt.Println(res)
}
