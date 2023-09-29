package modules

import (
	"database/sql"
	"regate-core/core/cronjob"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"


	_salaR "regate-core/core/modules/sala/repository/pg"
	_salaU "regate-core/core/modules/sala/use_case"

	_authHttpDelivery "regate-core/core/modules/user/auth/delivery/http"
	_authR "regate-core/core/modules/user/auth/repository/pg"
	_authU "regate-core/core/modules/user/auth/usecase"

	_cronjobHttpDelivery "regate-core/core/modules/user/auth/delivery/http"
	_cronjobR "regate-core/core/modules/user/auth/repository/pg"
	_cronjobU "regate-core/core/modules/user/auth/usecase"
)

func InitModules(db *sql.DB){ 
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		// AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept,echo.HeaderAccessControlAllowCredentials},
	}))
	timeout := time.Duration(15) * time.Second

	salaR :=  _salaR.NewRepository(db)
	salaU := _salaU.NewSalaUseCase(salaR,timeout)

	authR :=  _authR.NewRepository(db)
	authU := _authU.NewUseCase(timeout,authR)
	_authHttpDelivery.NewHandler(e,authU)

	
	s := gocron.NewScheduler(time.UTC)
	scheduling.BeginScheduling(s,salaU)
	e.Start("0.0.0.0:8000")
}