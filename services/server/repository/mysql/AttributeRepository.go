//author: richard
package mysql

import (
	"database/sql"
	ls "github.com/advancevillage/letaotao/services"
	lssr "github.com/advancevillage/letaotao/services/server/repository"
	_ "github.com/go-sql-driver/mysql"
)

type AttributeRepository struct {
	DB *sql.DB
}

func (r *AttributeRepository) Attribute(atrbtID int) (*ls.Attribute, error) {
	var table = "attribute"
	var attr  = new(ls.Attribute)
	var err	  error
	var str   = "SELECT atrbt_id,atrbt_key,atrbt_name,atrbt_delete,cat_id,atrbt_create_time,atrbt_last_update_time FROM " + table + " WHERE atrbt_id = ? LIMIT 1"
	stmt, err := r.DB.Prepare(str)
	lssr.Checker(err)
	row := stmt.QueryRow(atrbtID)
	err = row.Scan(&attr.AtrbtID, &attr.AtrbtKey, &attr.AtrbtName, &attr.AtrbtDeleted, &attr.CatID, &attr.CreateTime, &attr.UpdateTime)
	lssr.Checker(err)

	defer func(){
		err := stmt.Close()
		lssr.Checker(err)
	}()

	return attr, err
}
