package main

import (
	"context"
	"embed"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"tapi/desktop/hub"
	"tapi/desktop/logger"
	"tapi/desktop/service"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	system := service.System()
	store := service.Store()
	collection := service.Collection()
	environment := service.Environment()
	group := hub.Group()
	hub := hub.Hub()
	project := service.Project()
	request := service.Request()
	session := service.Session()
	// Create application with options
	err := wails.Run(&options.App{
		Title:  "tapi",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup: func(ctx context.Context) {
			logger.SetupLogger(ctx)

			// initialize
			system.Init(ctx)
			store.Init(ctx)
			hub.Init(ctx)
			group.Init(ctx)
			project.Init(ctx)
			environment.Init(ctx)
			collection.Init(ctx)
			request.Init(ctx)
			session.Init(ctx)

		},
		OnShutdown: func(ctx context.Context) {
			session.Close()
			request.Close()
			collection.Close()
			environment.Close()
			project.Close()
			group.Close()
			hub.Close()
			store.Close()
			system.Close()
		},
		Bind: []interface{}{
			system,
			hub,
			group,
			session,
			project,
			environment,
			collection,
			request,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
