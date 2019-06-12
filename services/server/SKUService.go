package server

import la "github.com/advancevillage/letaotao/services"

type SKURepository interface {
	SKU(sku_id int) (*la.SKU, error)
	SKUs() ([]*la.SKU, error)
	SKUsBy(spuID int) ([]*la.SKU, error)
	SKUByKey(skuKey string) (*la.SKU, error)
}


type SKUService struct {
	Repo SKURepository
}

func (s *SKUService) SKU(sku_id int) (*la.SKU, error) {
	return s.Repo.SKU(sku_id)
}

func (s *SKUService) SKUs() ([]*la.SKU, error) {
	return s.Repo.SKUs()
}

func (s *SKUService) SKUsBy(spuID int) ([]*la.SKU, error) {
	return s.Repo.SKUsBy(spuID)
}

func (s *SKUService) SKUByKey(skuKey string) (*la.SKU, error) {
	return s.Repo.SKUByKey(skuKey)
}