//author: richard
package services

type Attribute struct {
	AtrbtID			int 	`json:"atrbt_id"`
	AtrbtKey		string 	`json:"atrbt_key"`
	AtrbtName		string 	`json:"atrbt_name"`
	AtrbtDeleted	byte	`json:"atrbt_deleted"`
	CatID			int 	`json:"cat_id"`
	CreateTime		string 	`json:"create_time"`
	UpdateTime		string 	`json:"update_time"`
}

type AttributeService interface {
	Attribute(atrbtID int) (*Attribute, error)
}