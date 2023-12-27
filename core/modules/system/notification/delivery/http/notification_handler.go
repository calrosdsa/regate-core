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


