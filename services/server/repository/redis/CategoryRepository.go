//author: richard
package redis

import (
	"database/sql"
	"encoding/json"
	"fmt"
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
	var key = fmt.Sprintf("%s_%s_%s_%d", "001", "category", "pcatid", p_cat_id)
	var cats   []*la.Category
	var values []string
	var value  []byte

	ret, err := r.Client.Exists(key).Result()
	if  ret > 0 {
		values, err = r.Client.SMembers(key).Result()
		pack := make(map[string]interface{})
		for _, v :=range  values {
			var cat = new(la.Category)
			err = json.Unmarshal([]byte(v), &pack)
			var CatID,PCatID float64
			CatID, 		      _ = pack["CatID"].(float64)
			PCatID,		      _ = pack["PCatID"].(float64)
			cat.CatID		    = int(CatID)
			cat.PCatID  	    = int(PCatID)
			cat.CatNameKey,   _ = pack["CatNameKey"].(string)
			cat.CatNameValue, _ = pack["CatNameValue"].(string)
			cat.CreateTime,   _ = pack["CreateTime"].(string)
			cat.UpdateTime,   _ = pack["UpdateTime"].(string)
			cats = append(cats, cat)
		}
	} else {
		repo := lssrm.CategoryRepository{DB:r.DB}

		cats,err = repo.CategoryBy(p_cat_id)
		pack := make(map[string]interface{})
		for _,v :=range cats {
			pack["CatID"] 		 = v.CatID
			pack["PCatID"] 		 = v.PCatID
			pack["CatNameKey"] 	 = v.CatNameKey
			pack["CatNameValue"] = v.CatNameValue
			pack["CreateTime"] 	 = v.CreateTime
			pack["UpdateTime"] 	 = v.UpdateTime
			value, err = json.Marshal(pack)
			r.Client.SAdd(key,string(value))
		}
	}
	return cats, err
}

func (r *CategoryRepository) CreateCategory(cat *la.Category) error {
	var repo = lssrm.CategoryRepository{DB:r.DB}
	return repo.CreateCategory(cat)
}

func (r *CategoryRepository) DeleteCategory(cat_id int) error {
	var repo = lssrm.CategoryRepository{DB:r.DB}
	return repo.DeleteCategory(cat_id)
}

