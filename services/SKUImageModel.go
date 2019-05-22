package services


type SKUImage struct {
	SiID		int 	`json:"si_id"`
	SiDirection	int8 	`json:"si_direction"`
	SiUrl		string 	`json:"si_url"`
	SiDelete	int8	`json:"si_delete"`
	SiDisplay	int8 	`json:"si_display"`
	SiType		string 	`json:"si_type"`
	CreateTime 	string 	`json:"create_time"`
	UpdateTime 	string 	`json:"update_time"`
	SkuID		int 	`json:"sku_id"`
}

type SKUImageService interface {
	SKUImage(si_id int) (*SKUImage, error)
	SKUImages()	([]*SKUImage, error)
	SKUImageOfSKU(sku_id int) ([]*SKUImage, error)
}