package server

import "github.com/advancevillage/letaotao"

type SPURepository interface {
	SPU(spu_id int) (*letaotao.SPU, error)
	SPUs() ([]*letaotao.SPU, error)
	//CreateSPU(spu *letaotao.SPU) error
	//DeleteSPU(spu_id int) error
}

type SPUService struct {
	Repo SPURepository
}

func (s *SPUService) SPU(spu_id int) (*letaotao.SPU, error) {
	return s.Repo.SPU(spu_id)
}

func (s *SPUService) SPUs() ([]*letaotao.SPU, error) {
	return s.Repo.SPUs()
}