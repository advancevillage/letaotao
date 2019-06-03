package services

type Category struct {
	CatID 			int 		`json:"cat_id"`
	PCatID			int 		`json:"p_cat_id"`
	CatNameKey		string		`json:"cat_name_key"`
	CatNameValue	string		`json:"cat_name_key"`
	CreateTime  	string		`json:"create_time"`
	UpdateTime		string		`json:"update_time"`
}

type CategoryService interface {
	Category(cat_id	int) (*Category, error)
	Categories() ([]*Category, error)
	CategoryBy(p_cat_id int) ([]*Category, error)
	CreateCategory(cat *Category) error
	DeleteCategory(cat_id int) error
	CategorySubTree(p_cat_id int) (map[int][]int, error)
	CategoryKey (catID int) (string, error)
}

