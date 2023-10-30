package loader

import (
	tools "bangbang/tools"
	"strings"

	"github.com/laof/ssdata"
)

type Blogpwen struct{}

func (b Blogpwen) Run() (info ssdata.List) {
	return tools.WenGet(func(i int, val string) bool {
		return strings.HasPrefix(val, "ssr://") || strings.HasPrefix(val, "ss://")
	})
}
