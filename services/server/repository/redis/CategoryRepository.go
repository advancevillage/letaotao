//author: richard
package redis

import (
	"database/sql"
	la "github.com/advancevillage/letaotao/services"
	lssrm "github.com/advancevillage/letaotao/services/server/repository/mysql"
	"github.com/go-redis/redis"
)

type CategoryRepository struct {
	Client 	*redis.Client
	DB 		*sql.DB
}

func (r *CategoryRepository) Category(cat_id int) (*la.Category, error) {
	var repo = lssrm.CategoryRepository{DB:r.DB}
	return repo.Category(cat_id)
}

func (r *CategoryRepository) Categories() ([]*la.Category, error) {
	var repo = lssrm.CategoryRepository{DB:r.DB}
	return repo.Categories()
}

func (r *CategoryRepository) CategoryBy(p_cat_id int) ([]*la.Category, error) {
	var repo = lssrm.CategoryRepository{DB:r.DB}
	return repo.CategoryBy(p_cat_id)
}

func (r *CategoryRepository) CreateCategory(cat *la.Category) error {
	var repo = lssrm.CategoryRepository{DB:r.DB}
	return repo.CreateCategory(cat)
}

func (r *CategoryRepository) DeleteCategory(cat_id int) error {
	var repo = lssrm.CategoryRepository{DB:r.DB}
	return repo.DeleteCategory(cat_id)
}

