package repository

import "context"

type EmpresaUseCase interface {
	GetEmpresasByEstado(ctx context.Context,estado EmpresaEstado)(res []Empresa,err error)
	CreateEmpresa(ctx context.Context,d *Empresa)(err error)
	GetUuidEmpresa(ctx context.Context,id int)(uuid string,err error)
	GetEmpresa(ctx context.Context,uuid string)(res Empresa,err error)
}

type EmpresaRepository interface {
	GetEmpresasByEstado(ctx context.Context,estado EmpresaEstado)(res []Empresa,err error)
	CreateEmpresa(ctx context.Context,d *Empresa)(err error)
	GetUuidEmpresa(ctx context.Context,id int)(uuid string,err error)
	GetEmpresa(ctx context.Context,uuid string)(res Empresa,err error)
}

type Empresa struct {
	Id           int      `json:"id"`
	Uuid         string   `json:"uuid"`
	Name         string   `json:"name"`
	CreatedAt    string   `json:"created_at"`
	UpdatedAt    string   `json:"updated_at,omitempty"`
	AdminId      *string   `json:"admin_id,omitempty"`
	PhoneNumber  *string `json:"phone_number,omitempty"`
	Estado EmpresaEstado `json:"estado"`
	Email        *string   `json:"email,omitempty"`
	Address      *string   `json:"address,omitempty"`
	Latidud      *string     `json:"latitud,omitempty"`
	Longitud     *string     `json:"longitud,omitempty"`
}


type EmpresaEstado int8
const (
	EmpresaEstadoAll EmpresaEstado = 0
	EmpresaEstadoActive EmpresaEstado = 1
	EmpresaEstadoInactive EmpresaEstado = 2
)

type EmpresaSetting struct {
	Id           int     `json:"id,omitempty"`
	EmpresaId    int     `json:"empresa_id,omitempty"`
	CurrencyId   int     `json:"currency_id,omitempty"`
	Tarifa       float32 `json:"tarifa,omitempty"`
	CurrencyAbb  string  `json:"currency_abb,omitempty"`
	CurrencyName string  `json:"currency_name,omitempty"`
}
