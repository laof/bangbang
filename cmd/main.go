package main

import (
	"bangbang/configs"
	"bangbang/loader/launcher"
	"bangbang/tools"
	v2 "bangbang/v2"
	"fmt"

	"github.com/laof/ssdata"
)

func main() {
	run()
}

func test() {
	n := []string{configs.Blogpwen}

	info := launcher.Start(n)

	fmt.Println(info)

	// tools.CreateTxt(info)

	// if data, err := ssdata.GetDataString(info); err == nil {
	// 	tools.CreateDataFile(data)
	// }

}

func run() {

	arr, exiti := tools.Check()

	tools.GoLoad()

	if len(arr) == 0 {
		tools.Cancel()
		return
	}

	btChan := make(chan bool, 1)
	infoChan := make(chan []ssdata.List, 1)

	go tools.BackupTxtAsync(btChan)
	go launcher.StartAsync(arr, infoChan)

	<-btChan
	info := <-infoChan

	if len(info) == 0 {
		tools.Cancel()
		return
	}

	v2.Start()

	tools.CreateTxt(append(info, exiti...))

	if data, err := ssdata.GetDataString(append(info, exiti...)); err == nil {
		tools.CreateDataFile(data)
		tools.End()
	} else {
		tools.Cancel()
	}

}
