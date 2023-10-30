package site

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/laof/proxy"
	"github.com/laof/ssdata"
)

func Free(c chan<- ssdata.List) {

	info := ssdata.List{}

	str := proxy.Get("https://vo" + "freed" + "om.org/Home/Get" + "Nodes")

	query, err := goquery.NewDocumentFromReader(strings.NewReader(str))

	if err != nil {
		c <- info
		return
	}

	// Find the review items
	pre := query.Find("pre").First()

	txt := pre.Text()

	info.Nodes = ssdata.FilterSlice[string](strings.Split(txt, "\n"), func(i int, val string) bool {
		return strings.HasPrefix(val, "vmess://")
	})

	c <- info

}
