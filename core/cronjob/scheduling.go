package scheduling

import (
	"context"
	r "regate-core/domain/repository"
	"log"
	"time"

	"github.com/go-co-op/gocron"
)
type schedulingModule struct {
	s *gocron.Scheduler
}

func BeginScheduling(s *gocron.Scheduler,salaU r.SalaUseCase) {
	s.Every(5).Second().Tag("tag").WaitForSchedule().Do(func(){ 
		// s.Every(5).Second().Tag("tag").WaitForSchedule().Do(func(){ 
			log.Println("CORE APP ")
			salaU.DeleteUnAvailablesSalas(context.Background())
		})
		
		s.Every(1).Day().WaitForSchedule().Do(func(){ 
			salaU.DisabledExpiredRooms(context.Background())
		})
	go func() {
		err :=  s.RunByTag("tag")
		time.Sleep(6 * time.Second)
		jobs,err := s.FindJobsByTag("tag")
		for _,job := range jobs {
			log.Println(job.IsRunning(),"JOB")
		}
		if err != nil {
			log.Println(err)
		}
	    log.Println(jobs[0].IsRunning())
		time.Sleep(16 * time.Second)
		err = s.RemoveByTag("tag")
		if err != nil {
			log.Println(err)
		}

	    log.Println(jobs[0].IsRunning())

		// job.LimitRunsTo(2)
		
	}()
    s.WaitForScheduleAll()
	s.StartAsync()
}