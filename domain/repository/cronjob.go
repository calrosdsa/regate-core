package repository

import "context"

type CronJobUseCase interface {
	GetCronJobs(ctx context.Context)(res []CronJob,err error)
	DeleteCronJob(ctx context.Context,d CronJob)(err error)
	StartCronJob(ctx context.Context,d CronJob)(err error)
	// GenereateDepositos(ctx context.Context,d CronJob)(err error)
}

type CronJobRepository interface {
	GetCronJobs(ctx context.Context)(res []CronJob,err error)
	DeleteCronJob(ctx context.Context,d CronJob)(err error)
	StartCronJob(ctx context.Context,d CronJob)(err error)
	// GenereateDepositos(ctx context.Context,d CronJob)(err error)
}


type CronJob struct {
	Id int16 `json:"id"`
	Name string `json:"name"`
	Expression string `json:"expression"`
	Tag string `json:"tag"`
}

const (
	DeleteUnAvailablesSalasTag = "delete-unavailables-salas"
	DisabledExpiredRoomsTag = "disable-expired-rooms"
	CreateDepositoTag = "create-deposito"
	UpdateTsvColumn = "update-tsv-column"
)