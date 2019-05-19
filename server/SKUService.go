package server

import "github.com/advancevillage/letaotao"

type SKURepository interface {
	SKU(sku_id int) (*letaotao.SKU, error)
	SKUs() ([]*letaotao.SKU, error)
}


type SKUService struct {
	Repo SKURepository
}

func (s *SKUService) SKU(sku_id int) (*letaotao.SKU, error) {
	return s.Repo.SKU(sku_id)
}

func (s *SKUService) SKUs() ([]*letaotao.SKU, error) {
	return s.Repo.SKUs()
}