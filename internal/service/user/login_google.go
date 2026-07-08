package user

import (
	"context"
	"database/sql"
	"errors"
	"kulkasku/internal/dto"
	"kulkasku/internal/helper"
	"net/http"
)

func (r *userService) LoginGoogle(ctx context.Context, loginGoogleRequest *dto.LoginGoogleRequest) (*helper.WebResponse, error) {
	userExist, err := r.userRepository.GetUserByEmail(ctx, loginGoogleRequest.Email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &helper.WebResponse{
				Code:   http.StatusNotFound,
				Status: "User Not Found",
				Data:   nil,
			}, nil
		}

		return &helper.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "Internal Server Error At Login Google",
			Data:   err.Error(),
		}, nil
	}

	if userExist.GoogleID == nil || *userExist.GoogleID != loginGoogleRequest.GoogleID {
		return &helper.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "Invalid Google Account",
			Data:   nil,
		}, nil
	}

	return &helper.WebResponse{
		Code:   http.StatusOK,
		Status: "Berhasil Login Google",
		Data:   userExist.ID,
	}, nil
}
