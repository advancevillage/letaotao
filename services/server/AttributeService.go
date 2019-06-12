//author: richard
package server

import ls "github.com/advancevillage/letaotao/services"

type AttributeRepository interface {
	Attribute(atrbtID int) (*ls.Attribute, error)
}


type AttributeService struct {
	Repo AttributeRepository
}

func (s *AttributeService) Attribute(atrbtID int) (*ls.Attribute, error) {
	return s.Repo.Attribute(atrbtID)
}