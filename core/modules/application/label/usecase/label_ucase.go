package usecase

import (
	"context"
	r "regate-core/domain/repository"
	"time"
)

type labelUseCase struct {
	timeout   time.Duration
	labelRepo r.LabelRepository
	utilU     r.UtilUseCase
}

func NewUseCase(timeout time.Duration, labelRepo r.LabelRepository, utilU r.UtilUseCase) r.LabelUseCase {
	return &labelUseCase{
		timeout:   timeout,
		labelRepo: labelRepo,
		utilU:     utilU,
	}
}

func (u *labelUseCase) CreateOrUpdateLabel(ctx context.Context, d r.LabelRequest) (err error) {
	ctx, cancel := context.WithTimeout(ctx, u.timeout)
	defer cancel()
	switch d.LabelType {
	case r.LabelDeporte:
		err = u.labelRepo.CreateDeporte(ctx, d.Label)
		if err != nil {
			u.utilU.LogError("CreateDeporte", "info_usecase", err.Error())
		}
	case r.LabelCategory:
		err = u.labelRepo.CreateCategory(ctx, d.Label)
		if err != nil {
			u.utilU.LogError("CreateCategory", "info_usecase", err.Error())
		}
	case r.LabelAmenity:
		err = u.labelRepo.CreateAmenity(ctx, d.Label)
		if err != nil {
			u.utilU.LogError("CreateAmenity", "info_usecase", err.Error())
		}
	case r.LabelRule:
		err = u.labelRepo.CreateRule(ctx, d.Label)
		if err != nil {
			u.utilU.LogError("CreateRule", "info_usecase", err.Error())
		}
	}
	
	return
}

func (u *labelUseCase)DeleteLabel(ctx context.Context,d r.LabelRequest)(err error) {
	ctx, cancel := context.WithTimeout(ctx, u.timeout)
	defer cancel()
	switch d.LabelType {
	case r.LabelAmenity:
		err = u.labelRepo.DeleteAmenity(ctx,d.Label.Id)
		if err != nil {
			u.utilU.LogError("DeleteAmenity", "info_usecase", err.Error())
		}
	case r.LabelRule:
		err = u.labelRepo.DeleteRule(ctx,d.Label.Id)
		if err != nil {
			u.utilU.LogError("DeleteRule", "info_usecase", err.Error())
		}
	}
	return
}
