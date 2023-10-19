package usecase

import (
	"context"
	"errors"
	r "regate-core/domain/repository"
	"time"
)

type establecimientoUseCase struct {
	establecimientoRepo r.EstablecimientoRepository
	timeout     time.Duration
	utilU       r.UtilUseCase
}

func NewEmpresaUseCase(timeout time.Duration, establecimientoRepo r.EstablecimientoRepository,
	utilU r.UtilUseCase) r.EstablecimientoUseCase {
	return &establecimientoUseCase{
		establecimientoRepo: establecimientoRepo,
		timeout:     timeout,
		utilU:       utilU,
	}
}
func (u *establecimientoUseCase)VerificarEstablecimiento(ctx context.Context,id int)(err error){
	ctx, cancel := context.WithTimeout(ctx, u.timeout)
	defer cancel()
	err = u.establecimientoRepo.UpdateEstadoEstablecimiento(ctx,id,r.EstablecimientoVerificado,true)
	return
}
func (u *establecimientoUseCase)BloquearEstablecimiento(ctx context.Context,id int)(err error){
	ctx, cancel := context.WithTimeout(ctx, u.timeout)
	defer cancel()
	err = u.establecimientoRepo.UpdateEstadoEstablecimiento(ctx,id,r.EstablecimentoBloqueado,false)
	return
}		
func (u *establecimientoUseCase) GetEstablecimientosEmpresa(ctx context.Context,empresaUuid string) (res []r.Establecimiento,
	err error) {
	ctx, cancel := context.WithTimeout(ctx, u.timeout)
	defer cancel()
	res, err = u.establecimientoRepo.GetEstablecimientosEmpresa(ctx,empresaUuid)
	return
}
func (u *establecimientoUseCase) GetEstablecimientosByEstado(ctx context.Context,estado r.EstablecimientoEstado) (res []int,err error) {
	ctx, cancel := context.WithTimeout(ctx, u.timeout)
	defer cancel()
	res,err = u.establecimientoRepo.GetEstablecimientosByEstado(ctx,estado)
	return
}
func (u *establecimientoUseCase) UpdateEstablecimientosTsv(ctx context.Context) (err error) {
	ctx, cancel := context.WithTimeout(ctx, u.timeout)
	defer cancel()
	ids,err :=  u.GetEstablecimientosByEstado(ctx,r.EstablecimientoVerificado)
	err =errors.New("daskdmk samdkasmd kamsd")
	if err != nil {
		u.utilU.LogError("GetDepositosEmpresa","billing_usecase.go",err.Error())
		return 
	}
	for _,id := range ids{
		err = u.establecimientoRepo.UpdateEstablecimientoTsv(ctx,id)
		if err != nil {
			u.utilU.LogError("GetDepositosEmpresa","billing_usecase.go",err.Error())
			return
		}
	}
	return
}