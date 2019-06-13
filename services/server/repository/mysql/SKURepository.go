package mysql

import (
	"database/sql"
	la "github.com/advancevillage/letaotao/services"
	_ "github.com/go-sql-driver/mysql"
)

type SKURepository struct {
	DB *sql.DB
}


func (r *SKURepository) SKU(sku_id int) (*la.SKU, error) {
	var table = "sku"
	var sku = new(la.SKU)
	var str = "SELECT * FROM " + table + " WHERE sku_id = ?"
	stmt, err := r.DB.Prepare(str)
	defer stmt.Close()
	row := stmt.QueryRow(sku_id)
	err = row.Scan(&sku.SkuID, &sku.SkuKey, &sku.SkuPrice, &sku.SpuID, &sku.CreateTime, &sku.UpdateTime, &sku.SkuOnSale, &sku.DesID, &sku.SkuStock)
	return sku, err
}

func (r *SKURepository) SKUs() ([]*la.SKU, error) {
	var table = "sku"
	var skus []*la.SKU
	var str = "SELECT * FROM " + table
	stmt, err := r.DB.Prepare(str)
	defer stmt.Close()
	rows, err := stmt.Query()
	defer rows.Close()
	for rows.Next() {
		var sku = new(la.SKU)
		err = rows.Scan(&sku.SkuID, &sku.SkuKey, &sku.SkuPrice, &sku.SpuID, &sku.CreateTime, &sku.UpdateTime, &sku.SkuOnSale, &sku.DesID, &sku.SkuStock)
		skus = append(skus, sku)
	}
	return skus, err
}

func (r *SKURepository) SKUsBy(spuID int) ([]*la.SKU, error) {
	var table = "sku"
	var skus []*la.SKU
	var str = "SELECT * FROM " + table + " WHERE spu_id = ?"
	stmt, err := r.DB.Prepare(str)
	defer stmt.Close()
	rows, err := stmt.Query(spuID)
	defer rows.Close()
	for rows.Next() {
		var sku = new(la.SKU)
		err = rows.Scan(&sku.SkuID, &sku.SkuKey, &sku.SkuPrice, &sku.SpuID, &sku.CreateTime, &sku.UpdateTime, &sku.SkuOnSale, &sku.DesID, &sku.SkuStock)
		skus = append(skus, sku)
	}
	return skus, err
}

func (r *SKURepository) SKUByKey(skuKey string) (*la.SKU, error) {
	var table = "sku"
	var sku = new(la.SKU)
	var str = "SELECT * FROM " + table + " WHERE sku_key = ? LIMIT 1"
	stmt, err := r.DB.Prepare(str)
	defer stmt.Close()
	row := stmt.QueryRow(skuKey)
	err = row.Scan(&sku.SkuID, &sku.SkuKey, &sku.SkuPrice, &sku.SpuID, &sku.CreateTime, &sku.UpdateTime, &sku.SkuOnSale, &sku.DesID, &sku.SkuStock)
	return sku, err
}