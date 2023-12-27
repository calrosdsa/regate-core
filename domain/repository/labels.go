package repository

import "context"

type LabelRepository interface {
	CreateCategory(ctx context.Context, d Label) (err error)
	CreateDeporte(ctx context.Context, d Label) (err error)
	CreateRule(ctx context.Context, d Label) (err error)
	CreateAmenity(ctx context.Context, d Label) (err error)

	DeleteAmenity(ctx context.Context,id int)(err error)
	DeleteRule(ctx context.Context,id int)(err error)
}

type LabelUseCase interface {
	CreateOrUpdateLabel(ctx context.Context, d LabelRequest)(err error)
	DeleteLabel(ctx context.Context,d LabelRequest)(err error)
}

type LabelRequest struct {
	Label     Label     `json:"label"`
	LabelType LabelType `json:"label_type"`
}

type Label struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Thumnail string `json:"thumbnail,omitempty"`
}

type LabelType byte

const (
	LabelCategory LabelType = 1
	LabelDeporte  LabelType = 2
	LabelAmenity  LabelType = 3
	LabelRule     LabelType = 4
)
