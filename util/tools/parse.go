package tools

import (
	"io/ioutil"
	"net/http"

	"github.com/json-iterator/go"
)

// ParseBody 获取http.Request中的body并解析
func ParseBody(req *http.Request, ins interface{}) {
	body, _ := ioutil.ReadAll(req.Body)
	jsoniter.Unmarshal(body, ins)
}
