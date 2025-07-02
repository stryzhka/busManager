package main

import (
	"busManager/routers"
	"context"
	"embed"
	"fmt"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	//// Create an instance of the app structure
	app, err := NewApp()
	if err != nil {
		fmt.Println(err)
	}
	busRouter, err := routers.NewBusRouter()
	if err != nil {
		fmt.Println(err)
	}
	busStopRouter, err := routers.NewBusStopRouter()
	if err != nil {
		fmt.Println(err)
	}
	driverRouter, err := routers.NewDriverRouter()
	if err != nil {
		fmt.Println(err)
	}
	routeRouter, err := routers.NewRouteRouter()
	if err != nil {
		fmt.Println(err)
	}
	// Create application with options
	err = wails.Run(&options.App{
		Title:  "busManager",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup: func(ctx context.Context) {
			app.startup(ctx)
			busRouter.Startup(ctx)
			busStopRouter.Startup(ctx)
			driverRouter.Startup(ctx)
			routeRouter.Startup(ctx)
		},
		Bind: []interface{}{
			app,
			busRouter,
			busStopRouter,
			driverRouter,
			routeRouter,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}

}
