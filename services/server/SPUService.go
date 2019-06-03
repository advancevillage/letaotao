package server

import la "github.com/advancevillage/letaotao/services"

type SPURepository interface {
	SPU(spu_id int) (*la.SPU, error)
	SPUs() ([]*la.SPU, error)
	SPUsBy(catIDs []int) ([]*la.SPU, error)
	//CreateSPU(spu *letaotao.SPU) error
	//DeleteSPU(spu_id int) error
}

type SPUService struct {
	Repo SPURepository
}

func (s *SPUService) SPU(spu_id int) (*la.SPU, error) {
	return s.Repo.SPU(spu_id)
}

func (s *SPUService) SPUs() ([]*la.SPU, error) {
	return s.Repo.SPUs()
}

func (s *SPUService) SPUsBy(catIDs []int) ([]*la.SPU, error) {
	return s.Repo.SPUsBy(catIDs)
}