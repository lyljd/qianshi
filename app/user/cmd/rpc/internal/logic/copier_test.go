package logic

import (
	"fmt"
	"github.com/jinzhu/copier"
	"testing"
	"time"
)

type sourInner struct {
	ID       uint
	Birthday time.Time
}

type sour struct {
	sourInner
	Name string
	Age  int
	Tags string
}

type dest struct {
	ID       int64
	Birthday int
	Name     string
	Age      int
	Tags     []string
}

func TestCopier(t *testing.T) {
	s := sour{
		sourInner: sourInner{
			ID:       233,
			Birthday: time.Now(),
		},
		Name: "ljd",
		Age:  18,
		Tags: "sing;dance;rap;basketball",
	}

	var d dest
	if err := copier.Copy(&d, &s); err != nil {
		t.Fail()
		fmt.Println(err)
	}

	fmt.Println(d)
}
