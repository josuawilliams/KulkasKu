package user

import (
	"context"
	"kulkasku/internal/helper"
	"net/http"
)

func (r *userService) Me(ctx context.Context, userId int64) (*helper.WebResponse, error) {
	user, err := r.userRepository.GetUserById(ctx, userId)
	if err != nil {
		return &helper.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "Internal Server Error",
			Data:   nil,
		}, nil
	}

	if user == nil {
		return &helper.WebResponse{
			Code:   http.StatusNotFound,
			Status: "User Not Found",
			Data:   nil,
		}, nil
	}

	return &helper.WebResponse{
		Code:   http.StatusOK,
		Status: "Success",
		Data: map[string]interface{}{
			"name":  user.Name,
			"email": user.Email,
		},
	}, nil
}
