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
		{Method:"GET", Path:"/goods/", HandlerFunc:GoodsPage},
	}
	return policy
}

func NewRouter() (err error){
	policy := RoutePolicy()
	router := httprouter.New()
	for _, v := range policy {
		router.Handle(v.Method, v.Path, v.HandlerFunc)
	}
	err = http.ListenAndServe(":13147", router)
	return
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
	response, err := lws.IndexPageProcessor(catID,pageID)
	//err 记录log
	fmt.Println(err)
	_, _ = w.Write(response)
}

//商品详情页
func GoodsPage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	qs := r.URL.Query()
	//@param: skuKey   商品SKU
	var skuKey string
	if _, ok := qs["sku_key"]; ok {
		skuKey = qs["sku_key"][0]
	} else {
		skuKey = "e411525a2af8b9922a9880a975a62b3e"
	}
	w = SetApiHeader(w)
	response, err := lws.GoodsPageProcessor(skuKey)
	// err 记录log
	fmt.Println(err)
	_, _ = w.Write(response)
}


