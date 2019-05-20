package mysql

import (
	"database/sql"
	"github.com/advancevillage/letaotao"
)

type BrandRepository struct {
	DB	*sql.DB
}

func (r *BrandRepository) Brand(brd_id int) (*letaotao.Brand, error) {
	var table = "brand"
	var brd = new(letaotao.Brand)
	var str = "SELECT * FROM " + table + " WHERE brd_id = ?"
	stmt, err := r.DB.Prepare(str)
	row := stmt.QueryRow(brd_id)
	err = row.Scan(&brd.BrdID, &brd.BrdName, &brd.BrdCode, &brd.CreateTime, &brd.UpdateTime)
	defer stmt.Close()
	return brd, err
}

func (r *BrandRepository) Brands() ([]*letaotao.Brand, error) {
	var table = "brand"
	var brds []*letaotao.Brand
	var str = "SELECT * FROM " + table
	stmt, err := r.DB.Prepare(str)
	rows, err := stmt.Query()
	for rows.Next() {
		var brd = new(letaotao.Brand)
		err = rows.Scan(&brd.BrdID, &brd.BrdName, &brd.BrdCode, &brd.CreateTime, &brd.UpdateTime)
		brds = append(brds, brd)
	}
	return brds, err
}