package server

import la "github.com/advancevillage/letaotao/services"

type BrandRepository interface {
	Brand(brd_id int) (*la.Brand, error)
	Brands() ([]*la.Brand, error)
}

type BrandService struct {
	Repo BrandRepository
}

func (s *BrandService) Brand(brd_id int) (*la.Brand, error) {
	return s.Repo.Brand(brd_id)
}

func (s *BrandService) Brands() ([]*la.Brand, error) {
	return s.Repo.Brands()
}
