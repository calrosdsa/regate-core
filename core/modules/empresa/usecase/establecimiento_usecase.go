package usecase

import (
	"context"
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

func (u *establecimientoUseCase) GetEstablecimientosEmpresa(ctx context.Context,empresaUuid string) (res []r.Establecimiento,
	err error) {
	ctx, cancel := context.WithTimeout(ctx, u.timeout)
	defer cancel()
	res, err = u.establecimientoRepo.GetEstablecimientosEmpresa(ctx,empresaUuid)
	return
}
