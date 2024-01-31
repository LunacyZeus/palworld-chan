package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"net/http"
	"palworld-chan/internal/service/api/controllers"
	"palworld-chan/pkg/utility/embeds"
)

func Setup(app *fiber.App) {
	// public routes // Unauthenticated route
	app.Use("/", filesystem.New(filesystem.Config{
		Root:       http.FS(embeds.EmbedDir),
		PathPrefix: "frontend/dist/index.html",
		Browse:     true,
	}))

	app.Use("/app.config.js", filesystem.New(filesystem.Config{
		Root:       http.FS(embeds.EmbedDir),
		PathPrefix: "frontend/dist/app.config.js",
		Browse:     true,
	}))

	app.Use("/logo.svg", filesystem.New(filesystem.Config{
		Root:       http.FS(embeds.EmbedDir),
		PathPrefix: "frontend/dist/logo.svg",
		Browse:     true,
	}))

	// Access file "image.png" under `static/` directory via URL: `http://<server>/static/image.png`.
	// Without `PathPrefix`, you have to access it via URL:
	// `http://<server>/static/static/image.png`.
	app.Use("/static", filesystem.New(filesystem.Config{
		Root:       http.FS(embeds.EmbedDir),
		PathPrefix: "frontend/dist/static",
		Browse:     true,
	}))

	// Login route
	app.Post("/api/login", controllers.Login)

	// JWT Middleware
	//app.Use(jwtware.New(jwtware.Config{
	//	SigningKey: jwtware.SigningKey{Key: []byte("secret")},
	//}))

	app.Get("/api/getUserInfo", controllers.GetUserInfo)

	// private routes, authenticatd by middleware
	app.Get("/api/GameServer/get", controllers.GetGameServerInfo)
	app.Get("/api/ServerInfo/get", controllers.GetServerInfo)

	app.Post("/api/BroadCast/send", controllers.SendBroadCast)
	app.Get("/api/Players/show", controllers.ShowPlayers)
	app.Get("/api/rcon/info", controllers.RconInfo)
	app.Get("/api/rcon/restart", controllers.RestartServer)
	app.Post("/api/serverSetting/update", controllers.UpdateServerSetting)
	app.Get("/api/serverSetting/get", controllers.GetServerSetting)
}
