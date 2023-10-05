package usecase

import (
	"context"
	r "regate-core/domain/repository"
	"time"

)

type infoUseCase struct {
	timeout time.Duration
	infoU r.InfoRepository
	utilU r.UtilUseCase
}

func NewUseCase(timeout time.Duration,infoU r.InfoRepository,utilU r.UtilUseCase) r.InfoUseCase{
	return &infoUseCase{
		timeout: timeout,
		infoU: infoU,
		utilU: utilU,
	}
}

func (u *infoUseCase) GetInfoText(ctx context.Context,id int)(res r.InfoText,err error){
	ctx,cancel := context.WithTimeout(ctx,u.timeout)
	defer cancel()
	res,err = u.infoU.GetInfoText(ctx,id)
	if err != nil {
		u.utilU.LogError("GetInfoText","info_usecase",err.Error())
	}
	return
}

