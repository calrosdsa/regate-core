package usecase

import (
	"context"
	r "regate-core/domain/repository"
	"time"
)

type salaUseCase struct {
	salaRepo r.SalaRepository
	timeout time.Duration
}

func NewSalaUseCase(salaRepo r.SalaRepository,timeout time.Duration) r.SalaUseCase{
	return &salaUseCase{
		salaRepo: salaRepo,
		timeout: timeout,
	}
}
func (u *salaUseCase)DisabledExpiredRooms(ctx context.Context){
	ctx,cancel := context.WithTimeout(ctx,u.timeout)
	defer cancel()
	u.salaRepo.DisabledExpiredRooms(ctx)
}

func (u *salaUseCase)DeleteUnAvailablesSalas(ctx context.Context){
	ctx,cancel := context.WithTimeout(ctx,u.timeout)
	defer cancel()
	u.salaRepo.DeleteUnAvailablesSalas(ctx)
}