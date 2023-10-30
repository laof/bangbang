package loader

import (
	"bangbang/configs"
	tools "bangbang/tools"
	"fmt"

	"github.com/laof/ssdata"
)

type Nivla9999 struct{}

func (nivla Nivla9999) Run() ssdata.List {

	info := ssdata.List{Name: configs.Nivla9999, Remarks: ssdata.ReverseString("矶杉洛-国美")}
	fmt.Println("GO ==> " + info.Name)
	url :=
		"https://github.com/Alv" +
			"in9999" +
			"/new-" +
			"pac" +
			"/wiki/ss%E5%85%8D%E8%B4%B9%E8%B4%A6%E5%8F%B7"

	data := tools.GetNodesFrom999(url, []string{"ssr://", "ss://"})

	info.Nodes = data.Nodes
	info.Datetime = data.Datetime

	return info
}
