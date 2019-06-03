package services

type SPU struct {
	SpuID 		int 	`json:"spu_id"`
	SpuKey		string 	`json:"spu_key"`
	CatID 		int 	`json:"cat_id"`
	CreateTime 	string 	`json:"create_time"`
	UpdateTime 	string  `json:"update_time"`
	BrdID		int 	`json:"brd_id"`
}

type SPUService interface {
	SPU(spu_id int) (*SPU, error)
	SPUs() ([]*SPU, error)
	SPUsBy(catIDs []int) ([]*SPU, error)
	//CreateSPU(spu *SPU) error
	//DeleteSPU(spu_id int) error
}
