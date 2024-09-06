package main

import (
	"log"

	"ahyalfan.my.id/chat_rom_management/internal/api"
	"ahyalfan.my.id/chat_rom_management/internal/component"
	"ahyalfan.my.id/chat_rom_management/internal/config"
	"ahyalfan.my.id/chat_rom_management/internal/middleware"
	"ahyalfan.my.id/chat_rom_management/internal/repository"
	"ahyalfan.my.id/chat_rom_management/internal/service"
	"ahyalfan.my.id/chat_rom_management/internal/ws"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	cnf := config.NewConfig()
	app := fiber.New(
		fiber.Config{
			ErrorHandler: func(c *fiber.Ctx, err error) error {
				log.Printf("Error: %v\n", err)
				return nil
			},
		},
	)

	// middleware
	app.Use(recover.New())
	corsConfig := cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
		AllowMethods: "GET, POST, PUT, DELETE",
		MaxAge:       86400, // 24 hours
	}
	app.Use(cors.New(corsConfig))

	jwtMid := middleware.GetJWT(cnf)

	db := component.GetConnection(cnf)

	userRepository := repository.NewUserRepository(db)

	userService := service.NewUser(cnf, userRepository)

	// api
	api.NewUser(app, userService)

	// ws
	hub := ws.NewHub()
	ws.NewHandler(app, hub, jwtMid, cnf)

	// jalanain channel
	go hub.Run()

	err := app.Listen(cnf.Server.Host + ":" + cnf.Server.Port)
	if err != nil {
		log.Fatal(err)
	}
}
