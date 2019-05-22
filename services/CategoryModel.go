package services

type Category struct {
	CatID 		int 		`json:"cat_id"`
	CatName		string		`json:"cat_name"`
	CatCode		int 		`json:"cat_code"`
	CreateTime  string		`json:"create_time"`
	UpdateTime	string		`json:"update_time"`
}


type CategoryService interface {
	Category(cat_id	int) (*Category, error)
	Categories() ([]*Category, error)
	CreateCategory(cat *Category) error
	DeleteCategory(cat_id int) error
}

