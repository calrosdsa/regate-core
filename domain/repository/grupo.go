package repository

import "context"

type GrupoUseCase interface {
	UpdateTsvColumn(ctx context.Context)(err error)
}

type GrupoRepository interface {
	UpdateEstablecimientosTsv(ctx context.Context)(err error)
}