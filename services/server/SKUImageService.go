package server

import la "github.com/advancevillage/letaotao/services"

type SKUImageRepository interface {
	SKUImage(si_id int) (*la.SKUImage, error)
	SKUImages() ([]*la.SKUImage, error)
	SKUImageOfSKU(sku_id int) ([]*la.SKUImage, error)
}

type SKUImageService struct {
	Repo SKUImageRepository
}

func (s *SKUImageService) SKUImage(si_id int) (*la.SKUImage, error) {
	return s.Repo.SKUImage(si_id)
}

func (s *SKUImageService) SKUImages() ([]*la.SKUImage, error) {
	return s.Repo.SKUImages()
}

func (s *SKUImageService) SKUImageOfSKU(sku_id int) ([]*la.SKUImage, error) {
	return s.Repo.SKUImageOfSKU(sku_id)
}