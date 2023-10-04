package usecase

import (
	"context"
	r "regate-core/domain/repository"
	"time"
)

type empresaUseCase struct {
	empresaRepo r.EmpresaRepository
	timeout time.Duration
	utilU r.UtilUseCase
	mediaU r.MediaUseCase
}

func NewUseCase(timeout time.Duration,empresaRepo r.EmpresaRepository,
	utilU r.UtilUseCase,mediaU r.MediaUseCase) r.EmpresaUseCase{
	return &empresaUseCase{
		empresaRepo: empresaRepo,
		timeout: timeout,
		utilU: utilU,
		mediaU: mediaU,
	}
}
func (u *empresaUseCase)GetUuidEmpresa(ctx context.Context,id int)(uuid string,err error){
	ctx,cancel := context.WithTimeout(ctx,u.timeout)
	defer cancel()
	uuid,err = u.empresaRepo.GetUuidEmpresa(ctx,id)
	return
}

func(u *empresaUseCase)CreateEmpresa(ctx context.Context,d *r.Empresa)(err error){
	ctx,cancel := context.WithTimeout(ctx,u.timeout)
	defer cancel()
	err = u.empresaRepo.CreateEmpresa(ctx,d)
	if err != nil{
		u.utilU.LogError("CreaeteEmpresa","empresa_usecase.go",err.Error())
		return 
	}
	err = u.mediaU.CreateBucket(ctx,d.Uuid)
	if err != nil {
		u.utilU.LogError("CreaeteBucket","empresa_usecase.go",err.Error())
	}
	return 
}

func(u *empresaUseCase)GetEmpresasByEstado(ctx context.Context,estado r.EmpresaEstado)(res []r.Empresa,err error){
	ctx,cancel := context.WithTimeout(ctx,u.timeout)
	defer cancel()
	res,err = u.empresaRepo.GetEmpresasByEstado(ctx,estado)
	return 
}