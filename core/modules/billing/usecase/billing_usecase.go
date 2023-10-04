package usecase

import (
	"context"
	"fmt"
	"log"
	"mime/multipart"
	r "regate-core/domain/repository"
	"time"
)

type billingUseCase struct {
	billingRepo r.BillingRepository
	empresaU r.EmpresaUseCase
	timeout time.Duration
	mediaU r.MediaUseCase
	utilU r.UtilUseCase
}

func NewUseCase(timeout time.Duration,billingRepo r.BillingRepository,
	empresaU r.EmpresaUseCase,utilU r.UtilUseCase,mediaU r.MediaUseCase) r.BillingUseCase{
	return &billingUseCase{
		billingRepo: billingRepo,
		timeout: timeout,
		empresaU: empresaU,
		utilU: utilU,
		mediaU: mediaU,
	}
}

func (u *billingUseCase)GetDepositosEmpresa(ctx context.Context,d r.DepositoFilterData,page int16,size int8)(res []r.DepositoBancarioEmpresa,
	count int,nextPage int16,err error){
	ctx,cancel := context.WithTimeout(ctx,u.timeout)
	defer cancel()
	page = u.utilU.PaginationValues(page)
	res,count, err = u.billingRepo.GetDepositosEmpresa(ctx,d,page,size)
	if err != nil {
		u.utilU.LogError("GetDepositosEmpresa","billing_usecase.go",err.Error())
		return 
	}
	nextPage = u.utilU.GetNextPage(int8(len(res)),size,page +1)
	return 
}

func(u *billingUseCase)CreateDepositos(ctx context.Context){
	log.Println("CREATING DEPOSITO")
	res,err := u.empresaU.GetEmpresasByEstado(context.Background(),r.EmpresaEstadoActive)
	if err != nil {
		u.utilU.LogError("CreateDeposito","billing_usecase",err.Error())
	}
	for _,empresa := range res{
		err  = u.billingRepo.CreateDeposito(ctx,empresa.Id)
		if err != nil {
			u.utilU.LogError("CreateDeposito","billing_usecase",err.Error())
		}	
	}
}

func(u *billingUseCase)UploadComprobanteDeposito(ctx context.Context,file *multipart.FileHeader,d r.DepositoBancario)(url string,err error){
	ctx,cancel := context.WithTimeout(ctx,u.timeout)
	defer cancel()
	log.Println(d.ParentId)
	empresaId,err := u.billingRepo.GetEmpresaIdByDepositoDetailId(ctx,d.ParentId)
	if err != nil {
		u.utilU.LogError("UploadCombrobanteDeposito_GetEmpresaIdByDepositoDetail","billing_usecase",err.Error())
		return
	}
	uuid,err := u.empresaU.GetUuidEmpresa(ctx,empresaId)
	if err != nil {
		u.utilU.LogError("UploadCombrobanteDeposito_GetUuidEmpresa","billing_usecase",err.Error())
		return
	}
	log.Println(uuid)
	path :=  fmt.Sprintf("/depositos/%s/",d.DatePaid)
	log.Println(path)
	url,err = u.mediaU.UploadFile(ctx,file,path,uuid)
	if err != nil {
		u.utilU.LogError("UploadCombrobanteDeposito_UplaodFile","billing_usecase",err.Error())
		return
	}
	d.ComprobanteUrl = url
	err = u.billingRepo.UploadComprobanteDeposito(ctx,d)
	if err != nil {
		u.utilU.LogError("UploadCombrobanteDeposito_UplaodFile","billing_usecase",err.Error())
		return
	}
	return
}