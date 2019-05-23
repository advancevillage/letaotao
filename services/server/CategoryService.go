package server

import la "github.com/advancevillage/letaotao/services"

type CategoryRepository interface {
	Category(cat_id int) (*la.Category, error)
	Categories() ([]*la.Category, error)
	CreateCategory(cat *la.Category) error
	DeleteCategory(cat_id int) error
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