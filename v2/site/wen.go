package site

import (
	"bangbang/tools"
	"strings"

	"github.com/laof/ssdata"
)

func Wen(c chan<- ssdata.List) {
	c <- tools.WenGet(func(i int, val string) bool {
		return strings.HasPrefix(val, "vmess://")
	})
}
