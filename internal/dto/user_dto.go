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
		Email    string `json:"email" validate:"required,email"`
		GoogleID string `json:"google_id" validate:"required"`
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
	RefreshTokenRequest struct {
		RefreshToken string `json:"refresh_token" validate:"required"`
	}

	RefreshTokenResponse struct {
		Token        string `json:"token"`
		RefreshToken string `json:"refresh_token"`
	}
)
