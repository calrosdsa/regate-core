package repository

import (
	"context"
	"mime/multipart"
	"os"
	"sync"
	"time"
)
type CasoFile struct{
	Id int `json:"id"`
	FileUrl string `json:"file_url"`
	Extension string `json:"extension"`
	Descripcion string `json:"descripcion"`
	CasoId string `json:"caso_id"`
	CreatedOn time.Time `json:"created_on"`
}


type MediaUseCase interface{
	// UploadFileCaso(ctx context.Context,file *multipart.FileHeader,id string,descripcion string,ext string) (CasoFile,error)
	// GetFileCasos(ctx context.Context,id string)([]CasoFile,error)
	UploadImage(ctx context.Context,file *multipart.FileHeader,filename string,folder string)(url string,err error)
	UploadImage2(ctx context.Context,file *os.File,filename string,folder string)(url string,err error)
	UploadImageWithoutCtx(wg *sync.WaitGroup,file *multipart.FileHeader,urls *[]string)
	GenereteAddressImage(ctx context.Context,lng string,lat string,uuid string)(urlImg string,err error)
	CreateBucket(ctx context.Context,bucket string)(err error)
	UploadFile(ctx context.Context,file *multipart.FileHeader,path string,bucket string)(url string,err error)
}

type MediaRepository interface {
	// UploadFileCaso(ctx context.Context,url string,id string,descripcion string,ext string) (CasoFile,error)
	// GetFileCasos(ctx context.Context,id string)([]CasoFile,error)
}