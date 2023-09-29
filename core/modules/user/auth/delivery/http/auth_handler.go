package http

import (
	// "context"
	"log"
	"net/http"
	r "regate-core/domain/repository"

	// "strconv"

	// "strings"

	_jwt "regate-core/domain/util"

	"github.com/labstack/echo/v4"
)

type AccoutnHandler struct {
	authUseCase r.AuthUseCase
	// mailU r.MailUseCase
}

type Res struct {
	Result string `json:"result"`
}

func NewHandler(e *echo.Echo, authUseCase r.AuthUseCase) {
	handler := &AccoutnHandler{
		authUseCase: authUseCase,
		// mailU: mailU,
	}
	e.POST("v1/auth/login/", handler.Login)
	e.GET("v1/auth/", handler.Me)

	e.GET("v1/auth/send-again-email-verfication/:id/", handler.SendEmailVerification)
	e.GET("v1/auth/verify-email/:userId/:otp/", handler.VerifyEmail)
}

func (h *AccoutnHandler) VerifyEmail(c echo.Context) (err error) {
	// userId,_ := strconv.Atoi(c.Param("userId"))
	// otp,_ := strconv.Atoi(c.Param("otp"))
	// ctx := c.Request().Context()
	// res,err := h.authUseCase.VerifyEmail(ctx,userId,otp)
	// if err != nil {
	// 	log.Println(err)
	// 	return c.JSON(http.StatusBadRequest, r.ResponseMessage{Message: err.Error()})
	// }
	// token, err := _jwt.GenerateJWT(res.UserId, *res.Email, *res.Username,res.ProfileId)
	// if err != nil {
	// 	log.Println(err)
	// 	return c.JSON(http.StatusExpectationFailed, r.ResponseMessage{Message: err.Error()})
	// }
	// response := r.AuthenticationResponse{
	// 	User:  res,
	// 	Token: token,
	// }
	return c.JSON(http.StatusOK, "")
}

func (h *AccoutnHandler) SendEmailVerification(c echo.Context) (err error) {
	// auth := c.Request().	Header["Authorization"][0]
	// token := _jwt.GetToken(auth)
	// claims, err := _jwt.ExtractClaims(token)
	// if err != nil {
	// 	// log.Println(err)
	// 	return c.JSON(http.StatusUnauthorized, r.ResponseMessage{Message: err.Error()})
	// }
	// id ,_:=  strconv.Atoi(c.Param("id"))
	// ctx := c.Request().Context()
	// res,err := h.authUseCase.GetUser(ctx,id)
	// if err != nil {
	// 	log.Println(err)
	// 	return c.JSON(http.StatusBadRequest, r.ResponseMessage{Message: err.Error()})
	// }
	// h.mailU.SendEmailVerification(*res.Email,*res.Otp,*res.Username)
	// return
	return
}

func (h *AccoutnHandler) Me(c echo.Context) (err error) {
	auth := c.Request().Header["Authorization"][0]
	token := _jwt.GetToken(auth)
	claims, err := _jwt.ExtractClaims(token)
	if err != nil {
		// log.Println(err)
		return c.JSON(http.StatusUnauthorized, r.ResponseMessage{Message: err.Error()})
	}
	ctx := c.Request().Context()
	res, err := h.authUseCase.GetUser(ctx, claims.UserId)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, r.ResponseMessage{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, res)
}

func (h *AccoutnHandler) Login(c echo.Context) (err error) {
	var data r.LoginRequest
	err = c.Bind(&data)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, r.ResponseMessage{Message: err.Error()})
	}
	ctx := c.Request().Context()
	res, err := h.authUseCase.Login(ctx, &data)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, r.ResponseMessage{Message: err.Error()})
	}
	token, err := _jwt.GenerateJWT(res.Id, res.Uuid, *res.Email, *res.Username)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusExpectationFailed, r.ResponseMessage{Message: err.Error()})
	}
	response := r.AuthenticationResponse{
		User:  res,
		Token: token,
	}
	// go func() {
	// 	h.authUseCase.UpdateFcmToken(context.Background(),data.FcmToken,res.ProfileId)
	// }()
	return c.JSON(http.StatusOK, response)
}
