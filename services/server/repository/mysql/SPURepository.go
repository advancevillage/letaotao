package mysql

import (
	"database/sql"
	la "github.com/advancevillage/letaotao/services"
	lssr "github.com/advancevillage/letaotao/services/server/repository"
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
	lssr.Checker(err)
	row := stmt.QueryRow(spu_id)
	err = row.Scan(&spu.SpuID, &spu.SpuKey, &spu.CatID, &spu.CreateTime, &spu.UpdateTime, &spu.BrdID)
	lssr.Checker(err)

	//@param: 底层报错不直接处理由Wrapper层recover恢复处理
	defer func() {
		err := stmt.Close()
		lssr.Checker(err)
	}()

	return spu, err
}

func (r *SPURepository) SPUs() ([]*la.SPU, error) {
	var table = "spu"
	var spus []*la.SPU
	var str = "SELECT * FROM " + table
	stmt, err := r.DB.Prepare(str)
	lssr.Checker(err)
	rows, err := stmt.Query()
	lssr.Checker(err)
	for rows.Next() {
		var spu = new(la.SPU)
		err = rows.Scan(&spu.SpuID, &spu.SpuKey, &spu.CatID, &spu.CreateTime, &spu.UpdateTime, &spu.BrdID)
		lssr.Checker(err)
		spus = append(spus, spu)
	}

	//@param: 底层报错不直接处理由Wrapper层recover恢复处理
	defer func() {
		err := stmt.Close()
		lssr.Checker(err)
	}()
	defer func() {
		err := rows.Close()
		lssr.Checker(err)
	}()

	return spus, err
}

func (r *SPURepository) SPUsBy(catIDs []int) ([]*la.SPU, error) {
	var table = "spu"
	var spus []*la.SPU
	var temp []interface{}
	var s = ""
	for _, v :=range catIDs {
		temp = append(temp, v)
		s += "?,"
	}
	var str = "SELECT * FROM " + table + " WHERE cat_id in (" + s[0:len(s)-1]  + ")"
	stmt, err := r.DB.Prepare(str)
	lssr.Checker(err)
	rows, err := stmt.Query(temp...)
	lssr.Checker(err)
	for rows.Next() {
		var spu = new(la.SPU)
		err = rows.Scan(&spu.SpuID, &spu.SpuKey, &spu.CatID, &spu.CreateTime, &spu.UpdateTime, &spu.BrdID)
		lssr.Checker(err)
		spus = append(spus, spu)
	}

	//@param: 底层报错不直接处理由Wrapper层recover恢复处理
	defer func() {
		err := stmt.Close()
		lssr.Checker(err)
	}()
	defer func() {
		err := rows.Close()
		lssr.Checker(err)
	}()

	return spus, err
}
