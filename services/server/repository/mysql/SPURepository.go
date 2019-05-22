package mysql

import (
	"database/sql"
	la "github.com/advancevillage/letaotao/services"
	_ "github.com/go-sql-driver/mysql"
)

type SPURepository struct {
	DB *sql.DB
}

func (r *SPURepository) SPU(spu_id int) (*la.SPU, error){
	var table = "spu"
	var spu = new(la.SPU)
	var str = "SELECT * FROM " + table + " WHERE spu_id = ?"
	stmt, err := r.DB.Prepare(str)
	row := stmt.QueryRow(spu_id)
	err = row.Scan(&spu.SpuID, &spu.SpuName, &spu.CatID, &spu.CreateTime, &spu.UpdateTime)
	defer  stmt.Close()
	return spu, err
}

func (r *SPURepository) SPUs() ([]*la.SPU, error) {
	var table = "spu"
	var spus []*la.SPU
	var str = "SELECT * FROM " + table
	stmt, err := r.DB.Prepare(str)
	rows, err := stmt.Query()
	for rows.Next() {
		var spu = new(la.SPU)
		err = rows.Scan(&spu.SpuID, &spu.SpuName, &spu.CatID, &spu.CreateTime, &spu.UpdateTime)
		spus = append(spus, spu)
	}
	defer stmt.Close()
	defer rows.Close()
	return spus, err
}