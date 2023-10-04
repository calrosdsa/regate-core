package repository

import (
	"bytes"
)

type Ids struct {
	Id int `json:"id"`
}

type UtilUseCase interface {
	PaginationValues(page int16) int16
	GetNextPage(results int8, pageSize int8, page int16) (nextPage int16)
	LogError(method string, file string, err string)
	LogInfo(method string, file string, err string)
	CustomLog(method string, file string, err string, payload map[string]interface{})
	LogFatal(method string, file string, err string, payload map[string]interface{})
}

type MailUseCase interface {
	SendEmailVerification(mail string, otp int, username string)
	SendEmail(emails []string, body *bytes.Buffer, subject string) (err error)
	SendEmailUserAdmin(mail string, password string)
	SendResetPasswordEmail(email string, token string)
}

type MessageNotification struct {
	EnitityId int    `json:"id"`
	Message   string `json:"message"`
	Title     string `json:"title"`
}
