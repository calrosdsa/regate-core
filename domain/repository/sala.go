package repository

import "context"

type SalaUseCase interface {
	DeleteUnAvailablesSalas(ctx context.Context)
	DisabledExpiredRooms(ctx context.Context)
}

type SalaRepository interface {
	DeleteUnAvailablesSalas(ctx context.Context)
	DisabledExpiredRooms(ctx context.Context)(res []int,err error)
}

type SalaEstado int8

const (
	SalaAvailable SalaEstado = iota
	SalaUnAvailable 
	SalaReserved
)