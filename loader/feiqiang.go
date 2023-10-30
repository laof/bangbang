package loader

import (
	configs "bangbang/configs"
	tools "bangbang/tools"
	"fmt"
	"regexp"

	"github.com/laof/ssdata"
)

type Feiqiang struct{}

func (fq Feiqiang) Run() ssdata.List {

	info := ssdata.List{Name: configs.Feiqiang, Remarks: ssdata.ReverseString("知未-他其")}

	fmt.Println("GO ==> " + info.Name)

	url := "https://raw.githubusercontent.com/mian" + "fei" + "fq/share/main/README.md"

	html, err := tools.LoadHtml(url)

	if err != nil {
		return info
	}

	reg := regexp.MustCompile(`\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}\+\d{2}:\d{2}`)
	time := reg.FindString(html)

	// 优化的正则表达式
	reg1 := regexp.MustCompile(`(?m)^(ssr://|ssr://).*[^[:space:]]$`)

	list := reg1.FindAllString(html, -1)

	info.Nodes = list

	// 2023-09-27T17:29:32+08:00
	info.Datetime = tools.RFC3339(time)

	return info

}
