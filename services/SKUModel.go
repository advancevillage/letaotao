package services


type SKU struct {
	SkuID		int 	`json:"sku_id"`
	SkuKey		string 	`json:"sku_key"`
	SkuPrice	float32	`json:"sku_price"`
	SpuID		int		`json:"spu_id"`
	BrdID		int 	`json:"brd_id"`
	CreateTime	string 	`json:"create_time"`
	UpdateTime	string 	`json:"update_time"`
	SkuOnSale	byte 	`json:"sku_onsale"`
	DesID	    int 	`json:"des_id"`
	SkuStock	int 	`json:"sku_stock"`
}


type SKUService interface {
	SKU(sku_id int) (*SKU, error)
	SKUs() ([]*SKU, error)
	SKUsBy(spuID int) ([]*SKU, error)
}
