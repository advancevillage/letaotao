package server

import "github.com/advancevillage/letaotao"

type SKUImageRepository interface {
	SKUImage(si_id int) (*letaotao.SKUImage, error)
	SKUImages() ([]*letaotao.SKUImage, error)
	SKUImageOfSKU(sku_id int) ([]*letaotao.SKUImage, error)
}

type SKUImageService struct {
	Repo SKUImageRepository
}

func (s *SKUImageService) SKUImage(si_id int) (*letaotao.SKUImage, error) {
	return s.Repo.SKUImage(si_id)
}

func (s *SKUImageService) SKUImages() ([]*letaotao.SKUImage, error) {
	return s.Repo.SKUImages()
}

func (s *SKUImageService) SKUImageOfSKU(sku_id int) ([]*letaotao.SKUImage, error) {
	return s.Repo.SKUImageOfSKU(sku_id)
}