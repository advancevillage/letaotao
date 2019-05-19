package mysql

import (
	"database/sql"
	"github.com/advancevillage/letaotao"
	_ "github.com/go-sql-driver/mysql"
)

type SPURepository struct {
	DB *sql.DB
}

func (r *SPURepository) SPU(spu_id int) (*letaotao.SPU, error){
	var table = "spu"
	var spu = new(letaotao.SPU)
	var str = "SELECT * FROM " + table + " WHERE spu_id = ?"
	stmt, err := r.DB.Prepare(str)
	row := stmt.QueryRow(spu_id)
	err = row.Scan(&spu.SpuID, &spu.SpuName, &spu.CatID, &spu.CreateTime, &spu.UpdateTime)
	defer  stmt.Close()
	return spu, err
}

func (r *SPURepository) SPUs() ([]*letaotao.SPU, error) {
	var table = "spu"
	var spus []*letaotao.SPU
	var str = "SELECT * FROM " + table
	stmt, err := r.DB.Prepare(str)
	rows, err := stmt.Query()
	for rows.Next() {
		var spu = new(letaotao.SPU)
		err = rows.Scan(&spu.SpuID, &spu.SpuName, &spu.CatID, &spu.CreateTime, &spu.UpdateTime)
		spus = append(spus, spu)
	}
	defer stmt.Close()
	defer rows.Close()
	return spus, err
}