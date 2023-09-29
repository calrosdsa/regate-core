package repository

import "context"

type CronJobUseCase interface {
	RunDeleteUnAvailablesSalaJob(ctx context.Context)(err error)
	RemoveDeleteUnAvailablesSalaJob(ctx context.Context)(err error)
}
