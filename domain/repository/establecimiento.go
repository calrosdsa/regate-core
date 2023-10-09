package repository

import (
	"context"
	"time"
)

type Establecimiento struct {
	Id           int                   `json:"id"`
	Uuid         string                `json:"uuid,omitempty"`
	Name         string                `json:"name"`
	Description  string                `json:"description,omitempty"`
	Address      string                `json:"address,omitempty"`
	Portada      *string               `json:"portada,omitempty"`
	Photo        *string               `json:"photo"`
	AddressPhoto *string               `json:"address_photo,omitempty"`
	Estado       EstablecimientoEstado `json:"estado,omitempty"`
	Latidud      string                `json:"latitud,omitempty"`
	Longitud     string                `json:"longitud,omitempty"`
	PhoneNumber  string                `json:"phone_number,omitempty"`
	IsOpen       bool                  `json:"is_open,omitempty"`
	Visibility   bool                  `json:"visibility"`
	Email        string                `json:"email,omitempty"`
	EmpresaId    int                   `json:"empresa_id,omitempty"`
	CreatedAt    *time.Time            `json:"created_at,omitempty"`
	UpdatedAt    *time.Time            `json:"updated_at,omitempty"`
	Amenities    []int64               `json:"amenities,omitempty"`
	Rules        []int64               `json:"rules,omitempty"`
	Distance     float64               `json:"distance,omitempty"`
}

type EstablecimientoRepository interface {
	GetEstablecimientosEmpresa(ctx context.Context,empresaUuid string)(res []Establecimiento,err error)
	GetEstablecimientosByEstado(ctx context.Context,estado EstablecimientoEstado)(res []int,err error)
	UpdateEstablecimientoTsv(ctx context.Context,id int)(err error)
}

type EstablecimientoUseCase interface {
	GetEstablecimientosEmpresa(ctx context.Context,empresaUuid string)(res []Establecimiento,err error)
	GetEstablecimientosByEstado(ctx context.Context,estado EstablecimientoEstado)(res []int,err error)
	UpdateEstablecimientosTsv(ctx context.Context)(err error)
}

type EstablecimientoEstado int8

const (
	EstablecimientoVerificado EstablecimientoEstado = 1
	EstablecimientoPendiente  EstablecimientoEstado = 2
)