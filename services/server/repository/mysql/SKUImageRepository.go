package mysql

import (
	"database/sql"
	la "github.com/advancevillage/letaotao/services"
	_ "github.com/go-sql-driver/mysql"
)

type SKUImageRepository struct {
	DB *sql.DB
}

func (r *SKUImageRepository) SKUImage(si_id int) (*la.SKUImage, error) {
	var table = "sku_image"
	var si = new(la.SKUImage)
	var str = "SELECT * FROM " + table + " WHERE si_id = ?"
	stmt, err := r.DB.Prepare(str)
	row := stmt.QueryRow(si_id)
	err = row.Scan(&si.SiID, &si.SiDirection, &si.SiUrl, &si.SiDelete, &si.SiDisplay, &si.SiType, &si.CreateTime, &si.UpdateTime, &si.SkuID)
	defer stmt.Close()
	return si, err
}

func (r *SKUImageRepository) SKUImages() ([]*la.SKUImage, error) {
	var table = "sku_image"
	var sis []*la.SKUImage
	var str = "SELECT * FROM " + table
	stmt, err := r.DB.Prepare(str)
	rows, err := stmt.Query()
	for rows.Next() {
		var si = new(la.SKUImage)
		err = rows.Scan(&si.SiID, &si.SiDirection, &si.SiUrl, &si.SiDelete, &si.SiDisplay, &si.SiType, &si.CreateTime, &si.UpdateTime, &si.SkuID)
		sis = append(sis, si)
	}
	defer stmt.Close()
	defer rows.Close()
	return sis, err
}

func (r *SKUImageRepository) SKUImageOfSKU(sku_id int) ([]*la.SKUImage, error) {
	var table = "sku_image"
	var sis []*la.SKUImage
	var str = "SELECT * FROM " + table + " WHERE sku_id = ? AND si_display = 1 AND si_delete = 0"
	stmt, err := r.DB.Prepare(str)
	rows, err := stmt.Query(sku_id)
	for rows.Next() {
		var si = new(la.SKUImage)
		err = rows.Scan(&si.SiID, &si.SiDirection, &si.SiUrl, &si.SiDelete, &si.SiDisplay, &si.SiType, &si.CreateTime, &si.UpdateTime, &si.SkuID)
		sis = append(sis, si)
	}
	return sis, err
}