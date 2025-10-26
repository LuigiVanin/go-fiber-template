package dto

type LoginPaylod struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=72"`
}

type SignUpPaylod struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=72"`
}

type SignUpResponse struct {
	UserId  uint   `json:"user_id"`
	Message string `json:"message"`
}

type SignInResponse struct {
	Token string `json:"token"`
	User  User   `json:"user"`
}
