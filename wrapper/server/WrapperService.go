package server

import (
	"encoding/json"
	lms "github.com/advancevillage/letaotao/manager/server"
	ls "github.com/advancevillage/letaotao/services"
	lss "github.com/advancevillage/letaotao/services/server"
	lssrr "github.com/advancevillage/letaotao/services/server/repository/redis"
	lw "github.com/advancevillage/letaotao/wrapper"
	"strconv"
)

func CategoryProcessor(cat_id int) (interface{}, error) {
	//@param: cats 表示第一级分类，应用于网站menu导航栏
	var cats []*ls.Category
	var service = &lss.CategoryService{Repo:&lssrr.CategoryRepository{DB:lms.Conn(),Client:lms.ConnCache()}}
	cats,err := service.CategoryBy(cat_id)
	//获取分类树
	//navigation
	//	  |------ main
	//    |------ ...
	var navigate map[string]interface{}
	var main []interface{}
	navigate = make(map[string]interface{})
	main = make([]interface{},len(cats))
	for k,v :=range cats {
		item := make(map[string]string)
		item["key"] = v.CatNameKey
		item["value"] = v.CatNameValue
		item["cat"] = strconv.Itoa(v.CatID)
		main[k] = item
	}
	navigate["main"] = main

	for _, v := range cats {
		var temp []*ls.Category
		temp, err = service.CategoryBy(v.CatID)
		tm := make([]interface{},len(temp))
		for tk, tv :=range  temp {
			item := make(map[string]string)
			item["key"] = tv.CatNameKey
			item["value"] = tv.CatNameValue
			item["cat"] = strconv.Itoa(tv.CatID)
			tm[tk] = item
		}
		if len(temp) > 0 {
			navigate[v.CatNameKey] = tm
		}
	}
	return navigate, err
}


func IndexProcessor(cat_id int) ([]byte, error) {
	var wrapper = new(lw.WrapperModel).Init()
	//@brief: 组装页面接口
	//@interface: navigation   导航栏接口
	//@interface: goods		   商品列表页接口
	//@interface: carts		   购物车接口
	//@interface: logo		   登录接口
	//@interface: users		   用户接口
	navigate , err := CategoryProcessor(cat_id)
	wrapper.Set("navigate", navigate, err)
	var response []byte
	response, err = json.Marshal(*wrapper)
	return response, err
}

