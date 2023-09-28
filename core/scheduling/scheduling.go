package scheduling

import (
	"context"
	r "core_app/domain/repository"
	"log"

	"github.com/go-co-op/gocron"
)
type schedulingModule struct {
	s *gocron.Scheduler
}

func BeginScheduling(s *gocron.Scheduler,salaU r.SalaUseCase) {
	s.Every(5).Second().WaitForSchedule().Do(func(){ 
		log.Println("CORE APP ")
		salaU.DeleteUnAvailablesSalas(context.Background())
	 })
	 s.Every(1).Day().WaitForSchedule().Do(func(){ 
		salaU.DisabledExpiredRooms(context.Background())
	 })
	
	 
    s.WaitForScheduleAll()
	s.StartBlocking()
}