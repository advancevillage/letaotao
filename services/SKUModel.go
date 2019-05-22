package services


type SKU struct {
	SkuID		int 	`json:"sku_id"`
	SkuName		string 	`json:"sku_name"`
	SkuPrice	float32	`json:"sku_price"`
	SpuID		int		`json:"spu_id"`
	BrdID		int 	`json:"brd_id"`
	CreateTime	string 	`json:"create_time"`
	UpdateTime	string 	`json:"update_time"`
}


type SKUService interface {
	SKU(sku_id int) (*SKU, error)
	SKUs() ([]*SKU, error)
}
