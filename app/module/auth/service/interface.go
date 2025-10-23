package auth

import "boilerplate/app/models/dto"

type IAuthService interface {
	SignIn(payload dto.LoginPaylod) error
	SignUp(payload dto.SignUpPaylod) error
}
