//author: richard
package server

import (
	lms "github.com/advancevillage/letaotao/manager/server"
	ls "github.com/advancevillage/letaotao/services"
	lss "github.com/advancevillage/letaotao/services/server"
	lssrm "github.com/advancevillage/letaotao/services/server/repository/mysql"
)

func SKUBaseInfoByKeyUnit(skuKey string) (map[string]interface{}, *ls.SKU, error) {
	var skuService = &lss.SKUService{Repo:&lssrm.SKURepository{DB:lms.Conn()}}
	var sku 	*ls.SKU
	var err 	error
	var skuBaseInfo = make(map[string]interface{})
	sku, err = skuService.SKUByKey(skuKey)

	skuBaseInfo["SkuID"] 		= sku.SkuID
	skuBaseInfo["SkuKey"] 		= sku.SkuKey
	skuBaseInfo["SkuPrice"] 	= sku.SkuPrice
	skuBaseInfo["BrdID"] 		= sku.BrdID
	skuBaseInfo["SkuOnSale"] 	= sku.SkuOnSale
	skuBaseInfo["DesID"] 		= sku.DesID
	skuBaseInfo["SkuStock"] 	= sku.SkuStock
	skuBaseInfo["CreateTime"] 	= sku.CreateTime
	skuBaseInfo["UpdateTime"] 	= sku.UpdateTime

	return skuBaseInfo, sku, err
}

func SKUAttrBySKUID(skuBaseInfo *ls.SKU) (map[string]interface{}, error){
	var skuAttrService = lss.SKUAttributeService{Repo:&lssrm.SKUAttributeRepository{DB:lms.Conn()}}
	var attrService = lss.AttributeService{Repo:&lssrm.AttributeRepository{DB:lms.Conn()}}
	var skuAttrsInfo   = make(map[string]interface{})
	var skuAttrs  []*ls.SKUAttribute
	var attValue  *ls.Attribute
	var err 	  error
	skuAttrs, err = skuAttrService.SKUAttributeBySKU(skuBaseInfo.SkuID)
	for _, v :=range skuAttrs {
		attValue, err = attrService.Attribute(v.AtrbtID)
		skuAttrsInfo[attValue.AtrbtName] = v.AtrbtValue
	}
	return skuAttrsInfo, err
}

