package http

import (
	"log"
	"net/http"
	r "regate-core/domain/repository"
	_jwt "regate-core/domain/util"
	"strconv"

	// "strconv"

	"github.com/labstack/echo/v4"
)

type EmpresaHandler struct {
	empresaU r.EmpresaUseCase
}

func NewHandler(e *echo.Echo,empresaU r.EmpresaUseCase){
	handler := EmpresaHandler{
		empresaU:empresaU,
	}
	e.GET("v1/empresas/",handler.GetEmpresasByEstado)
	e.POST("v1/empresa/",handler.CreateEmpresa)
}

func (h *EmpresaHandler)CreateEmpresa(c echo.Context)(err error){
	auth := c.Request().Header["Authorization"][0]
	token := _jwt.GetToken(auth)
	_, err = _jwt.ExtractClaimsAdmin(token)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusUnauthorized, r.ResponseMessage{Message: err.Error()})
	}
	ctx := c.Request().Context()
	var data r.Empresa
	err = c.Bind(&data)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusUnauthorized, r.ResponseMessage{Message: err.Error()})
	}
	err = h.empresaU.CreateEmpresa(ctx,&data)
	 if err != nil {
		return c.JSON(http.StatusBadRequest, r.ResponseMessage{Message: err.Error()})
	}
	return c.JSON(http.StatusOK,data)
}


func (h *EmpresaHandler)GetEmpresasByEstado(c echo.Context)(err error){
	auth := c.Request().Header["Authorization"][0]
	token := _jwt.GetToken(auth)
	_, err = _jwt.ExtractClaimsAdmin(token)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusUnauthorized, r.ResponseMessage{Message: err.Error()})
	}
	estado,err := strconv.Atoi(c.QueryParam("estado"))
	if err != nil {
		estado = 1
	}
	ctx := c.Request().Context()
	res,err := h.empresaU.GetEmpresasByEstado(ctx,r.EmpresaEstado(estado))
	 if err != nil {
		return c.JSON(http.StatusBadRequest, r.ResponseMessage{Message: err.Error()})
	}
	return c.JSON(http.StatusOK,res)
}

