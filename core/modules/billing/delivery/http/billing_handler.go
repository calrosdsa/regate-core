package http

import (
	"log"
	"net/http"
	r "regate-core/domain/repository"
	_jwt "regate-core/domain/util"
	"strconv"

	"github.com/labstack/echo/v4"
)

type BillingHandler struct {
	billingU r.BillingUseCase
}

func NewHandler(e *echo.Echo,billingU r.BillingUseCase){
	handler := BillingHandler{
		billingU:billingU,
	}
	e.POST("v1/depositos/list/",handler.GetDepositos)
	e.POST("v1/deposito/upload/comprobante/",handler.UploadCombrobanteDeposito)
}

func (h *BillingHandler) UploadCombrobanteDeposito(c echo.Context) (err error) {
	auth := c.Request().Header["Authorization"][0]
	token := _jwt.GetToken(auth)
	_, err = _jwt.ExtractClaimsAdmin(token)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusUnauthorized, r.ResponseMessage{Message: err.Error()})
	}
	var data r.DepositoBancario
	data.DatePaid = c.FormValue("date_paid")
	data.Id,err = strconv.Atoi(c.FormValue("id"))
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, r.ResponseMessage{Message: err.Error()})
	}
	data.ParentId,err = strconv.Atoi(c.FormValue("parentId"))
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, r.ResponseMessage{Message: err.Error()})
	}
	file, err := c.FormFile("file")
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, r.ResponseMessage{Message: err.Error()})
	}
	ctx := c.Request().Context()
	url,err := h.billingU.UploadComprobanteDeposito(ctx,file,data)
	if err != nil {
		return c.JSON(http.StatusBadRequest, r.ResponseMessage{Message: err.Error()})
	}
	// data.ParentId 
	return c.JSON(http.StatusOK, url)
}


func (h *BillingHandler) GetDepositos(c echo.Context) (err error) {
	log.Println("depositos")
	auth := c.Request().Header["Authorization"][0]
	token := _jwt.GetToken(auth)
	_, err = _jwt.ExtractClaimsAdmin(token)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusUnauthorized, r.ResponseMessage{Message: err.Error()})
	}
	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil {
		page = 1
	}
	var data r.DepositoFilterData
	err = c.Bind(&data)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, r.ResponseMessage{Message: err.Error()})
	}
	ctx := c.Request().Context()
	res,count,nextPage, err := h.billingU.GetDepositosEmpresa(ctx,data,int16(page),20)
	if err != nil {
		return c.JSON(http.StatusBadRequest, r.ResponseMessage{Message: err.Error()})
	}
	response := struct {
		NextPage int16       `json:"next_page"`
		Results  []r.DepositoBancarioEmpresa `json:"results"`
		Count    int         `json:"count"`
		PageSize int8        `json:"page_size"`
	}{
		NextPage: nextPage,
		Results:  res,
		Count:    count,
		PageSize: 20,
	}
	return c.JSON(http.StatusOK, response)
}
