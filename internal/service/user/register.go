package user

import (
	"context"
	"database/sql"
	"errors"
	"kulkasku/internal/dto"
	"kulkasku/internal/helper"
	"kulkasku/internal/model"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func (r *userService) Register(ctx context.Context, registerRequest *dto.RegisterRequest) (*helper.WebResponse, error) {
	//Cek User Udah Ada Atau Belum
	userExist, err := r.userRepository.GetUserByEmail(ctx, registerRequest.Email)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return &helper.WebResponse{
				Code:   http.StatusInternalServerError,
				Status: "Internal Server Error At Register User (23)",
				Data:   err.Error(),
			}, nil
		}
	}

	if userExist != nil {
		return &helper.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "User Already Exist",
			Data:   nil,
		}, nil
	}

	//hash password

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(registerRequest.Password), bcrypt.DefaultCost)
	if err != nil {
		return &helper.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "Internal Server Error At Register User (Hash Password)",
			Data:   nil,
		}, nil
	}

	//Create User
	now := time.Now()
	userModel := &model.User{
		Name:      registerRequest.Name,
		Email:     registerRequest.Email,
		Password:  string(passwordHash),
		GoogleID:  nil,
		AvatarURL: nil,
		CreatedAt: now,
		UpdatedAt: now,
	}

	userId, err := r.userRepository.CreateUser(ctx, userModel)
	if err != nil {
		return &helper.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "Internal Server Error At Register User (Create User 65)",
			Data:   err.Error(),
		}, nil
	}

	return &helper.WebResponse{
		Code:   http.StatusOK,
		Status: "Berhasil Register User",
		Data:   userId,
	}, nil
}

func (r *userService) RegisterGoogle(ctx context.Context, registerGoogleRequest *dto.RegisterRequest) (*helper.WebResponse, error) {
	return nil, nil
}
