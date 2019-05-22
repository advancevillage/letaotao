package services

type SPU struct {
	SpuID 		int 	`json:"spu_id"`
	SpuName		string 	`json:"spu_name"`
	CatID 		int 	`json:"cat_id"`
	CreateTime 	string 	`json:"create_time"`
	UpdateTime 	string  `json:"update_time"`
}

type SPUService interface {
	SPU(spu_id int) (*SPU, error)
	SPUs() ([]*SPU, error)
	//CreateSPU(spu *SPU) error
	//DeleteSPU(spu_id int) error
}
