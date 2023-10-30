package launcher

import (
	"bangbang/configs"
	"bangbang/loader"

	"github.com/laof/ssdata"
)

type Loader interface {
	Run() ssdata.List
}

func Start(arr []string) []ssdata.List {
	max := len(arr)
	dc := make(chan ssdata.List, max)
	for _, value := range arr {
		go run(value, dc)
	}

	list := []ssdata.List{}
	for {
		data := <-dc
		list = append(list, data)
		if len(list) >= max {
			close(dc)
			break
		}
	}

	return filterSlice[ssdata.List](list, func(i int, val ssdata.List) bool {
		return len(val.Nodes) > 0
	})
}

func StartAsync(arr []string, c chan<- []ssdata.List) {
	c <- Start(arr)
	close(c)
}

func run(name string, c chan<- ssdata.List) {

	var load Loader

	switch name {
	case configs.Ncnl:
		load = loader.Ncnl{}
	case configs.Feiqiang:
		load = loader.Feiqiang{}
	case configs.Nivla9999:
		load = loader.Nivla9999{}
	case configs.Blogpwen:
		load = loader.Blogpwen{}
	default:
		c <- ssdata.List{}
		return
	}
	c <- load.Run()

}

func filterSlice[T any](arr []T, filter func(i int, val T) bool) []T {
	list := []T{}

	for i, v := range arr {
		if filter(i, v) {
			list = append(list, v)
		}
	}

	return list

}
