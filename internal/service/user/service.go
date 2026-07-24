package user

import (
	"context"
	"kulkasku/internal/dto"
	"kulkasku/internal/helper"
	"kulkasku/internal/model"
	"kulkasku/internal/repository/user"
)

type UserService interface {
	Register(ctx context.Context, registerRequest *dto.RegisterRequest) (*helper.WebResponse, error)
	Login(ctx context.Context, loginRequest *dto.LoginRequest) (*helper.WebResponse, error)
	RefreshToken(ctx context.Context, refreshTokenRequest *dto.RefreshTokenRequest, userId int64) (*helper.WebResponse, error)
	LoginGoogle(ctx context.Context, loginGoogleRequest *dto.LoginGoogleRequest) (*helper.WebResponse, error)
	UpdatePassword(ctx context.Context, req *dto.UpdatePasswordRequest, userId int64) (*helper.WebResponse, error)
	Me(ctx context.Context, userId int64) (*helper.WebResponse, error)
}

type userService struct {
	cfg            *model.Config
	userRepository user.UserRepository
}

func NewService(cfg *model.Config, userRepository user.UserRepository) UserService {
	return &userService{
		cfg:            cfg,
		userRepository: userRepository,
	}
}
