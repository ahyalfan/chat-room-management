package middleware

import (
	"github.com/gofiber/fiber/v2"
)

func AuthAdmin() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		ctx.Get("au")
		return ctx.Next()
	}
}
