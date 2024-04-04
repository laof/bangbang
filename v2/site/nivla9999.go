package site

import (
	"bangbang/tools"

	"github.com/laof/ssdata"
)

func A999(c chan<- ssdata.List) {
	url := "https://github.com/Alv" + "in9999/new" + "-" + "pac/wiki/v2ray%E5%85%8D%E8%B4%B9%E8%B4%A6%E5%8F%B7"
	c <- tools.GetNodesFrom999(url, []string{"vmess://"})
}

func A999r(c chan<- ssdata.List) {
	url := "https://github.com/Alv" + "in9999/new-" + "pac/wiki/ss%E5%85%8D%E8%B4%B9%E8%B4%A6%E5%8F%B7"
	c <- tools.GetNodesFrom999(url, []string{"ssr://", "ss://"})
}
