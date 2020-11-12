package types

import (
	"log"
	"net/http"

	"shaos/gateway/code"

	"github.com/json-iterator/go"
)

// AjaxResult
type AjaxResult struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// ToJSON
func (result AjaxResult) ToJSON() []byte {
	res, err := jsoniter.Marshal(result)
	if err != nil {
		return nil
	}
	return res
}

// AjaxResp
func AjaxResp(res http.ResponseWriter, code code.HTTPCode, message string, data interface{}) http.ResponseWriter {
	result := &AjaxResult{
		Code:    int(code),
		Message: message,
		Data:    data,
	}

	if code != http.StatusOK {
		log.Println(message, data)
		result.Data = nil
		code = code / 10000
	}
	res.WriteHeader(int(code))
	_, err := res.Write(result.ToJSON())
	if err != nil {
		code = 400
		result.Data = nil
		log.Printf("wirte response body fail: %v, data: %v\n", err, data)
	} else {
		log.Println("query success")
	}
	return res
}

