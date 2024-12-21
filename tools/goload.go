package tools

import (
	"io"
	"log"
	"net/http"
)

func GoLoad() {

	res, err := http.Get("https://hwfmd6-8080.csb.app/keeplive")

	if err != nil {
		log.Println("Error hwfmd6 : " + err.Error())
		return
	}

	txt, err := io.ReadAll(res.Body)
	log.Println("ok hwfmd6: " + string(txt))
	defer res.Body.Close()

}
