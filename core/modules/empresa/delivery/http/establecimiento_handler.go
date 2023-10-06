package http

import (
	"log"
	"net/http"
	r "regate-core/domain/repository"
	_jwt "regate-core/domain/util"
	// "strconv"

	// "strconv"

	"github.com/labstack/echo/v4"
)

type EstablecimientoHandler struct {
	establecimientoU r.EstablecimientoUseCase
}

func NewEstablecimientoHandler(e *echo.Echo,establecimientoU r.EstablecimientoUseCase){
	handler := EstablecimientoHandler{
		establecimientoU:establecimientoU,
	}
	e.GET("v1/empresa/establecimientos/:empresaUuid/",handler.GetEstablecimientosEmpresa)
}
func (h *EstablecimientoHandler)GetEstablecimientosEmpresa(c echo.Context)(err error){
	auth := c.Request().Header["Authorization"][0]
	token := _jwt.GetToken(auth)
	_, err = _jwt.ExtractClaimsAdmin(token)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusUnauthorized, r.ResponseMessage{Message: err.Error()})
	}
	empresaUuid := c.Param("empresaUuid")
	
	ctx := c.Request().Context()
	res,err := h.establecimientoU.GetEstablecimientosEmpresa(ctx,empresaUuid)
	 if err != nil {
		return c.JSON(http.StatusBadRequest, r.ResponseMessage{Message: err.Error()})
	}
	return c.JSON(http.StatusOK,res)
}

