package repository

import (
	"context"
	"time"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthenticationResponse struct {
	Token string `json:"access_token"`
	User  User   `json:"user"`
}

type User struct {
	Id          int        `json:"user_id"`
	Uuid        string     `json:"uuid"`
	Username    *string    `json:"username,omitempty"`
	PhoneNumber *string     `json:"phone_number,omitempty"`
	LastLogin   *time.Time `json:"last_Login,omitempty"`
	CreatedAt   *time.Time `json:"created_at,omitempty"`
	Email       *string    `json:"email"`
	Otp         *int       `json:"otp,omitempty"`
	Estado      int        `json:"estado"`
	Rol      int        `json:"rol"`
	Password    string     `json:"password,omitempty"`
}

type AuthUseCase interface {
	Login(ctx context.Context, d *LoginRequest) (User, error)
	VerifyEmail(ctx context.Context, userId int, otp int) (res User, err error)
	GetUser(ctx context.Context, userId int) (User, error)
}

// ArticleRepository represent the article's repository contract
type AuthRepository interface {
	Login(ctx context.Context, d *LoginRequest) (res User, err error)
	VerifyEmail(ctx context.Context, userId int, otp int) (err error)
	GetUser(ctx context.Context, userId int) (User, error)
}
