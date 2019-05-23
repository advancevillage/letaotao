package server

import (
	"encoding/json"
	"fmt"
	lms "github.com/advancevillage/letaotao/manager/server"
	lss "github.com/advancevillage/letaotao/services/server"
	lssrm "github.com/advancevillage/letaotao/services/server/repository/mysql"
	lw "github.com/advancevillage/letaotao/wrapper"

	"net/http"
)

func IndexProcessor(cookies []*http.Cookie) ([]byte, error) {
	var wrapper = new(lw.WrapperModel)
	wrapper.Set("richard", "")
	var sis = &lss.SKUImageService{Repo:&lssrm.SKUImageRepository{DB:lms.Conn()}}
	si, err := sis.SKUImage(1)
	fmt.Println(si)
	//wrapper.Set(si, err)
	var response []byte
	response, err = json.Marshal(*wrapper)
	return response, err
}