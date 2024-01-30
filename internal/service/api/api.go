package api

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"palworld-chan/internal/service/api/pkg/custom"
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
		// Global custom error handler
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(fiber.StatusBadRequest).JSON(custom.GlobalErrorHandlerResp{
				Code:    500,
				Type:    "error",
				Message: err.Error(),
			})
		},
	})

	routes.Setup(app)

	log.Fatal(app.Listen(port))
}
