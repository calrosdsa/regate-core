package http

import (
	"net/http"
	r "regate-core/domain/repository"

	// _jwt "regate-core/domain/util"
	"strconv"

	"github.com/labstack/echo/v4"
)

type InfoHandler struct {
	infoU r.InfoUseCase
}

func NewHandler(e *echo.Echo,infoU r.InfoUseCase){
	handler := InfoHandler{
		infoU:infoU,
	}
	e.GET("v1/info/:id/",handler.GetInfoText)
}

func(h *InfoHandler)GetInfoText(c echo.Context)(err error){
	// auth := c.Request().Header["Authorization"][0]
	// token := _jwt.GetToken(auth)
	// _, err = _jwt.ExtractClaims(token)
	// if err != nil {
	// 	// log.Println(err)
	// 	return c.JSON(http.StatusUnauthorized, r.ResponseMessage{Message: err.Error()})
	// }
	id,err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, r.ResponseMessage{Message: err.Error()})
	}
	ctx := c.Request().Context()
	res,err := h.infoU.GetInfoText(ctx,id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, r.ResponseMessage{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, res)
}

