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

func TestMD5(t *testing.T) {
	sigKey := "/video/standard/test.mp4-1444435200-0-0-aliyuncdnexp1234"
	res := "23bf85053008f5c0e791667a313e28ce"
	if calcRes := MD5(sigKey); calcRes != res {
		fmt.Println(calcRes)
		t.Fail()
	}
}

func TestEncEmail(t *testing.T) {
	ems := []string{
		"a@qianshi.fun",
		"ab@qianshi.fun",
		"abc@qianshi.fun",
		"abcd@qianshi.fun",
		"abcde@qianshi.fun",
		"abcdef@qianshi.fun",
		"abcdefg@qianshi.fun",
	}

	for _, e := range ems {
		fmt.Println(EncEmail(e))
	}
}
