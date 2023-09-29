package http

import (
	"net/http"
	r "regate-core/domain/repository"
	_jwt "regate-core/domain/util"

	"github.com/labstack/echo/v4"
)

type CronJobHandler struct {
	cronjobU r.CronJobUseCase
}

func NewHandler(e *echo.Echo,cronjobU r.CronJobUseCase){
	handler := CronJobHandler{
		cronjobU:cronjobU,
	}
	e.GET("v1/cronjob/run/delete-unavailables-salas",handler.RunDeleteUnAvailablesSalaJob)
	e.GET("v1/cronjob/remove/delete-unavailables-salas",handler.RemoveDeleteUnAvailablesSalaJob)

}

func(h *CronJobHandler)RunDeleteUnAvailablesSalaJob(c echo.Context)(err error){
	auth := c.Request().Header["Authorization"][0]
	token := _jwt.GetToken(auth)
	_, err = _jwt.ExtractClaims(token)
	if err != nil {
		// log.Println(err)
		return c.JSON(http.StatusUnauthorized, r.ResponseMessage{Message: err.Error()})
	}
	ctx := c.Request().Context()
	err = h.cronjobU.RunDeleteUnAvailablesSalaJob(ctx)
	if err != nil {
		return c.JSON(http.StatusBadRequest, r.ResponseMessage{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, nil)
}

func(h *CronJobHandler)RemoveDeleteUnAvailablesSalaJob(c echo.Context)(err error){
	auth := c.Request().Header["Authorization"][0]
	token := _jwt.GetToken(auth)
	_, err = _jwt.ExtractClaims(token)
	if err != nil {
		// log.Println(err)
		return c.JSON(http.StatusUnauthorized, r.ResponseMessage{Message: err.Error()})
	}
	ctx := c.Request().Context()
	err = h.cronjobU.RunDeleteUnAvailablesSalaJob(ctx)
	if err != nil {
		return c.JSON(http.StatusBadRequest, r.ResponseMessage{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, nil)
}