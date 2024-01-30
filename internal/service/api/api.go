package api

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"palworld-chan/internal/service/api/routes"
	"palworld-chan/internal/service/cron"
)

func Init(port string) {

	cron.StartCron()

	app := fiber.New(fiber.Config{
		Prefork:       false,
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "Fiber",
		AppName:       "palworld-chan v0.0.1",
	})

	routes.Setup(app)

	log.Fatal(app.Listen(port))
}
