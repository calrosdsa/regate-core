package ucases

import (
	"context"
	"io"
	"log"
	"mime/multipart"
	"os"
	"strings"

	// "path/filepath"
	aws_service "regate-core/core/modules/util/media/services"
	r "regate-core/domain/repository"

	// "strings"
	"sync"
	"time"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/nickalie/go-webpbin"

	"github.com/google/uuid"
	"github.com/spf13/viper"

    "net/http"

)

// const

type mediaUseCase struct {
	contextTimeout time.Duration
	sess           *session.Session
}

func NewUseCase(sess *session.Session) r.MediaUseCase {
	timeout := time.Duration(20) * time.Second
	return &mediaUseCase{
		// mediaRepo:      mu,
		contextTimeout: timeout,
		sess:           sess,
	}
}

func (mu *mediaUseCase) CreateBucket(ctx context.Context,bucket string)(err error){
	err = aws_service.CreateBucket(mu.sess,bucket)
	return
}

func (mu *mediaUseCase)UploadFile(ctx context.Context,file *multipart.FileHeader,path string,bucket string) (url string ,err error){
	url,err = aws_service.UplaodObject(ctx,file,path,bucket,mu.sess)
	return
}

func (mu *mediaUseCase) GenereteAddressImage(ctx context.Context,lng string,lat string,uuid string)(urlImg string,err error){
	var b strings.Builder
	b.WriteString("https://api.mapbox.com/styles/v1/mapbox/streets-v12/static/geojson(%7B%22type%22%3A%22Point%22%2C%22coordinates%22%3A%5B")
	b.WriteString(lng + "%2C" + lat + "%5D%7D)/")
	b.WriteString(lng+","+lat+",14/500x300?access_token=")
	b.WriteString(viper.GetString(`mapbox_api_key`))
	url := b.String()
	response, err := http.Get(url)
    if err != nil {
        log.Println(err)
    }
    defer response.Body.Close()
	 //open a file for writing
	file, err := os.Create("place.png")
	 if err != nil {
		 log.Fatal(err)
	 }
	defer func(){
		file.Close()
		err := os.Remove(file.Name())
		if err != nil {
			log.Println(err)
		}
	}()
	// Use io.Copy to just dump the response body to the file. This supports huge files
	_, err = io.Copy(file, response.Body)
	if err != nil {
		log.Fatal(err)
	}
	filename := uuid + ".webp"
	path := uuid + "/address/" 
	urlImg,err = mu.UploadImage2(ctx,file,filename,path)
	return
}

func (mu *mediaUseCase) UploadImage2(ctx context.Context, file *os.File,filename string,folder string) (url string, err error) {
	// filename := strings.TrimSuffix(file.Filename, filepath.Ext(file.Filename)) + ".wepb"
	
	
	err = webpbin.NewCWebP().
		Quality(80).
		InputFile(file.Name()).
		OutputFile(filename).
		Run()
	if err != nil {
		log.Println(err)
	}
	fileWebp, err := os.Open(filename)
	if err != nil {
		log.Println(err)
	}
	url, err = aws_service.UplaodObjectWebp(ctx,fileWebp, "teclu-soporte", mu.sess,folder)
	
	defer func() {
		if err := file.Close(); err != nil {
			log.Println(err)
		}
		err := os.Remove(file.Name())
		if err != nil {
			log.Println(err)
		}
		if err := fileWebp.Close(); err != nil {
			log.Println(err)
		}
		err1 := os.Remove(filename)
		if err1 != nil {
			log.Println(err1)
		}
	}()
	return
}


func (mu *mediaUseCase) UploadImageWithoutCtx(wg *sync.WaitGroup,file *multipart.FileHeader,urls *[]string) {
	filename := uuid.New().String() + ".webp"
	src, err := file.Open()
	if err != nil {
		return
	}
	// Destination
	dst, err := os.Create(file.Filename)
	if err != nil {
		return
	}
	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		log.Println(err)
	}

	err = webpbin.NewCWebP().
		Quality(40).
		InputFile(dst.Name()).
		OutputFile(filename).
		Run()
	if err != nil {
		log.Println(err)
	}
	fileWebp, err := os.Open(filename)
	if err != nil {
		log.Println(err)
	}
	url, err := aws_service.UplaodObjectWebpWithoutCxt(fileWebp, "teclu-soporte", mu.sess)
	if err != nil {
		log.Println(err)
	}
	*urls = append(*urls, url)
	
	defer func() {
		src.Close()
		if err := dst.Close(); err != nil {
			log.Println(err)
		}
		err := os.Remove(dst.Name())
		if err != nil {
			log.Println(err)
		}
		if err := fileWebp.Close(); err != nil {
			log.Println(err)
		}
		err1 := os.Remove(filename)
		if err1 != nil {
			log.Println(err1)
		}
		wg.Done()
	}()
}

func (mu *mediaUseCase) UploadImage(ctx context.Context, file *multipart.FileHeader,filename string,folder string) (url string, err error) {
	// filename := strings.TrimSuffix(file.Filename, filepath.Ext(file.Filename)) + ".wepb"
	
	src, err := file.Open()
	if err != nil {
		return
	}
	// Destination
	dst, err := os.Create(file.Filename)
	if err != nil {
		return
	}
	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		log.Println(err)
	}

	
	err = webpbin.NewCWebP().
		Quality(80).
		InputFile(dst.Name()).
		OutputFile(filename).
		Run()
	if err != nil {
		log.Println(err)
	}
	fileWebp, err := os.Open(filename)
	if err != nil {
		log.Println(err)
	}
	url, err = aws_service.UplaodObjectWebp(ctx,fileWebp, "teclu-soporte", mu.sess,folder)
	
	defer func() {
		src.Close()
		if err := dst.Close(); err != nil {
			log.Println(err)
		}
		err := os.Remove(dst.Name())
		if err != nil {
			log.Println(err)
		}
		if err := fileWebp.Close(); err != nil {
			log.Println(err)
		}
		err1 := os.Remove(filename)
		if err1 != nil {
			log.Println(err1)
		}
	}()
	return
}

