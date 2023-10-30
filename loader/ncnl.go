package loader

import (
	configs "bangbang/configs"
	tools "bangbang/tools"
	"fmt"
	"log"

	"github.com/go-rod/rod"
	"github.com/laof/ssdata"
)

type Ncnl struct{}

func (ncnl Ncnl) Run() (info ssdata.List) {
	info = ssdata.List{Name: configs.Ncnl, Remarks: ssdata.ReverseString("科斯莫-斯罗俄")}
	fmt.Println("GO ==> " + info.Name)

	data := ncnl.getData()
	if data.Time != "" {
		// 09月27日 10点15分
		info.Datetime = tools.NcnlTime(data.Time)
	}
	info.Nodes = data.List
	return
}

type Data struct {
	List []string `json:"list"`
	Time string   `json:"time"`
}

func (ncnl Ncnl) getData() Data {

	target := "https://ln" + "cn.org"
	browser := rod.New().MustConnect()
	page := browser.MustPage()
	page.MustNavigate(target)
	page.MustElement(".copy" + "-all")
	// Eval js on the page
	page.MustEval(`() => console.log("hello world")`)
	script := `() => new Promise((resolve) => {
		let timer
		const btn = ['.', 'copy', '-all'].join('')
		const obj = { time: '', list: [] }
		navigator.clipboard.writeText = function (txt) {
		  obj.time = document.querySelector('.time').innerText || ''
		  obj.list = txt.split('\n')
		  return resolve(obj), Promise.resolve('')
		}
		const ele = document.querySelector(btn)
		timer = setInterval(() => ele.click(), 600)
		setTimeout(() => [clearInterval(timer), resolve(obj)], 1000 * 60)
	  })
	`
	promise := page.MustEval(script)

	data := &Data{}

	err := promise.Unmarshal(data)

	if err != nil {
		log.Println("err " + err.Error())
	}
	browser.MustClose()
	return *data
}
