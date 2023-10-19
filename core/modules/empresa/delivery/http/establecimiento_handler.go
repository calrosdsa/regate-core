package http

import (
	"log"
	"net/http"
	r "regate-core/domain/repository"
	_jwt "regate-core/domain/util"
	"strconv"

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
	e.GET("v1/empresa/establecimientos/util/",handler.GetUtil)
	e.PUT("v1/core/establecimiento/verificar/:id/",handler.VerificarEstablecimiento)
	e.PUT("v1/core/establecimiento/bloquear/:id/",handler.BloquearEstablecimiento)
}
func (h *EstablecimientoHandler)GetUtil(c echo.Context)(err error){
	categories := []string{"Futbol","Futbol sala"}
	appendTsv(categories)
	ctx := c.Request().Context()
	err = h.establecimientoU.UpdateEstablecimientosTsv(ctx)
	if err != nil {
		return c.JSON(http.StatusBadRequest, r.ResponseMessage{Message: err.Error()})
	}
	return c.JSON(http.StatusOK,"")
}

func appendTsv(values []string){
    var res string
	for _,val := range values {
		t := val + " || "
		res = res + t
		// values = append(values, t)
	}

	log.Println(res)
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

func (h *EstablecimientoHandler)VerificarEstablecimiento(c echo.Context)(err error){
	auth := c.Request().Header["Authorization"][0]
	token := _jwt.GetToken(auth)
	_, err = _jwt.ExtractClaimsAdmin(token)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusUnauthorized, r.ResponseMessage{Message: err.Error()})
	}
	id,err :=strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, r.ResponseMessage{Message: err.Error()})
	}
	ctx := c.Request().Context()
	err = h.establecimientoU.VerificarEstablecimiento(ctx,id)
	 if err != nil {
		return c.JSON(http.StatusBadRequest, r.ResponseMessage{Message: err.Error()})
	}
	return c.JSON(http.StatusOK,err)
}


func (h *EstablecimientoHandler)BloquearEstablecimiento(c echo.Context)(err error){
	auth := c.Request().Header["Authorization"][0]
	token := _jwt.GetToken(auth)
	_, err = _jwt.ExtractClaimsAdmin(token)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusUnauthorized, r.ResponseMessage{Message: err.Error()})
	}
	id,err :=strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, r.ResponseMessage{Message: err.Error()})
	}
	ctx := c.Request().Context()
	err = h.establecimientoU.BloquearEstablecimiento(ctx,id)
	 if err != nil {
		return c.JSON(http.StatusBadRequest, r.ResponseMessage{Message: err.Error()})
	}
	return c.JSON(http.StatusOK,err)
}