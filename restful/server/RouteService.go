package server

import (
	"fmt"
	lr "github.com/advancevillage/letaotao/restful"
	lws "github.com/advancevillage/letaotao/wrapper/server"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

// HandlerFunc(w http.ResponseWriter, r *http.Request, _ httprouter.Params)
func RoutePolicy() []*lr.Route {
	var policy = []*lr.Route{
		{Method:"GET", Path:"/", HandlerFunc:IndexPage},
	}
	return policy
}

func NewRouter() *httprouter.Router {
	policy := RoutePolicy()
	router := httprouter.New()
	for _, v := range policy {
		router.Handle(v.Method, v.Path, v.HandlerFunc)
	}
	return router
}

func SetApiHeader (w http.ResponseWriter) http.ResponseWriter {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	return  w
}

func IndexPage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	//url := r.URL
	//cookies := r.Cookies()
	qs := r.URL.Query()
	//@param: catID			分类
	//@param: pageID		页面
	var catID,pageID int
	if _, ok := qs["cat"]; ok {
		catID, _ = strconv.Atoi(qs["cat"][0])
	} else {
		catID = 1
	}
	if _, ok := qs["page"]; ok {
		pageID, _ = strconv.Atoi(qs["page"][0])
	} else {
		pageID = 1
	}

	w = SetApiHeader(w)
	response, err := lws.IndexProcessor(catID,pageID)
	fmt.Println(err)
	_, _ = w.Write(response)
}



