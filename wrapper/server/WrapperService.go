package server

import (
	"encoding/json"
	lw "github.com/advancevillage/letaotao/wrapper"
	"net/http"
)

func IndexProcessor(cookies []*http.Cookie) ([]byte, error) {
	var wrapper = new(lw.WrapperModel)
	wrapper.Set("richard", "")
	var response []byte
	response, err := json.Marshal(*wrapper)
	return response, err
}