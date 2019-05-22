package server

import (
	"fmt"
	lr "github.com/advancevillage/letaotao/restful"
	lws "github.com/advancevillage/letaotao/wrapper/server"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

// HandlerFunc(w http.ResponseWriter, r *http.Request, _ httprouter.Params)
func RoutePolicy() []*lr.Route {
	var policy = []*lr.Route{
		{Method:"GET", Path:"/", HandlerFunc:Index},
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

func Index(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	//url := r.URL
	cookies := r.Cookies()
	//headers := r.Header
	w = SetApiHeader(w)
	response, err := lws.IndexProcessor(cookies)
	fmt.Println(err)
	w.Write(response)
}


