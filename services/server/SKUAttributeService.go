//author: richard
package server

import la "github.com/advancevillage/letaotao/services"

type SKUAttributeRepository interface {
	SKUAttributeBySKU(skuID int) ([]*la.SKUAttribute, error)
	SKUAttributeByAttr(atrbtID int) ([]*la.SKUAttribute, error)
	SKUAttributeBySKUAndAttr(skuID int, atrbtID int) ([]*la.SKUAttribute, error)
}

type SKUAttributeService struct {
	Repo SKUAttributeRepository
}

func (s *SKUAttributeService) SKUAttributeBySKU(skuID int) ([]*la.SKUAttribute, error) {
	return s.Repo.SKUAttributeBySKU(skuID)
}

func (s *SKUAttributeService) SKUAttributeByAttr(atrbtID int) ([]*la.SKUAttribute, error) {
	return s.Repo.SKUAttributeByAttr(atrbtID)
}

func (s *SKUAttributeService) SKUAttributeBySKUAndAttr(skuID int, atrbtID int) ([]*la.SKUAttribute, error) {
	return s.Repo.SKUAttributeBySKUAndAttr(skuID, atrbtID)
}