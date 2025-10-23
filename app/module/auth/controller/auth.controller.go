package auth

import (
	"boilerplate/app/common"
	"boilerplate/app/middleware"
	dto "boilerplate/app/models/dto"
	auth "boilerplate/app/module/auth/service"

	"github.com/gofiber/fiber/v2"
)

var _ common.IController = &AuthController{}

type AuthController struct {
	authService auth.IAuthService
}

func New(authService auth.IAuthService) *AuthController {
	return &AuthController{
		authService: authService,
	}
}

func (controller *AuthController) SignIn(ctx *fiber.Ctx) error {
	payload := dto.LoginPaylod{}
	ctx.BodyParser(&payload)

	err := controller.authService.SignIn(payload)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to login",
			"error":   err.Error(),
		})
	}

	return ctx.SendString("Hello, Login Page!")
}

func (controller *AuthController) SignUp(ctx *fiber.Ctx) error {
	payload := dto.SignUpPaylod{}
	ctx.BodyParser(&payload)

	err := controller.authService.SignUp(payload)

	if err != nil {
		return err
	}

	return ctx.SendString("Hello, Sign Up Page!")
}

func (controller *AuthController) Register(app *fiber.App) {

	app.Post(
		"/login",
		middleware.BodyValidator[dto.LoginPaylod](),
		controller.SignIn,
	)
	app.Post(
		"/signup",
		middleware.BodyValidator[dto.SignUpPaylod](),
		controller.SignUp,
	)
}
