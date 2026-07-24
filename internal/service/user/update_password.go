package user

import (
	"context"
	"kulkasku/internal/dto"
	"kulkasku/internal/helper"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func (r *userService) UpdatePassword(ctx context.Context, req *dto.UpdatePasswordRequest, userId int64) (*helper.WebResponse, error) {
	userExist, err := r.userRepository.GetUserById(ctx, userId)
	if err != nil {
		return &helper.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "Internal Server Error",
			Data:   nil,
		}, nil
	}

	if userExist == nil {
		return &helper.WebResponse{
			Code:   http.StatusNotFound,
			Status: "User Not Found",
			Data:   nil,
		}, nil
	}

	if userExist.Password != "" {
		err = bcrypt.CompareHashAndPassword([]byte(userExist.Password), []byte(req.CurrentPassword))
		if err != nil {
			return &helper.WebResponse{
				Code:   http.StatusUnauthorized,
				Status: "Current Password Salah",
				Data:   nil,
			}, nil
		}
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(req.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		return &helper.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "Gagal hash password",
			Data:   nil,
		}, nil
	}

	err = r.userRepository.UpdateUserPassword(ctx, userId, string(passwordHash))
	if err != nil {
		return &helper.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "Gagal update password",
			Data:   nil,
		}, nil
	}

	return &helper.WebResponse{
		Code:   http.StatusOK,
		Status: "Password Berhasil Diubah",
		Data:   nil,
	}, nil
}
