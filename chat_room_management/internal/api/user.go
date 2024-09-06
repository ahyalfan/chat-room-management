package api

import (
	"context"
	"time"

	"ahyalfan.my.id/chat_rom_management/domain"
	"ahyalfan.my.id/chat_rom_management/dto"
	"ahyalfan.my.id/chat_rom_management/internal/util"
	"github.com/gofiber/fiber/v2"
)

type userApi struct {
	userService domain.UserService
}

func NewUser(app *fiber.App, userService domain.UserService) {
	api := &userApi{userService: userService}
	app.Post("/api/v1/auth/signup", api.SignUp)
	app.Post("/api/v1/auth/login", api.Login)
}

func (api *userApi) SignUp(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	var req dto.UserCreatedReq
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.CreateResponseError(fiber.ErrBadRequest.Code, "Invalid :"+err.Error()))
	}

	fails := util.Validate(req)
	if len(fails) > 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.CreateResponseErrorData(fiber.StatusBadRequest, "creating request failed", fails))
	}

	res, err := api.userService.CreateUser(c, req)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(dto.CreateResponseError(fiber.StatusInternalServerError, err.Error()))
	}

	return ctx.JSON(dto.CreateResponseSuccess(fiber.StatusOK, res))
}

func (api *userApi) Login(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	var req dto.LoginUserReq
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.CreateResponseError(fiber.ErrBadRequest.Code, "Invalid :"+err.Error()))
	}

	fails := util.Validate(req)
	if len(fails) > 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.CreateResponseErrorData(fiber.StatusBadRequest, "creating request failed", fails))
	}

	res, err := api.userService.LoginUser(c, req)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(dto.CreateResponseError(fiber.StatusInternalServerError, err.Error()))
	}

	return ctx.JSON(dto.CreateResponseSuccess(fiber.StatusOK, res))
}
