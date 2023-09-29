package usecase

import (
	"context"
	"crypto/rand"
	"io"
	"log"
	r "regate-core/domain/repository"
	"strconv"
	"time"
)

type authUseCase struct {
	timeout time.Duration
	authRepo r.AuthRepository
}

func NewUseCase(timeout time.Duration,authRepo r.AuthRepository) r.AuthUseCase {
	return &authUseCase{
		timeout: timeout,
		authRepo: authRepo,
	}
}

func(u *authUseCase)VerifyEmail(ctx context.Context,userId int,otp int)(res r.User,err error){
	ctx,cancel := context.WithTimeout(ctx,u.timeout)
	defer cancel()
	err = u.authRepo.VerifyEmail(ctx,userId,otp)
	if err != nil{
		return
	}
	res,err = u.authRepo.GetUser(ctx,userId)
	return
}

func(u *authUseCase) GetUser(ctx context.Context,userId int)(res r.User,err error){
	ctx,cancel := context.WithTimeout(ctx,u.timeout)
	defer cancel()
	res,err = u.authRepo.GetUser(ctx,userId)
	return
}



func(u *authUseCase) Login(ctx context.Context,d *r.LoginRequest)(res r.User,err error){
	ctx,cancel := context.WithTimeout(ctx,u.timeout)
	defer cancel()
	res,err = u.authRepo.Login(ctx,d)
	return
}


func(u *authUseCase)generateOtp(max int) int {
	b := make([]byte, max)
    n, err := io.ReadAtLeast(rand.Reader, b, max)
    if n != max {
        log.Println(err)
    }
    for i := 0; i < len(b); i++ {
        b[i] = table[int(b[i])%len(table)]
    }
	otp,err := strconv.Atoi(string(b))
	if err != nil {
		log.Println(err)
		return 51432
	}
	return otp
}

var table = [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}