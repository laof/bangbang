package tools

import (
	"io"
	"log"
	"net/http"
	"strconv"
)

func Keeplive() {
	refresh("https://hwfmd6-8080.csb.app/keeplive")
	refresh("https://nzv26-8080.csb.app/get?target=https://laof.github.io/blob/files/modao/info.json")
}

func refresh(url string) {

	res, err := http.Get(url)

	if err != nil {
		log.Println("Keep live error：" + url + " " + err.Error())
		return
	}

	txt, _ := io.ReadAll(res.Body)

	log.Println("Keep live success：" + strconv.Itoa(len(string(txt))) + " " + url)
	defer res.Body.Close()

}
