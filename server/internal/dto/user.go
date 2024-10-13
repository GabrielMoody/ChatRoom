package dto

type (
	UserRegistration struct {
		Name            string `json:"name"`
		Email           string `json:"email"`
		Password        string `json:"password" validate:"min=8"`
		PasswordConfirm string `json:"password_confirm" validate:"required,eqfield=Password"`
	}

	UserLogin struct {
		Email    string `json:"email" form:"email" validate:"required,email"`
		Password string `json:"password" form:"password" validate:"required"`
	}
)
