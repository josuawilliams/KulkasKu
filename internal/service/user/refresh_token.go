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

func (s *userService) RefreshToken(ctx context.Context, refreshTokenRequest *dto.RefreshTokenRequest, userId int64) (*helper.WebResponse, error) {
	//Check User Exist
	userExist, err := s.userRepository.GetUserById(ctx, userId)
	if err != nil {
		return &helper.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "Internal Server Error At Refresh Token",
			Data:   err.Error(),
		}, nil
	}

	if userExist == nil {
		return &helper.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "Invalid or expired token (30)",
			Data:   nil,
		}, nil
	}

	// Get refresh token by User Id (only valid / not expired)
	refreshTokenExists, err := s.userRepository.GetRefreshToken(ctx, userId, time.Now())
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			refreshTokenExists = nil
		} else {
			return &helper.WebResponse{
				Code:   http.StatusInternalServerError,
				Status: "Internal Server Error (43)",
				Data:   err.Error(),
			}, nil
		}
	}

	if refreshTokenExists == nil {
		return &helper.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "Refresh Token Expired",
			Data:   nil,
		}, nil
	}

	//check refresh token is match with request body
	if refreshTokenExists.RefreshToken != refreshTokenRequest.RefreshToken {
		return &helper.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "Invalid Refresh Token",
			Data:   nil,
		}, nil
	}

	//Generate Token
	token, err := jwt.GenerateToken(userExist.ID, userExist.Email, userExist.Name, s.cfg.SecretJwt)
	if err != nil {
		return &helper.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "Error At Generate Token (71)",
			Data:   err.Error(),
		}, nil
	}

	//delete old refresh token & generate new refresh token
	err = s.userRepository.DeleteRefreshToken(ctx, userId)
	if err != nil {
		return &helper.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "Error At Delete Old Refresh Token",
			Data:   err.Error(),
		}, nil
	}

	refreshtoken, err := refreshtoken.GenerateRefreshToken()
	if err != nil {
		return &helper.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "Error At Generate New Refresh Token (90)",
			Data:   err.Error(),
		}, nil
	}

	//store New Refresh Token
	err = s.userRepository.StoreRefreshToken(ctx, &model.RefreshToken{
		UserID:       userId,
		RefreshToken: refreshtoken,
		ExpiresAt:    time.Now().Add(24 * time.Hour),
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	})
	if err != nil {
		return &helper.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "Error At Create New Refresh Token (105)",
			Data:   err.Error(),
		}, nil
	}

	return &helper.WebResponse{
		Code:   http.StatusOK,
		Status: "New Refresh Token Success",
		Data: map[string]interface{}{
			"token":        token,
			"refreshToken": refreshtoken,
		},
	}, nil
}
