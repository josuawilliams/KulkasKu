package user

import (
	"context"
	"database/sql"
	"errors"
	"kulkasku/internal/dto"
	"kulkasku/internal/helper"
	"kulkasku/internal/model"
	"kulkasku/pkg/jwt"
	refreshtoken "kulkasku/pkg/refresh_token.go"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func (r *userService) Login(ctx context.Context, loginRequest *dto.LoginRequest) (*helper.WebResponse, error) {
	//Cek Email
	userExist, err := r.userRepository.GetUserByEmail(ctx, loginRequest.Email)
	if err != nil {
		return &helper.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "Error Get User (24)",
			Data:   err.Error(),
		}, nil
	}

	if userExist == nil {
		return &helper.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "User Not Found",
			Data:   nil,
		}, nil
	}

	// Cek Password
	err = bcrypt.CompareHashAndPassword([]byte(userExist.Password), []byte(loginRequest.Password))
	if err != nil {
		return &helper.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "Invalid Password",
			Data:   nil,
		}, nil
	}

	//Generate JWT
	token, err := jwt.GenerateToken(userExist.ID, userExist.Email, userExist.Name, r.cfg.SecretJwt)
	if err != nil {
		return &helper.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "Error At JWT Generate Token",
			Data:   err.Error(),
		}, nil
	}

	//Check Refresh Token Apakah Ada
	now := time.Now()
	refreshTokenExist, err := r.userRepository.GetRefreshToken(ctx, userExist.ID, now)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			refreshTokenExist = nil
		} else {
			return &helper.WebResponse{
				Code:   http.StatusInternalServerError,
				Status: "Error At Checking Refresh Token",
				Data:   err.Error(),
			}, nil
		}
	}

	if refreshTokenExist != nil {
		return &helper.WebResponse{
			Code:   http.StatusOK,
			Status: "Refresh Token Found",
			Data: map[string]interface{}{
				"token":         token,
				"refresh_token": refreshTokenExist.RefreshToken,
			},
		}, nil
	}

	refreshToken, err := refreshtoken.GenerateRefreshToken()
	if err != nil {
		return &helper.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "Error At Generate Refresh Token",
			Data:   err.Error(),
		}, nil
	}

	//Store Refresh Token
	err = r.userRepository.StoreRefreshToken(ctx, &model.RefreshToken{
		UserID:       userExist.ID,
		RefreshToken: refreshToken,
		ExpiresAt:    time.Now().Add(24 * time.Hour),
		CreatedAt:    now,
		UpdatedAt:    now,
	})

	if err != nil {
		return &helper.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "Error At Store Refresh Token",
			Data:   err.Error(),
		}, nil
	}

	return &helper.WebResponse{
		Code:   http.StatusOK,
		Status: "Login Success",
		Data: map[string]interface{}{
			"token":         token,
			"refresh_token": refreshToken,
		},
	}, nil

}
