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
)

func (r *userService) LoginGoogle(ctx context.Context, loginGoogleRequest *dto.LoginGoogleRequest) (*helper.WebResponse, error) {
	userExist, err := r.userRepository.GetUserByEmail(ctx, loginGoogleRequest.Email)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return &helper.WebResponse{
				Code:   http.StatusInternalServerError,
				Status: "Internal Server Error At Login Google",
				Data:   nil,
			}, nil
		}
	}

	var userID int64
	var name string
	hasPassword := false

	if userExist == nil {
		now := time.Now()
		userModel := &model.User{
			Name:      loginGoogleRequest.Name,
			Email:     loginGoogleRequest.Email,
			GoogleID:  &loginGoogleRequest.GoogleID,
			AvatarURL: &loginGoogleRequest.AvatarURL,
			CreatedAt: now,
			UpdatedAt: now,
		}

		userID, err = r.userRepository.CreateUser(ctx, userModel)
		if err != nil {
			return &helper.WebResponse{
				Code:   http.StatusInternalServerError,
				Status: "Gagal membuat user google",
				Data:   nil,
			}, nil
		}
		name = loginGoogleRequest.Name
	} else {
		if userExist.GoogleID == nil || *userExist.GoogleID != loginGoogleRequest.GoogleID {
			return &helper.WebResponse{
				Code:   http.StatusUnauthorized,
				Status: "Invalid Google Account",
				Data:   nil,
			}, nil
		}

		userID = userExist.ID
		name = userExist.Name
		hasPassword = userExist.Password != ""

		if loginGoogleRequest.AvatarURL != "" && (userExist.AvatarURL == nil || *userExist.AvatarURL == "") {
			_ = r.userRepository.UpdateUserGoogle(ctx, userID, loginGoogleRequest.GoogleID, loginGoogleRequest.AvatarURL)
		}
	}

	token, err := jwt.GenerateToken(userID, loginGoogleRequest.Email, name, r.cfg.SecretJwt)
	if err != nil {
		return &helper.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "Error At JWT Generate Token",
			Data:   nil,
		}, nil
	}

	now := time.Now()
	refreshTokenExist, err := r.userRepository.GetRefreshToken(ctx, userID, now)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return &helper.WebResponse{
				Code:   http.StatusInternalServerError,
				Status: "Error At Checking Refresh Token",
				Data:   nil,
			}, nil
		}
	}

	if refreshTokenExist != nil {
		return &helper.WebResponse{
			Code:   http.StatusOK,
			Status: "Success Login Google",
			Data: map[string]interface{}{
				"token":         token,
				"refresh_token": refreshTokenExist.RefreshToken,
				"has_password":  hasPassword,
			},
		}, nil
	}

	refreshToken, err := refreshtoken.GenerateRefreshToken()
	if err != nil {
		return &helper.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "Error At Generate Refresh Token",
			Data:   nil,
		}, nil
	}

	err = r.userRepository.StoreRefreshToken(ctx, &model.RefreshToken{
		UserID:       userID,
		RefreshToken: refreshToken,
		ExpiresAt:    time.Now().Add(24 * time.Hour),
		CreatedAt:    now,
		UpdatedAt:    now,
	})
	if err != nil {
		return &helper.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "Error At Store Refresh Token",
			Data:   nil,
		}, nil
	}

	return &helper.WebResponse{
		Code:   http.StatusOK,
		Status: "Success Login Google",
		Data: map[string]interface{}{
			"token":         token,
			"refresh_token": refreshToken,
			"has_password":  hasPassword,
		},
	}, nil

}
