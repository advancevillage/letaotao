package server

import (
	"encoding/json"
	lms "github.com/advancevillage/letaotao/manager/server"
	ls "github.com/advancevillage/letaotao/services"
	lss "github.com/advancevillage/letaotao/services/server"
	lssrm "github.com/advancevillage/letaotao/services/server/repository/mysql"
	lw "github.com/advancevillage/letaotao/wrapper"
	"strconv"
)

func CategoryUnit(catID int) (interface{},  error) {
	//@param: cats 表示第一级分类，应用于网站menu导航栏  分类ID的所有信息
	//@param: subCat 以catID 为父分类的子子分类树
	var cats []*ls.Category
	//var service = &lss.CategoryService{Repo:&lssrr.CategoryRepository{DB:lms.Conn(),Client:lms.ConnCache()}}
	var service = &lss.CategoryService{Repo:&lssrm.CategoryRepository{DB:lms.Conn()}}
	cats,err := service.CategoryBy(catID)

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

func GoodsUnit(catID int, pageID int) (map[string]interface{}, error){
	var spuService = &lss.SPUService{Repo:&lssrm.SPURepository{DB:lms.Conn()}}
	var catService = &lss.CategoryService{Repo:&lssrm.CategoryRepository{DB:lms.Conn()}}
	var skuService = &lss.SKUService{Repo:&lssrm.SKURepository{DB:lms.Conn()}}
	var skuImageService = &lss.SKUImageService{Repo:&lssrm.SKUImageRepository{DB:lms.Conn()}}
	var spus     []*ls.SPU
	var skus     []*ls.SKU
	var skuImage []*ls.SKUImage
	var err    error
	var goods  = make(map[string]interface{})

	tree, err  := catService.CategorySubTree(catID)

	for _, v :=range tree[catID] {
		var catIDs []int
		var t []interface{}
		cID := v
		subTree, _    := catService.CategorySubTree(cID)
		catNameKey, _ := catService.CategoryKey(cID)
		catIDs = append(catIDs, cID)
		for _, sv :=range subTree {
			catIDs = append(catIDs, sv...)
		}
		spus , _ = spuService.SPUsBy(catIDs)
		for _, v :=range spus {
			spuID := v.SpuID
			skus, err = skuService.SKUsBy(spuID)
			for _, tv :=range skus {
				var good = make(map[string]interface{})
				skuImage,_ = skuImageService.SKUImageOfSKU(tv.SkuID)
				// 临时取图逻辑
				for _, siv :=range skuImage {
					good["image_url"] = siv.SiUrl
				}
				good["sku_id"] = tv.SkuID
				good["price"]  = tv.SkuPrice
				good["stock"]  = tv.SkuStock
				good["sku_key"] = tv.SkuKey
				t = append(t, good)
			}
		}
		if len(t) > 0 {
			goods[catNameKey] = t
		} else {
			continue
		}
	}
	return goods, err
}
//必须有变量的返回值 response []byte, err error
func IndexPageProcessor(catID int, pageID int) (response []byte, err error) {
	var wrapper = new(lw.WrapperModel).Init()
	defer func() {
		response, err = json.Marshal(*wrapper)
	}()
	defer wrapper.Catcher()

	//@brief: 组装页面接口
	//@interface: navigate     导航栏接口
	//@interface: goods		   商品列表页接口
	//@interface: carts		   购物车接口
	//@interface: logo		   登录接口
	//@interface: users		   用户接口
	navigate, err := CategoryUnit(catID)
	goods, err    := GoodsUnit(catID, pageID)
	wrapper.Set("navigate", navigate, err)
	wrapper.Set("goods", goods, err)

	return
}

func GoodsPageProcessor(skuKey string) (response []byte, err error) {
	var wrapper = new(lw.WrapperModel).Init()
	defer func() {
		response, err = json.Marshal(*wrapper)
	}()
	defer wrapper.Catcher()

	//@interface: sku_attr   商品属性接口
	//@interface: sku		 商品信息接口
	sku, SKUBaseInfo, err := SKUBaseInfoByKeyUnit(skuKey)
	skuAttr, err := SKUAttrBySKUID(SKUBaseInfo)
	wrapper.Set("sku", sku, err)
	wrapper.Set("sku_attr", skuAttr, err)

	return
}
