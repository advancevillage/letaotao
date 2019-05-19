package server

import "github.com/advancevillage/letaotao"

type CategoryRepository interface {
	Category(cat_id int) (*letaotao.Category, error)
	Categories() ([]*letaotao.Category, error)
	CreateCategory(cat *letaotao.Category) error
	DeleteCategory(cat_id int) error
}

type CategoryService struct {
	Repo CategoryRepository
}

func (s *CategoryService) Category(cat_id int) (*letaotao.Category, error) {
	return s.Repo.Category(cat_id)
}

func (s *CategoryService) Categories() ([]*letaotao.Category, error) {
	return s.Repo.Categories()
}

func (s *CategoryService) CreateCategory(cat *letaotao.Category) error {
	return s.Repo.CreateCategory(cat)
}

func (s *CategoryService) DeleteCategory(cat_id int) error {
	return s.Repo.DeleteCategory(cat_id)
}