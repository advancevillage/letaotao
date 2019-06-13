package mysql

import (
	"database/sql"
	la "github.com/advancevillage/letaotao/services"
	lssr "github.com/advancevillage/letaotao/services/server/repository"
	_ "github.com/go-sql-driver/mysql"
)

type BrandRepository struct {
	DB	*sql.DB
}

func (r *BrandRepository) Brand(brd_id int) (*la.Brand, error) {
	var table = "brand"
	var brd = new(la.Brand)
	var str = "SELECT * FROM " + table + " WHERE brd_id = ?"
	stmt, err := r.DB.Prepare(str)
	lssr.Checker(err)
	row := stmt.QueryRow(brd_id)
	err = row.Scan(&brd.BrdID, &brd.BrdName, &brd.BrdCode, &brd.CreateTime, &brd.UpdateTime)
	lssr.Checker(err)

	defer func(){
		err := stmt.Close()
		lssr.Checker(err)
	}()

	return brd, err
}

func (r *BrandRepository) Brands() ([]*la.Brand, error) {
	var table = "brand"
	var brds []*la.Brand
	var str = "SELECT * FROM " + table
	stmt, err := r.DB.Prepare(str)
	lssr.Checker(err)
	rows, err := stmt.Query()
	lssr.Checker(err)
	for rows.Next() {
		var brd = new(la.Brand)
		err = rows.Scan(&brd.BrdID, &brd.BrdName, &brd.BrdCode, &brd.CreateTime, &brd.UpdateTime)
		lssr.Checker(err)
		brds = append(brds, brd)
	}

	defer func(){
		err := stmt.Close()
		lssr.Checker(err)
	}()
	defer func() {
		err := rows.Close()
		lssr.Checker(err)
	}()

	return brds, err
}