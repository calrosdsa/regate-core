package scheduling

import (
	"context"
	"log"
	r "regate-core/domain/repository"
	"time"

	// "time"

	// "time"

	"github.com/go-co-op/gocron"
)
// type schedulingModule struct {
// 	s *gocron.Scheduler
// }

func BeginScheduling(s *gocron.Scheduler,salaU r.SalaUseCase,billingU r.BillingUseCase) {
	loc,_ := time.LoadLocation("America/La_Paz")
	log.Println(time.Now().In(loc))
	s.Cron("0/5 * * * *").Tag(r.DeleteUnAvailablesSalasTag).WaitForSchedule().Do(func(){ 
			salaU.DeleteUnAvailablesSalas(context.Background())
		})
	s.Cron("0/5 * * * *").Tag(r.DisabledExpiredRoomsTag).WaitForSchedule().Do(func(){ 
		    salaU.DisabledExpiredRooms(context.Background())
	})
	s.Cron("0/1 * * * *").Tag(r.CreateDepositoTag).WaitForSchedule().Do(func(){ 
		billingU.CreateDepositos(context.Background())
    })
		// s.Every(1).Day().WaitForSchedule().Do(func(){ 
		// 	salaU.DisabledExpiredRooms(context.Background())
		// })
    s.WaitForScheduleAll()
	s.StartAsync()
}