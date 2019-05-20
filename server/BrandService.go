package server

import "github.com/advancevillage/letaotao"

type BrandRepository interface {
	Brand(brd_id int) (*letaotao.Brand, error)
	Brands() ([]*letaotao.Brand, error)
}

type BrandService struct {
	Repo BrandRepository
}

func (s *BrandService) Brand(brd_id int) (*letaotao.Brand, error) {
	return s.Repo.Brand(brd_id)
}

func (s *BrandService) Brands() ([]*letaotao.Brand, error) {
	return s.Repo.Brands()
}
