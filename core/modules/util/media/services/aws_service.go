package services

import (
	"context"
	"fmt"
	"mime/multipart"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"

)

func ListObjects(bucket string, sess *session.Session) {
    svc := s3.New(sess)
	resp, err := svc.ListObjectsV2(&s3.ListObjectsV2Input{Bucket: aws.String(bucket)})
	if err != nil {
		exitErrorf("Unable to list items in bucket %q, %v", bucket, err)
	}

	for _, item := range resp.Contents {
		fmt.Println("Name:         ", *item.Key)
		fmt.Println("Last modified:", *item.LastModified)
		fmt.Println("Size:         ", *item.Size)
		fmt.Println("Storage class:", *item.StorageClass)
		fmt.Println("")
	}

	fmt.Println("Found", len(resp.Contents), "items in bucket", bucket)
	fmt.Println("")
}


func UplaodObject(ctx context.Context,file *multipart.FileHeader,path string,bucket string, sess *session.Session)(url string,err error) {
	src, err := file.Open()
	if err != nil {
		return
	}
	defer src.Close()
    uploader := s3manager.NewUploader(sess)
    output, err := uploader.UploadWithContext(ctx,&s3manager.UploadInput{
        Bucket: aws.String(bucket),
        Key: aws.String(path+file.Filename),
        Body: src,
    })
    if err != nil {
        // Print the error and exit.
		return
    }

    fmt.Printf("Successfully uploaded %q to %q\n", file.Filename, bucket)
	return output.Location,nil
}


func UplaodObjectWebp(ctx context.Context,file *os.File,bucket string, sess *session.Session,folder string)(url string,err error) {
	
    uploader := s3manager.NewUploader(sess)
    output, err := uploader.UploadWithContext(ctx,&s3manager.UploadInput{
        Bucket: aws.String(bucket),
        Key: aws.String(folder+file.Name()),
        Body: file,
    })
    if err != nil {
        // Print the error and exit.
		return
    }
    fmt.Printf("Successfully uploaded %q to %q\n", file.Name(), bucket)
	return output.Location,nil
}


func UplaodObjectWebpWithoutCxt(file *os.File,bucket string, sess *session.Session)(url string,err error) {
	
    uploader := s3manager.NewUploader(sess)
    output, err := uploader.Upload(&s3manager.UploadInput{
        Bucket: aws.String(bucket),
        Key: aws.String("test/"+file.Name()),
        Body: file,
    })
    if err != nil {
        // Print the error and exit.
		return
    }

    fmt.Printf("Successfully uploaded %q to %q\n", file.Name(), bucket)
	return output.Location,nil
}

func CreateBucket(sess *session.Session,bucket string)(err error){
	// Create S3 service client
	svc := s3.New(sess)
	_, err = svc.CreateBucket(&s3.CreateBucketInput{
		Bucket: aws.String(bucket),
	})
	if err != nil {
		exitErrorf("Unable to create bucket %q, %v", bucket, err)
		return
	}
	
	// Wait until bucket is created before finishing
	fmt.Printf("Waiting for bucket %q to be created...\n", bucket)
	
	err = svc.WaitUntilBucketExists(&s3.HeadBucketInput{
		Bucket: aws.String(bucket),
	})
	if err != nil {
		exitErrorf("Error occurred while waiting for bucket to be created, %v", bucket)
		return
	}
	
	fmt.Printf("Bucket %q successfully created\n", bucket)
	return
}



func exitErrorf(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, msg+"\n", args...)
	os.Exit(1)
}
