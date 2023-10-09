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

func BeginScheduling(s *gocron.Scheduler,salaU r.SalaUseCase,billingU r.BillingUseCase,
	establecimientoU r.EstablecimientoUseCase) {
	loc,_ := time.LoadLocation("America/La_Paz")
	log.Println(time.Now().In(loc))
	s.Cron("0/15 * * * *").Tag(r.DeleteUnAvailablesSalasTag).WaitForSchedule().Do(func(){ 
			salaU.DeleteUnAvailablesSalas(context.Background())
		})
	s.Cron("0 0 * * *").Tag(r.DisabledExpiredRoomsTag).WaitForSchedule().Do(func(){ 
		    salaU.DisabledExpiredRooms(context.Background())
	})
	s.Cron("0 0 * * *").Tag(r.CreateDepositoTag).WaitForSchedule().Do(func(){ 
		billingU.CreateDepositos(context.Background())
    })
	s.Cron("0 0 * * *").Tag(r.CreateDepositoTag).WaitForSchedule().Do(func(){ 
		log.Println("RUNING JOB")
		establecimientoU.UpdateEstablecimientosTsv(context.Background())
    })

		// s.Every(1).Day().WaitForSchedule().Do(func(){ 
		// 	salaU.DisabledExpiredRooms(context.Background())
		// })
    s.WaitForScheduleAll()
	s.StartAsync()
}