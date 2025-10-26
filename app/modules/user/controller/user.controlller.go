package user

import (
	"github.com/gofiber/fiber/v2"

	"boilerplate/app/common"
	"boilerplate/app/models/dto"
	us "boilerplate/app/modules/user/service"
)

type UserController struct {
	userService us.IUserService
	authGuard   common.IGuard
}

var _ common.IController = &UserController{}

func New(authGuard common.IGuard, userService us.IUserService) *UserController {
	return &UserController{
		userService: userService,
		authGuard:   authGuard,
	}
}

func (controller *UserController) GetCurrent(ctx *fiber.Ctx) error {
	userId := ctx.Locals("user").(dto.User).ID

	user, err := controller.userService.FindById(userId)

	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(user)
}

func (controller *UserController) Register(app *fiber.App) {
	group := app.Group("/users")

	group.Get(
		"/current",
		controller.authGuard.Activate,
		controller.GetCurrent,
	)

	common.Logger.Info("User controller registered successfully")
}
