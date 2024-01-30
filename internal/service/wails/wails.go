package wails

import (
	"context"
	"embed"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/logger"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
	"palworld-chan/internal/service/wails/app"
)

func Init(frontend embed.FS) {
	// Create an instance of the structs which are bound to the wails context
	app := app.NewApp()

	// Create application with options
	err := wails.Run(&options.App{
		Title:            "Wails Example",
		Width:            1024,
		Height:           768,
		WindowStartState: options.Maximised,
		Frameless:        false,
		AssetServer: &assetserver.Options{
			Assets: frontend,
		},
		Windows: &windows.Options{
			DisableWindowIcon: true,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 128},
		OnStartup: func(ctx context.Context) {
			app.OnStartup(ctx)
		},
		OnDomReady: func(ctx context.Context) {
			app.OnDomReady(ctx)
		},
		OnBeforeClose: func(ctx context.Context) (prevent bool) {
			return (app.OnBeforeClose(ctx))
		},
		Bind: []interface{}{
			app,
		},
		LogLevel:           logger.WARNING,
		LogLevelProduction: logger.ERROR,
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
