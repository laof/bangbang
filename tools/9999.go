package tools

import (
	"regexp"

	"github.com/laof/ssdata"
)

func GetNodesFrom999(url string, paa []string) ssdata.List {

	info := ssdata.List{}
	html, err := LoadHtml(url)
	if err != nil {
		return info
	}

	info.Datetime = ZuluTime(getTime(html))

	for _, val := range paa {
		info.Nodes = append(info.Nodes, getNodes(html, val)...)
	}

	return info
}

func getTime(txt string) string {
	pattern := `(\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}Z)`
	re := regexp.MustCompile(pattern)
	return re.FindString(txt)
}

func getNodes(txt, ss string) []string {
	arr := []string{}
	re := regexp.MustCompile(`<pre>` + ss + `(.*?)</pre>`)
	matches := re.FindAllStringSubmatch(txt, -1)
	for _, match := range matches {
		if len(match) >= 2 {
			arr = append(arr, removeEmail(ss+match[1]))
		}
	}
	return arr
}

func removeEmail(txt string) string {
	// txt := `xxx:<a href="mailto:33333@qq.com">33333@qq.com</a>cc`

	// 正则表达式模式
	pattern := `<a href="mailto:(.*?)">(.*?)</a>`

	// 编译正则表达式
	re := regexp.MustCompile(pattern)

	// 替换匹配的文本
	result := re.ReplaceAllString(txt, "$1")

	return result
}
