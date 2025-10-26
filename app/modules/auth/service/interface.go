package auth

import "boilerplate/app/models/dto"

type IAuthService interface {
	SignIn(payload dto.LoginPaylod) (dto.SignInResponse, error)
	SignUp(payload dto.SignUpPaylod) (uint, error)
}
