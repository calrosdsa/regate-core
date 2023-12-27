package http

import (
	"net/http"
	r "regate-core/domain/repository"

	_jwt "regate-core/domain/util"

	"github.com/labstack/echo/v4"
)

type LabelHandler struct {
	labelUCase r.LabelUseCase
}
func NewHandler(e *echo.Echo,labelUCase r.LabelUseCase){
	handler := LabelHandler{
		labelUCase:labelUCase,
	}
	e.POST("v1/create-update/label/",handler.CreateOrUpdateLabel)
	e.POST("v1/delete/label",handler.DeleteLabel)
}

func(h *LabelHandler)CreateOrUpdateLabel(c echo.Context)(err error){
	auth := c.Request().Header["Authorization"][0]
	token := _jwt.GetToken(auth)
	_, err = _jwt.ExtractClaims(token)
	if err != nil {
		// log.Println(err)
		return c.JSON(http.StatusUnauthorized, r.ResponseMessage{Message: err.Error()})
	}
	var data r.LabelRequest
	err = c.Bind(&data)
	if err != nil {
		return c.JSON(http.StatusBadRequest, r.ResponseMessage{Message: err.Error()})
	}
	ctx := c.Request().Context()
	err = h.labelUCase.CreateOrUpdateLabel(ctx,data)
	if err != nil {
		return c.JSON(http.StatusBadRequest, r.ResponseMessage{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, nil)
}


func(h *LabelHandler)DeleteLabel(c echo.Context)(err error){
	auth := c.Request().Header["Authorization"][0]
	token := _jwt.GetToken(auth)
	_, err = _jwt.ExtractClaims(token)
	if err != nil {
		// log.Println(err)
		return c.JSON(http.StatusUnauthorized, r.ResponseMessage{Message: err.Error()})
	}
	var data r.LabelRequest
	err = c.Bind(&data)
	if err != nil {
		return c.JSON(http.StatusBadRequest, r.ResponseMessage{Message: err.Error()})
	}
	ctx := c.Request().Context()
	err = h.labelUCase.DeleteLabel(ctx,data)
	if err != nil {
		return c.JSON(http.StatusBadRequest, r.ResponseMessage{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, nil)
}

