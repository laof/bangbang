package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/laof/ssdata"
)

func main() {

	txt := local()

	if txt == "" {
		log.Println("no data")
		return
	}

	list := strings.Split(txt, "\n")

	// ss := ssdata.FilterSlice[string](list, func(i int, val string) bool {
	// 	return strings.HasPrefix(val, "vmess://")
	// })

	dd, _ := ssdata.Test(strings.Join(list, "\n"))

	for _, val := range dd {

		if val.IsOk {
			log.Println(val.Id)
			log.Println(val.Link)
		}

	}

}

func load() string {
	res, err := http.Get("https://raw.githubusercontent.com/you" + "kai535" + "30100/you" + "kai/master/sub/sub_" + "merge.txt")

	if err != nil {
		return ""
	}

	defer res.Body.Close()

	txt, err := io.ReadAll(res.Body)

	if err != nil {
		return ""
	}

	return string(txt)
}

func local() string {
	data, err := os.ReadFile("test.txt")

	if err != nil {
		log.Println(err.Error())
		return ""
	}

	return string(data)
}
