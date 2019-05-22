package mysql

import (
	"database/sql"
	la "github.com/advancevillage/letaotao/services"
	_ "github.com/go-sql-driver/mysql"
)

type CategoryRepository struct {
	DB 	*sql.DB
}

func (r *CategoryRepository) Category(cat_id int) (*la.Category, error) {
	var table = "category"
	var cat = new(la.Category)
	var str = "SELECT * FROM " + table + " WHERE cat_id = ?"
	stmt, err := r.DB.Prepare(str)
	row := stmt.QueryRow(cat_id)
	err = row.Scan(&cat.CatID, &cat.CatName, &cat.CatCode, &cat.CreateTime, &cat.UpdateTime)
	defer stmt.Close()
	return cat , err
}

func (r *CategoryRepository) Categories() ([]*la.Category, error) {
	var table = "category"
	var cats []*la.Category
	var str = "SELECT * FROM " + table
	stmt, err := r.DB.Prepare(str)
	rows, _ := stmt.Query()
	for rows.Next() {
		var cat = new(la.Category)
		err = rows.Scan(&cat.CatID, &cat.CatName, &cat.CatCode, &cat.CreateTime, &cat.UpdateTime)
		cats = append(cats, cat)
	}
	defer stmt.Close()
	defer rows.Close()
	return cats, err
}

func (r *CategoryRepository) CreateCategory(cat *la.Category) error {
	var table = "category"
	var str = "INSERT INTO " + table + "(cat_name, cat_code, cat_create_time)VALUES(?,?,?)"
	stmt, err := r.DB.Prepare(str)
	_, err = stmt.Exec(cat.CatName, cat.CatCode, cat.CreateTime)
	defer  stmt.Close()
	return err
}

func (r *CategoryRepository) DeleteCategory(cat_id int) error {
	var table = "category"
	var str = "DELETE FROM " + table + " WHERE cat_id = ?"
	stmt, err := r.DB.Prepare(str)
	_, err = stmt.Exec(cat_id)
	defer stmt.Close()
	return err
}