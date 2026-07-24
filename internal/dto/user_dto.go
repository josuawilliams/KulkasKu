package dto

type (
	RegisterRequest struct {
		Name            string `json:"name" validate:"required"`
		Email           string `json:"email" validate:"required,email"`
		Password        string `json:"password" validate:"required"`
		PasswordConfirm string `json:"password_confirm" validate:"required,eqfield=Password"`
		GoogleID        string `json:"google_id"`
		AvatarURL       string `json:"avatar_url"`
	}

	LoginGoogleRequest struct {
		Email     string `json:"email" validate:"required,email"`
		GoogleID  string `json:"google_id" validate:"required"`
		Name      string `json:"name"`
		AvatarURL string `json:"avatar_url"`
	}
)

type (
	LoginRequest struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required"`
	}

	LoginResponse struct {
		Token        string `json:"token"`
		RefreshToken string `json:"refresh_token"`
	}
)

type (
	UpdatePasswordRequest struct {
		CurrentPassword    string `json:"current_password"`
		NewPassword        string `json:"new_password" validate:"required"`
		NewPasswordConfirm string `json:"new_password_confirm" validate:"required,eqfield=NewPassword"`
	}
)

type (
	RefreshTokenRequest struct {
		RefreshToken string `json:"refresh_token" validate:"required"`
	}

	RefreshTokenResponse struct {
		Token        string `json:"token"`
		RefreshToken string `json:"refresh_token"`
	}
)
