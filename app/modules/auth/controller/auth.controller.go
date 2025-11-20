package controller

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"

	"boilerplate/app/common"
	"boilerplate/app/middleware"
	dto "boilerplate/app/models/dto"
	as "boilerplate/app/modules/auth/service"
)

type AuthController struct {
	authService as.IAuthService
	log         *zap.Logger
}

var _ common.IController = &AuthController{}

func New(authService as.IAuthService, log *zap.Logger) *AuthController {
	return &AuthController{
		authService: authService,
		log:         log,
	}
}

func (controller *AuthController) SignIn(ctx *fiber.Ctx) error {
	payload := dto.LoginPaylod{}
	ctx.BodyParser(&payload)

	response, err := controller.authService.SignIn(payload)

	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (controller *AuthController) SignUp(ctx *fiber.Ctx) error {
	payload := dto.SignUpPaylod{}
	ctx.BodyParser(&payload)

	userId, err := controller.authService.SignUp(payload)

	if err != nil {
		return err
	}

	return ctx.
		Status(fiber.StatusCreated).
		JSON(
			dto.SignUpResponse{
				UserId:  userId,
				Message: "User created successfully",
			},
		)
}

func (controller *AuthController) Register(app *fiber.App) {

	group := app.Group("/auth")

	group.Post(
		"/login",
		middleware.BodyValidator[dto.LoginPaylod](),
		controller.SignIn,
	)
	group.Post(
		"/signup",
		middleware.BodyValidator[dto.SignUpPaylod](),
		controller.SignUp,
	)

	controller.log.Info("Auth controller registered successfully")
}
