package repository

import (
	"context"
	"mime/multipart"
)

type BillingUseCase interface {
	GetDepositosEmpresa(ctx context.Context, d DepositoFilterData, page int16, size int8) (res []DepositoBancarioEmpresa, count int, nextPage int16, err error)
	UploadComprobanteDeposito(ctx context.Context, file *multipart.FileHeader, d DepositoBancario) (url string,err error)
	CreateDepositos(ctx context.Context)
}

type BillingRepository interface {
	GetDepositosEmpresa(ctx context.Context, d DepositoFilterData, page int16, size int8) (res []DepositoBancarioEmpresa, count int, err error)
	CreateDeposito(ctx context.Context, empresaId int) (err error)
	UploadComprobanteDeposito(ctx context.Context, d DepositoBancario) (err error)
	GetEmpresaIdByDepositoDetailId(ctx context.Context, id int) (res int, err error)
}
type DepositoFilterData struct {
	EmpresaId int            `json:"empresa_id"`
	Estado    EstadoDeposito `json:"estado"`
}

type DepositoBancarioEmpresa struct {
	Id          int     `json:"id"`
	Uuid        string  `json:"uuid"`
	EmpresaId   int     `json:"empresa_id"`
	CreatedAt   string  `json:"created_at"`
	TotalIncome float64 `json:"total_income"`
}

type DepositoBancario struct {
	Id                  int            `json:"id"`
	Uuid                string         `json:"uuid,omitempty"`
	Gloss               string         `json:"gloss,omitempty"`
	CreatedAt           string         `json:"created_at,omitempty"`
	DatePaid            string         `json:"date_paid,omitempty"`
	Income              float64        `json:"income,omitempty"`
	Estado              EstadoDeposito `json:"estado,omitempty"`
	Tarifa              float32        `json:"tarifa,omitempty"`
	CurrencyAbb         string         `json:"currency_abb,omitempty"`
	EstablecimientoId   int            `json:"establecimiento_id,omitempty"`
	ParentId            int            `json:"parent_id,omitempty"`
	ComprobanteUrl      string        `json:"comprobante_url,omitempty"`
	EmitionDate         *string        `json:"emition_date"`
	EstablecimientoName string         `json:"establecimiento_name,omitempty"`
}

type IdsDepositoReturn struct {
	Id                int
	EstablecimientoId int
}

type EstadoDeposito int8

const (
	DEPOSITO_PENDIENTE = 1
	DEPOSITO_EMITIDO   = 2
)
