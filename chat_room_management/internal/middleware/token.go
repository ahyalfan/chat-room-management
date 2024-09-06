package middleware

import (
	"ahyalfan.my.id/chat_rom_management/dto"
	"ahyalfan.my.id/chat_rom_management/internal/config"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

func GetJWT(cnf *config.Config) func(*fiber.Ctx) error {
	return jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(cnf.JWT.Key)},
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(fiber.StatusUnauthorized).JSON(dto.CreateResponseError(fiber.StatusUnauthorized, err.Error()))
		},
	})
}
