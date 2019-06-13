package mysql

import (
	"database/sql"
	la "github.com/advancevillage/letaotao/services"
	lssr "github.com/advancevillage/letaotao/services/server/repository"
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
	lssr.Checker(err)
	row := stmt.QueryRow(cat_id)
	err = row.Scan(&cat.CatID, &cat.PCatID, &cat.CatNameKey, &cat.CatNameValue , &cat.CreateTime, &cat.UpdateTime)
	lssr.Checker(err)
	//@param: 底层报错不直接处理由Wrapper层recover恢复处理
	defer func() {
		err := stmt.Close()
		lssr.Checker(err)
	}()
	return cat , err
}

func (r *CategoryRepository) Categories() ([]*la.Category, error) {
	var table = "category"
	var cats []*la.Category
	var str = "SELECT * FROM " + table
	stmt, err := r.DB.Prepare(str)
	lssr.Checker(err)
	rows, err := stmt.Query()
	lssr.Checker(err)
	for rows.Next() {
		var cat = new(la.Category)
		err = rows.Scan(&cat.CatID, &cat.PCatID, &cat.CatNameKey, &cat.CatNameValue , &cat.CreateTime, &cat.UpdateTime)
		lssr.Checker(err)
		cats = append(cats, cat)
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

	return cats, err
}

func (r *CategoryRepository) CreateCategory(cat *la.Category) error {
	var table = "category"
	var str = "INSERT INTO " + table + "(cat_name, cat_code, cat_create_time)VALUES(?,?,?)"
	stmt, err := r.DB.Prepare(str)
	lssr.Checker(err)
	_, err = stmt.Exec(cat.PCatID, cat.CatNameKey, cat.CatNameValue ,cat.CreateTime)
	lssr.Checker(err)

	defer func(){
		err := stmt.Close()
		lssr.Checker(err)
	}()

	return err
}

func (r *CategoryRepository) DeleteCategory(cat_id int) error {
	var table = "category"
	var str = "DELETE FROM " + table + " WHERE cat_id = ?"
	stmt, err := r.DB.Prepare(str)
	lssr.Checker(err)
	_, err = stmt.Exec(cat_id)
	lssr.Checker(err)

	defer func(){
		err := stmt.Close()
		lssr.Checker(err)
	}()

	return err
}

func (r *CategoryRepository) CategoryBy(p_cat_id int) ([]*la.Category, error) {
	var table= "category"
	var str= "SELECT * FROM " + table + " WHERE p_cat_id = ? AND cat_id > 1"
	var cats []*la.Category
	stmt, err := r.DB.Prepare(str)
	lssr.Checker(err)
	rows, err := stmt.Query(p_cat_id)
	lssr.Checker(err)
	for rows.Next() {
		var cat = new(la.Category)
		err = rows.Scan(&cat.CatID, &cat.PCatID, &cat.CatNameKey, &cat.CatNameValue , &cat.CreateTime, &cat.UpdateTime)
		lssr.Checker(err)
		cats = append(cats,cat)
	}

	defer func(){
		err := stmt.Close()
		lssr.Checker(err)
	}()
	defer func() {
		err := rows.Close()
		lssr.Checker(err)
	}()

	return cats, err
}

func (r *CategoryRepository) CategorySubTree(p_cat_id int) (map[int][]int, error) {
	var table = "category"
	var str   = "SELECT cat_id FROM " + table + " WHERE p_cat_id = ? AND cat_id > 1"
	var tree  = make(map[int][]int)
	var i,j   = 0,1
	var queue = make([]int,10)
	var rows  *sql.Rows
	var err   error
	queue[i] = p_cat_id
	stmt, err := r.DB.Prepare(str)
	lssr.Checker(err)

	for ; j > i; {
		p_cat_id = queue[i % cap(queue)]
		i++
		rows, err = stmt.Query(p_cat_id)
		for rows.Next() {
			var t int
			err = rows.Scan(&t)
			lssr.Checker(err)
			tree[p_cat_id] = append(tree[p_cat_id], t)
			queue[j % cap(queue)] = t
			j++
		}
		err = rows.Close()
	}

	defer func(){
		err := stmt.Close()
		lssr.Checker(err)
	}()

	return tree, err
}

func (r *CategoryRepository) CategoryKey (catID int) (string, error) {
	var table = "category"
	var str   = "SELECT cat_name_key FROM " + table + " WHERE cat_id = ? AND cat_id > 1 LIMIT 1"
	var catNameKey = ""
	stmt, err := r.DB.Prepare(str)
	lssr.Checker(err)
	row := stmt.QueryRow(catID)
	err = row.Scan(&catNameKey)
	lssr.Checker(err)

	defer func() {
		err := stmt.Close()
		lssr.Checker(err)
	}()

	return catNameKey , err
}