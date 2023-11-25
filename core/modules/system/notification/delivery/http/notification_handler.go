package http

import (
	"log"
	"net/http"
	r "regate-core/domain/repository"
	_jwt "regate-core/domain/util"

	"github.com/labstack/echo/v4"
)

type NotificationHandler struct {
	notificationU r.NotificationuseCase
}

func NewHandler(e *echo.Echo,notificationU r.NotificationuseCase){
	handler := NotificationHandler{
		notificationU:notificationU,
	}
	e.GET("v1/info/:id/",handler.GetInfoText)
	e.POST("v1/notification/diffusion/",handler.SendNotificationDifussion)
}

func (h *NotificationHandler) SendNotificationDifussion(c echo.Context)(err error){
	auth := c.Request().Header["Authorization"][0]
	token := _jwt.GetToken(auth)
	_, err = _jwt.ExtractClaims(token)
	if err != nil {
		// log.Println(err)
		return c.JSON(http.StatusUnauthorized, r.ResponseMessage{Message: err.Error()})
	}
	var data r.RequestNotificationDiffusion
	err = c.Bind(&data)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusUnprocessableEntity, r.ResponseMessage{Message: err.Error()})
	}
	ctx := c.Request().Context()
	h.notificationU.SendNotification(ctx,data)
	log.Println("DIFFUSSION_DATA",data)
	return c.JSON(http.StatusOK,nil)
}

func(h *NotificationHandler)GetInfoText(c echo.Context)(err error){
	// auth := c.Request().Header["Authorization"][0]
	// token := _jwt.GetToken(auth)
	// _, err = _jwt.ExtractClaims(token)
	// if err != nil {
	// 	// log.Println(err)
	// 	return c.JSON(http.StatusUnauthorized, r.ResponseMessage{Message: err.Error()})
	// }
	// id,err := strconv.Atoi(c.Param("id"))
	// if err != nil {
		// return c.JSON(http.StatusBadRequest, r.ResponseMessage{Message: err.Error()})
	// }
	// ctx := c.Request().Context()
	// res,err := h.notificationU.GetInfoText(ctx,id)
	// if err != nil {
	// 	return c.JSON(http.StatusBadRequest, r.ResponseMessage{Message: err.Error()})
	// }
	return c.JSON(http.StatusOK, "")
}

