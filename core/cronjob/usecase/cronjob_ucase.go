package usecase

import (
	"context"
	r "regate-core/domain/repository"
	"time"
)

type cronjobUseCase struct {
	timeout time.Duration
	salaR r.SalaRepository
}

func NewUseCase(timeout time.Duration,salaR r.SalaRepository) r.CronJobUseCase{
	return &cronjobUseCase{
		timeout: timeout,
		salaR: salaR,
	}
}

func (u *cronjobUseCase) RunDeleteUnAvailablesSalaJob(ctx context.Context)(err error){
	return
}

func (u *cronjobUseCase) RemoveDeleteUnAvailablesSalaJob(ctx context.Context)(err error){
	return
}