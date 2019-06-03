package server

import la "github.com/advancevillage/letaotao/services"

type CategoryRepository interface {
	Category(cat_id int) (*la.Category, error)
	Categories() ([]*la.Category, error)
	CategoryBy(p_cat_id int) ([]*la.Category, error)
	CreateCategory(cat *la.Category) error
	DeleteCategory(cat_id int) error
	CategorySubTree(p_cat_id int) (map[int][]int, error)
	CategoryKey(catID int) (string, error)
}

type CategoryService struct {
	Repo CategoryRepository
}

func (s *CategoryService) Category(cat_id int) (*la.Category, error) {
	return s.Repo.Category(cat_id)
}

func (s *CategoryService) Categories() ([]*la.Category, error) {
	return s.Repo.Categories()
}

func (s *CategoryService) CreateCategory(cat *la.Category) error {
	return s.Repo.CreateCategory(cat)
}

func (s *CategoryService) DeleteCategory(cat_id int) error {
	return s.Repo.DeleteCategory(cat_id)
}

func (s *CategoryService) CategoryBy(p_cat_id int) ([]*la.Category, error) {
	return s.Repo.CategoryBy(p_cat_id)
}

func (s *CategoryService) CategorySubTree(p_cat_id int) (map[int][]int, error) {
	return s.Repo.CategorySubTree(p_cat_id)
}

func (s *CategoryService) CategoryKey (catID int) (string, error) {
	return s.Repo.CategoryKey(catID)
}