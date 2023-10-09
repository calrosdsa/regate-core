package modules

import (
	"database/sql"
	"regate-core/core/modules/system/cronjob"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gopkg.in/gomail.v2"

	"github.com/aws/aws-sdk-go/aws/session"

	_salaR "regate-core/core/modules/application/sala/repository/pg"
	_salaU "regate-core/core/modules/application/sala/use_case"

	_authHttpDelivery "regate-core/core/modules/user/auth/delivery/http"
	_authR "regate-core/core/modules/user/auth/repository/pg"
	_authU "regate-core/core/modules/user/auth/usecase"

	_cronjobHttpDelivery "regate-core/core/modules/system/cronjob/delivery/http"
	_cronjobR "regate-core/core/modules/system/cronjob/repository/pg"
	_cronjobU "regate-core/core/modules/system/cronjob/usecase"

	_infoHttpDelivery "regate-core/core/modules/system/info/delivery/http"
	_infoR "regate-core/core/modules/system/info/repository/pg"
	_infoU "regate-core/core/modules/system/info/usecase"

	_billingHttpDelivery "regate-core/core/modules/billing/delivery/http"
	_billingR "regate-core/core/modules/billing/repository/pg"
	_billingU "regate-core/core/modules/billing/usecase"

	_empresaHttpDelivery "regate-core/core/modules/empresa/delivery/http"
	_empresaR "regate-core/core/modules/empresa/repository/pg"
	_empresaU "regate-core/core/modules/empresa/usecase"

	_utilU "regate-core/core/modules/util/usecase"
	_mailU "regate-core/core/modules/util/usecase/mail"
	_mediaU "regate-core/core/modules/util/media/usecase"
)

func InitModules(db *sql.DB,sess *session.Session){ 
	e := echo.New()
	mUser := viper.GetString("mail.user")
	mPass := viper.GetString("mail.password")
	mHost := viper.GetString("mail.host")
	gomailAuth := gomail.NewDialer(mHost, 25, mUser, mPass)
	logger := logrus.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		// AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept,echo.HeaderAccessControlAllowCredentials},
	}))
	loc,_ := time.LoadLocation("America/La_Paz")

	s := gocron.NewScheduler(loc)
	s.TagsUnique()
	timeout := time.Duration(10) * time.Second

	utilU := _utilU.NewUseCase(logger)
	// mailU := _mailU.NewUseCase(gomailAuth)
	_mailU.NewUseCase(gomailAuth)
	mediaU := _mediaU.NewUseCase(sess)

	salaR :=  _salaR.NewRepository(db)
	salaU := _salaU.NewSalaUseCase(salaR,timeout)

	authR :=  _authR.NewRepository(db)
	authU := _authU.NewUseCase(timeout,authR)
	_authHttpDelivery.NewHandler(e,authU)

	cronjobR :=  _cronjobR.NewRepository(db)
	cronjobU := _cronjobU.NewUseCase(timeout,cronjobR,salaU,s)
	_cronjobHttpDelivery.NewHandler(e,cronjobU,utilU)

	infoR :=  _infoR.NewRepository(db)
	infoU := _infoU.NewUseCase(timeout,infoR,utilU)
	_infoHttpDelivery.NewHandler(e,infoU)

	establecimientoR :=  _empresaR.NewEstablecimientoRepository(db)
	establecimientoU := _empresaU.NewEmpresaUseCase(timeout,establecimientoR,utilU)
	_empresaHttpDelivery.NewEstablecimientoHandler(e,establecimientoU)

	empresaR :=  _empresaR.NewRepository(db)
	empresaU := _empresaU.NewUseCase(timeout,empresaR,utilU,mediaU)
	_empresaHttpDelivery.NewHandler(e,empresaU)

	billingR :=  _billingR.NewRepository(db)
	billingU := _billingU.NewUseCase(timeout,billingR,empresaU,utilU,mediaU)
	_billingHttpDelivery.NewHandler(e,billingU)
	
	scheduling.BeginScheduling(s,salaU,billingU,establecimientoU)
	e.Start("0.0.0.0:8000")
}