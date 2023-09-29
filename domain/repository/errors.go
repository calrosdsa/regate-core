package repository

import (
	"errors"
	"net/http"
)

var (
	// ErrInternalServerError will throw if any the Internal Server Error happen
	ErrInternalServerError = errors.New("internal Server Error")
	// ErrNotFound will throw if the requested item is not exists
	ErrNotFound = errors.New("your requested Item is not found")
	// ErrConflict will throw if the current action already exists
	ErrConflict = errors.New("your Item already exist")
	// ErrBadParamInput will throw if the given request-body or params is not valid
	ErrBadParamInput = errors.New("given Param is not valid")

	PageError        = errors.New("page param is needed:(?page=1)")
	ErrNotFoundEmail = errors.New("no se encontró ninguna cuenta vinculada a este correo electrónico")
)
const (
	DbItemNotFound = "sql: no rows in result set"
)

type ResponseMessage struct {
	Message string `json:"message"`
}

func GetStatusCode(err error) int {
	switch err {
	case ErrNotFoundEmail:
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}