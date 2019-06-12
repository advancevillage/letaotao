//author: richard
package mysql

import (
	"database/sql"
	ls "github.com/advancevillage/letaotao/services"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
)

type SKUAttributeRepository struct {
	 DB *sql.DB
}


func (r *SKUAttributeRepository) SKUAttributeBySKU(skuID int) ([]*ls.SKUAttribute, error) {
	var table = "sku_attribute"
	var skuAttrs []*ls.SKUAttribute
	var str = "SELECT sku_id, atrbt_id, atrbt_value, sa_delete, sa_create_time, sa_update_time FROM " + table + " WHERE sku_id = ?"
	stmt, err := r.DB.Prepare(str)
	rows, err := stmt.Query(skuID)
	exist := make(map[string]*ls.SKUAttribute)
	for rows.Next() {
		var tDeleted int8
		var tSkuID, tAtrbtID int
		var tAtrbtValue, tCreateTime, tUpdateTime string
		err = rows.Scan(&tSkuID, &tAtrbtID, &tAtrbtValue, &tDeleted, &tCreateTime, &tUpdateTime)
		tKey := strconv.Itoa(tSkuID) + strconv.Itoa(tAtrbtID)
		if _, ok := exist[tKey]; ok {
			skuAttr := exist[tKey]
			skuAttr.AtrbtValue = append(skuAttr.AtrbtValue, tAtrbtValue)
			skuAttr.SaDeleted  = append(skuAttr.SaDeleted, tDeleted)
			skuAttr.CreateTime = append(skuAttr.CreateTime, tCreateTime)
			skuAttr.UpdateTime = append(skuAttr.UpdateTime, tUpdateTime)
		} else {
			var skuAttr = new(ls.SKUAttribute)
			skuAttr.SkuID = tSkuID
			skuAttr.AtrbtID = tAtrbtID
			skuAttr.AtrbtValue = append(skuAttr.AtrbtValue, tAtrbtValue)
			skuAttr.SaDeleted  = append(skuAttr.SaDeleted, tDeleted)
			skuAttr.CreateTime = append(skuAttr.CreateTime, tCreateTime)
			skuAttr.UpdateTime = append(skuAttr.UpdateTime, tUpdateTime)
			exist[tKey] = skuAttr
			skuAttrs = append(skuAttrs, skuAttr)
		}
	}
	defer stmt.Close()
	defer rows.Close()
	return skuAttrs, err
}

func (r *SKUAttributeRepository) SKUAttributeByAttr(atrbtID int) ([]*ls.SKUAttribute, error) {
	var table = "sku_attribute"
	var skuAttrs []*ls.SKUAttribute
	var str = "SELECT sku_id, atrbt_id, atrbt_value, sa_delete, sa_create_time, sa_update_time FROM " + table + " WHERE atrbt_id = ?"
	stmt, err := r.DB.Prepare(str)
	rows, err := stmt.Query(atrbtID)
	exist := make(map[string]*ls.SKUAttribute)
	for rows.Next() {
		var tDeleted int8
		var tSkuID, tAtrbtID int
		var tAtrbtValue, tCreateTime, tUpdateTime string
		err = rows.Scan(&tSkuID, &tAtrbtID, &tAtrbtValue, &tDeleted, &tCreateTime, &tUpdateTime)
		tKey := strconv.Itoa(tSkuID) + strconv.Itoa(tAtrbtID)
		if _, ok := exist[tKey]; ok {
			skuAttr := exist[tKey]
			skuAttr.AtrbtValue = append(skuAttr.AtrbtValue, tAtrbtValue)
			skuAttr.SaDeleted  = append(skuAttr.SaDeleted, tDeleted)
			skuAttr.CreateTime = append(skuAttr.CreateTime, tCreateTime)
			skuAttr.UpdateTime = append(skuAttr.UpdateTime, tUpdateTime)
		} else {
			var skuAttr = new(ls.SKUAttribute)
			skuAttr.SkuID = tSkuID
			skuAttr.AtrbtID = tAtrbtID
			skuAttr.AtrbtValue = append(skuAttr.AtrbtValue, tAtrbtValue)
			skuAttr.SaDeleted  = append(skuAttr.SaDeleted, tDeleted)
			skuAttr.CreateTime = append(skuAttr.CreateTime, tCreateTime)
			skuAttr.UpdateTime = append(skuAttr.UpdateTime, tUpdateTime)
			skuAttrs = append(skuAttrs, skuAttr)
		}
	}
	defer stmt.Close()
	defer rows.Close()
	return skuAttrs, err
}

func (r *SKUAttributeRepository) SKUAttributeBySKUAndAttr(skuId int, atrbtID int) ([]*ls.SKUAttribute, error) {
	var table = "sku_attribute"
	var skuAttrs []*ls.SKUAttribute
	var str = "SELECT sku_id, atrbt_id, atrbt_value, sa_delete, sa_create_time, sa_update_time FROM " + table + " WHERE sku_id = ? AND atrbt_id = ?"
	stmt, err := r.DB.Prepare(str)
	rows, err := stmt.Query(skuId, atrbtID)
	exist := make(map[string]*ls.SKUAttribute)
	for rows.Next() {
		var tDeleted int8
		var tSkuID, tAtrbtID int
		var tAtrbtValue, tCreateTime, tUpdateTime string
		err = rows.Scan(&tSkuID, &tAtrbtID, &tAtrbtValue, &tDeleted, &tCreateTime, &tUpdateTime)
		tKey := strconv.Itoa(tSkuID) + strconv.Itoa(tAtrbtID)
		if _, ok := exist[tKey]; ok {
			skuAttr := exist[tKey]
			skuAttr.AtrbtValue = append(skuAttr.AtrbtValue, tAtrbtValue)
			skuAttr.SaDeleted  = append(skuAttr.SaDeleted, tDeleted)
			skuAttr.CreateTime = append(skuAttr.CreateTime, tCreateTime)
			skuAttr.UpdateTime = append(skuAttr.UpdateTime, tUpdateTime)
		} else {
			var skuAttr = new(ls.SKUAttribute)
			skuAttr.SkuID = tSkuID
			skuAttr.AtrbtID = tAtrbtID
			skuAttr.AtrbtValue = append(skuAttr.AtrbtValue, tAtrbtValue)
			skuAttr.SaDeleted  = append(skuAttr.SaDeleted, tDeleted)
			skuAttr.CreateTime = append(skuAttr.CreateTime, tCreateTime)
			skuAttr.UpdateTime = append(skuAttr.UpdateTime, tUpdateTime)
			skuAttrs = append(skuAttrs, skuAttr)
		}
	}
	defer stmt.Close()
	defer rows.Close()
	return skuAttrs, err
}