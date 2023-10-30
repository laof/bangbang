package tools

import (
	"bangbang/configs"
	"os"

	"github.com/laof/ssdata"
)

func Check() ([]string, []ssdata.List) {

	if force() {
		return configs.All, []ssdata.List{}
	}

	sucList := []ssdata.List{}

	data, err := ssdata.Get("https://laof.github.io/get-nodes-test-app/json/data.json")
	if err != nil {
		return configs.All, sucList
	}

	res := ssdata.PingAll(data, 3)

	for _, item := range data.List {
		if includes(configs.All, item.Name) && !includes(res.Failed, item.Name) {
			sucList = append(sucList, item)
		}
	}

	for _, nnname := range configs.All {
		if !includes(res.Services, nnname) {
			res.Failed = append(res.Failed, nnname)
		}
	}

	return res.Failed, sucList
}

func force() bool {
	f := false
	for _, str := range os.Args {
		if str == "Y" {
			f = true
			break
		}
	}
	return f
}
