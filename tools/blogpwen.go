package tools

import (
	"bangbang/configs"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/laof/ssdata"
)

func WenGet(filter func(i int, val string) bool) (info ssdata.List) {

	info = ssdata.List{Name: configs.Blogpwen, Remarks: ssdata.ReverseString("亚南东")}

	fmt.Println("GO ==> " + info.Name)

	doc, err := NewQuery("https://www.wen" + "p" + "blog.com")
	if err != nil {
		return
	}

	surl := ""

	// Find the review items
	doc.Find(".mdl-navigation__link").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the title
		target := s.Text()

		if strings.Contains(target, "免"+"费"+"节"+"点") {
			url, ok := s.Attr("href")
			if ok {
				surl = url
			}
		}

	})

	if surl == "" {
		return
	}

	nodedoc, er := NewQuery(surl)
	if er != nil {
		return
	}

	// Find the review items
	a := nodedoc.Find(".article-all a").First()

	detail := ""
	if a != nil {
		val, ok := a.Attr("href")
		if ok {
			detail = val
		}
	}

	if detail == "" {
		return
	}

	dc, err := NewQuery(detail)

	if err != nil {
		return
	}

	txt := ""
	dc.Find("blockquote").Each(func(i int, s *goquery.Selection) {
		txt += s.Text()
	})

	dc.Find("p").Each(func(i int, s *goquery.Selection) {
		txt, word := s.Text(), "最后更新时间："
		if strings.Contains(txt, word) {
			info.Datetime = strings.Replace(txt, word, "", 1)
		}
	})

	list := strings.Split(txt, "\n")

	info.Nodes = ssdata.FilterSlice[string](list, filter)

	return
}

func NewQuery(url string) (*goquery.Document, error) {
	res, err := http.Get(url)
	if err != nil {
		log.Println("load " + url + " error: " + err.Error())
		return &goquery.Document{}, err
	}

	defer res.Body.Close()

	// Load the HTML document
	query, faild := goquery.NewDocumentFromReader(res.Body)
	if faild != nil {
		log.Println("goquery " + url + " error: " + faild.Error())
	}

	return query, faild

}
