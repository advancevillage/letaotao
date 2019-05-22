package services


type Brand struct {
	BrdID		int		`json:"brd_id"`
	BrdName		string 	`json:"brd_name"`
	BrdCode		int 	`json:"brd_code"`
	CreateTime	string 	`json:"create_time"`
	UpdateTime 	string 	`json:"update_time"`
}


type BrandService interface {
	Brand(brd_id int) (*Brand, error)
	Brands() ([]*Brand, error)
}