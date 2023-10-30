package v2

import (
	"bangbang/tools"
	"bangbang/v2/site"
	"encoding/base64"
	"encoding/json"
	"strings"

	"github.com/laof/ssdata"
)

type VmessNode struct {
	Add            string `json:"add"`
	Aid            string `json:"aid"`
	Encryption     string `json:"encryption"`
	Host           string `json:"host"`
	ID             string `json:"id"`
	Net            string `json:"net"`
	Path           string `json:"path"`
	Port           string `json:"port"`
	Ps             string `json:"ps"` //remark
	Security       string `json:"security"`
	SkipCertVerify bool   `json:"skip-cert-verify"`
	TLS            string `json:"tls"`
	Type           string `json:"type"`
	URLGroup       string `json:"url_group"` // group
	V              string `json:"v"`
}

func Start() {

	c := make(chan ssdata.List, 3)

	go site.Free(c)
	go site.A999(c)
	go site.Wen(c)

	list := []ssdata.List{}
	txt := []string{}
	for {
		data := <-c

		list = append(list, data)

		for _, val := range data.Nodes {

			link := updateNode(val)
			if link == "" {
				txt = append(txt, val)
			} else {
				txt = append(txt, link)
			}
		}

		if len(list) >= 3 {
			close(c)
			break
		}

	}

	tools.CreateSub(txt, "v2.txt")

}

func updateNode(encodedNode string) string {
	decoded, err := base64.StdEncoding.DecodeString(encodedNode[8:])

	if err != nil {
		return ""
	}

	data := &VmessNode{}

	json.Unmarshal(decoded, data)

	newName := rename(*data)

	if newName == "" {
		return ""
	}

	data.Ps = "Smiles" + "_专用_" + strings.TrimSpace(newName)

	link, err := json.Marshal(data)

	if err != nil {
		return ""
	}

	return "vmess://" + base64.StdEncoding.EncodeToString(link)
}

func rename(node VmessNode) string {

	subtxt := "https://github.com/Alv" + "in99" + "99/new" + "-" + "pac/wiki/v2r" + "ay%E5%85%8D%E8%B4%B9%E8%B4%A6%E5%8F%B7"
	if strings.Contains(node.Ps, subtxt) {
		return strings.Replace(node.Ps, subtxt, "", 1)
	}
	return ""

}
