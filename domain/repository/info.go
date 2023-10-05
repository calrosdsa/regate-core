package repository

import "context"

type InfoText struct {
	Id      int     `json:"id"`
	Title   *string `json:"title"`
	Content string  `json:"content"`
}

type InfoRepository interface {
	GetInfoText(ctx context.Context,id int)(res InfoText,err error)
}

type InfoUseCase interface {
	GetInfoText(ctx context.Context,id int)(res InfoText,err error)
}