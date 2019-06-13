package mysql

import (
	"database/sql"
	la "github.com/advancevillage/letaotao/services"
	lssr "github.com/advancevillage/letaotao/services/server/repository"
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
	lssr.Checker(err)
	row := stmt.QueryRow(sku_id)
	err = row.Scan(&sku.SkuID, &sku.SkuKey, &sku.SkuPrice, &sku.SpuID, &sku.CreateTime, &sku.UpdateTime, &sku.SkuOnSale, &sku.DesID, &sku.SkuStock)
	lssr.Checker(err)

	//@param: 底层报错不直接处理由Wrapper层recover恢复处理
	defer func() {
		err := stmt.Close()
		lssr.Checker(err)
	}()

	return sku, err
}

func (r *SKURepository) SKUs() ([]*la.SKU, error) {
	var table = "sku"
	var skus []*la.SKU
	var str = "SELECT * FROM " + table
	stmt, err := r.DB.Prepare(str)
	lssr.Checker(err)
	rows, err := stmt.Query()
	lssr.Checker(err)
	for rows.Next() {
		var sku = new(la.SKU)
		err = rows.Scan(&sku.SkuID, &sku.SkuKey, &sku.SkuPrice, &sku.SpuID, &sku.CreateTime, &sku.UpdateTime, &sku.SkuOnSale, &sku.DesID, &sku.SkuStock)
		lssr.Checker(err)
		skus = append(skus, sku)
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

	return skus, err
}

func (r *SKURepository) SKUsBy(spuID int) ([]*la.SKU, error) {
	var table = "sku"
	var skus []*la.SKU
	var str = "SELECT * FROM " + table + " WHERE spu_id = ?"
	stmt, err := r.DB.Prepare(str)
	lssr.Checker(err)
	rows, err := stmt.Query(spuID)
	lssr.Checker(err)
	for rows.Next() {
		var sku = new(la.SKU)
		err = rows.Scan(&sku.SkuID, &sku.SkuKey, &sku.SkuPrice, &sku.SpuID, &sku.CreateTime, &sku.UpdateTime, &sku.SkuOnSale, &sku.DesID, &sku.SkuStock)
		lssr.Checker(err)
		skus = append(skus, sku)
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

	return skus, err
}

func (r *SKURepository) SKUByKey(skuKey string) (*la.SKU, error) {
	var table = "sku"
	var sku = new(la.SKU)
	var str = "SELECT * FROM " + table + " WHERE sku_key = ? LIMIT 1"
	stmt, err := r.DB.Prepare(str)
	lssr.Checker(err)
	row := stmt.QueryRow(skuKey)
	err = row.Scan(&sku.SkuID, &sku.SkuKey, &sku.SkuPrice, &sku.SpuID, &sku.CreateTime, &sku.UpdateTime, &sku.SkuOnSale, &sku.DesID, &sku.SkuStock)
	lssr.Checker(err)

	//@param: 底层报错不直接处理由Wrapper层recover恢复处理
	defer func() {
		err := stmt.Close()
		lssr.Checker(err)
	}()

	return sku, err
}