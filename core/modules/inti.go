package modules

import (
	"database/sql"
	"time"
	"core_app/core/scheduling"
	"github.com/go-co-op/gocron"
	_salaR "core_app/core/modules/sala/repository/pg"
	_salaU "core_app/core/modules/sala/use_case"
)

func InitModules(db *sql.DB){ 
	timeout := time.Duration(15) * time.Second

	salaR :=  _salaR.NewRepository(db)
	salaU := _salaU.NewSalaUseCase(salaR,timeout)

	s := gocron.NewScheduler(time.UTC)
	scheduling.BeginScheduling(s,salaU)

}