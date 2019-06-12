//author: richard
package services

type SKUAttribute struct {
	 SkuID			int 		`json:"sku_id"`
	 AtrbtID		int 		`json:"atrbt_id"`
	 AtrbtValue 	[]string 	`json:"atrbt_value"`
	 SaDeleted		[]int8 		`json:"sa_delete"`
	 CreateTime 	[]string 	`json:"create_time"`
	 UpdateTime 	[]string 	`json:"update_time"`
}

type SKUAttributeService interface {
	SKUAttributeBySKU(skuID int) ([]*SKUAttribute, error)
	SKUAttributeByAttr(atrbtID int) ([]*SKUAttribute, error)
	SKUAttributeBySKUAndAttr(skuID int, atrbtID int) ([]*SKUAttribute, error)
}


