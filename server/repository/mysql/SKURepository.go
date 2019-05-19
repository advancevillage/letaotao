package mysql

import (
	"database/sql"
	"github.com/advancevillage/letaotao"
	_ "github.com/go-sql-driver/mysql"
)

type SKURepository struct {
	DB *sql.DB
}


func (r *SKURepository) SKU(sku_id int) (*letaotao.SKU, error) {
	var table = "sku"
	var sku = new(letaotao.SKU)
	var str = "SELECT * FROM " + table + " WHERE sku_id = ?"
	stmt, err := r.DB.Prepare(str)
	row := stmt.QueryRow(sku_id)
	err = row.Scan(&sku.SkuID, &sku.SkuName, &sku.SpuID, &sku.BrdID, &sku.CreateTime, &sku.UpdateTime)
	defer stmt.Close()
	return sku, err
}

func (r *SKURepository) SKUs() ([]*letaotao.SKU, error) {
	var table = "sku"
	var skus []*letaotao.SKU
	var str = "SELECT * FROM " + table
	stmt, err := r.DB.Prepare(str)
	rows, err := stmt.Query(str)
	for rows.Next() {
		var sku = new(letaotao.SKU)
		err = rows.Scan(&sku.SkuID, &sku.SkuName, &sku.SpuID, &sku.BrdID, &sku.CreateTime, &sku.UpdateTime)
		skus = append(skus, sku)
	}
	return skus, err
}