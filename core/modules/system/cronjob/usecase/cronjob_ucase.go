package usecase

import (
	"context"
	"log"
	r "regate-core/domain/repository"
	"time"

	"github.com/go-co-op/gocron"
)

type cronjobUseCase struct {
	timeout time.Duration
	cronjobR r.CronJobRepository
	salaU r.SalaUseCase
	s *gocron.Scheduler
}

func NewUseCase(timeout time.Duration,cronjobR r.CronJobRepository,
	salaU r.SalaUseCase,s *gocron.Scheduler) r.CronJobUseCase{
	return &cronjobUseCase{
		timeout: timeout,
		cronjobR: cronjobR,
		salaU: salaU,
		s: s,
	}
}

func (u *cronjobUseCase) GetCronJobs(ctx context.Context)(res []r.CronJob,err error){
	ctx,cancel := context.WithTimeout(ctx,u.timeout)
	defer cancel()
	res,err = u.cronjobR.GetCronJobs(ctx)
	return
}

func (u *cronjobUseCase) DeleteCronJob(ctx context.Context,d r.CronJob)(err error){
	err = u.s.RemoveByTag(d.Tag)
	return
}

func (u *cronjobUseCase) StartCronJob(ctx context.Context,d r.CronJob)(err error){
	switch d.Tag {
		case r.DeleteUnAvailablesSalasTag:
			u.s.Cron(d.Expression).Tag(d.Tag).WaitForSchedule().Do(func(){ 
				// s.Every(5).Second().Tag(tag).WaitForSchedule().Do(func(){ 
					log.Println("init again")
					u.salaU.DeleteUnAvailablesSalas(context.Background())
					// salaU.DeleteUnAvailablesSalas(context.Background())
				})
		case r.DisabledExpiredRoomsTag:
			log.Println("init again")
		u.s.Cron(d.Expression).Tag(d.Tag).WaitForSchedule().Do(func(){ 
			// s.Every(5).Second().Tag(tag).WaitForSchedule().Do(func(){ 
			log.Println("init again disabled")
			u.salaU.DisabledExpiredRooms(context.Background())
				// salaU.DeleteUnAvailablesSalas(context.Background())
				})
		default:
			return	
	}
	// err = u.s.RemoveByTag(cronJob.Tag)
	return
}