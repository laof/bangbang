package tools

import (
	"bangbang/configs"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path"
	"regexp"
	"runtime"
	"strings"

	"github.com/laof/ssdata"
)

func LoadHtml(url string) (string, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.82 Safari/537.36")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return "", err
	}

	return string(body), nil
}

func CreateDataFile(data string) {
	createfile("data.json", []byte(data))
}

func CreateTxt(data []ssdata.List) {
	list := []string{}
	for _, v := range data {
		if v.Name == configs.Ncnl {
			list = append(list, v.Nodes...)
		}
	}

	if len(list) == 0 {
		return
	}

	CreateSub(list, "ssr.txt")

}

func CreateSub(list []string, name string) {
	str := strings.Join(list, "\n")

	txt := base64.StdEncoding.EncodeToString([]byte(str))

	reg := regexp.MustCompile(`=+$`)
	txt = reg.ReplaceAllString(txt, "")

	createfile(name, []byte(txt))
}

func createfile(filename string, data []byte) {
	// pa, _ := os.Getwd()
	dir := "output"
	_, err := os.Stat(dir)
	if err != nil {
		os.Mkdir(dir, 0700)
	}
	os.WriteFile(path.Join(dir, filename), []byte(data), 0700)
}

func BackupTxtAsync(c chan<- bool) {
	c <- BackupTxt()
	close(c)
}

func BackupTxt() bool {

	resp, err := http.Get("https://laof.github.io/get-nodes-test-app/json/ssr.txt")

	if err != nil {
		return false
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)

	if err != nil {
		return false
	}

	createfile("ssr.txt", data)
	return true
}

func TestFile() {

	//D:\Go\workspace\port  当前项目的路径
	// pa, _ := os.Getwd()
	// fmt.Println("ddd    " + pa)

	str, err := os.ReadFile("output/data.json")

	if err != nil {
		fmt.Println("???================================098")
		fmt.Println(err.Error())
		return
	}
	fmt.Println(string(str))
}

func Cancel() {
	PrintError(" Error Cancel !!! ")
	setGithub("")
}

func End() {
	PrintSuccess("=======  ok  =======")
	setGithub("OK")
}

func setGithub(str string) {

	if runtime.GOOS == "windows" {
		return
	}

	// fmt.Printf("\"%s=%s\" >> $GITHUB_OUTPUT", "result", s) // no no no......

	// # fuck requirement exec.Command!!!
	cmd := exec.Command("/bin/sh", "-c", fmt.Sprintf("echo \"result=%v\" >> $GITHUB_OUTPUT", str))
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		fmt.Printf("error executing shell command: %v", err.Error())
		os.Exit(1)
	}
	os.Exit(0)
}

func includes(arr []string, value string) bool {

	for _, item := range arr {
		if item == value {
			return true
		}
	}
	return false

}

func RemoveDuplicates(arr []string) []string {

	visited := make(map[string]bool)
	var list []string
	for _, str := range arr {
		if visited[str] {
			continue
		}
		visited[str] = true
		list = append(list, str)
	}
	return list

}
