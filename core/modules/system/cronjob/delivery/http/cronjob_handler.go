package http

import (
	"log"
	"net/http"
	r "regate-core/domain/repository"
	_jwt "regate-core/domain/util"

	"github.com/labstack/echo/v4"
)

type CronJobHandler struct {
	cronjobU r.CronJobUseCase
	utilU r.UtilUseCase
}

func NewHandler(e *echo.Echo,cronjobU r.CronJobUseCase,utilU r.UtilUseCase){
	handler := CronJobHandler{
		cronjobU:cronjobU,
		utilU: utilU,
	}
	e.GET("v1/cronjob/list/",handler.GetCronJobs)
	e.POST("v1/cronjob/delete/",handler.DeleteCronJob)
	e.POST("v1/cronjob/start/",handler.StartCronJob)
}

func(h *CronJobHandler)GetCronJobs(c echo.Context)(err error){
	auth := c.Request().Header["Authorization"][0]
	token := _jwt.GetToken(auth)
	_, err = _jwt.ExtractClaims(token)
	if err != nil {
		// log.Println(err)
		return c.JSON(http.StatusUnauthorized, r.ResponseMessage{Message: err.Error()})
	}
	ctx := c.Request().Context()
	res,err := h.cronjobU.GetCronJobs(ctx)
	if err != nil {
		return c.JSON(http.StatusBadRequest, r.ResponseMessage{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, res)
}

func(h *CronJobHandler)DeleteCronJob(c echo.Context)(err error){
	auth := c.Request().Header["Authorization"][0]
	token := _jwt.GetToken(auth)
	_, err = _jwt.ExtractClaims(token)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusUnauthorized, r.ResponseMessage{Message: err.Error()})
	}
	var data r.CronJob
	err = c.Bind(&data)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, r.ResponseMessage{Message: err.Error()})
	}
	ctx := c.Request().Context()
	err = h.cronjobU.DeleteCronJob(ctx,data)
	if err != nil {
	h.utilU.LogError("GetCronJobs","file.tsx",err.Error())
		return c.JSON(http.StatusBadRequest, r.ResponseMessage{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, nil)
}

func(h *CronJobHandler)StartCronJob(c echo.Context)(err error){
	auth := c.Request().Header["Authorization"][0]
	token := _jwt.GetToken(auth)
	_, err = _jwt.ExtractClaims(token)
	if err != nil {
		// log.Println(err)
		return c.JSON(http.StatusUnauthorized, r.ResponseMessage{Message: err.Error()})
	}
	var data r.CronJob
	err = c.Bind(&data)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, r.ResponseMessage{Message: err.Error()})
	}
	ctx := c.Request().Context()
	err = h.cronjobU.StartCronJob(ctx,data)
	if err != nil {
		return c.JSON(http.StatusBadRequest, r.ResponseMessage{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, nil)
}
