package coreapp

import (
	goflag "flag"

	"github.com/duongquoctue/insync/core/logger"
)

func (app *App) loadCommon() {
	if !app.ConfigLoaded {
		app.loadFlags()
		app.ConfigLoaded = true
	}

	app.initLogger()
}

func (app *App) loadFlags() {
	goflag.StringVar(&app.HTTPHost, "host", "", "HTTP listen host")
	goflag.Int64Var(&app.HTTPPort, "port", 40000, "HTTP listen port")
	goflag.StringVar(&app.ConfigRemoteAddress, "cfg-addr", "", "Configuration remote address. ip:port")
	goflag.StringVar(&app.ConfigRemoteKeys, "cfg-keys", "", "Configuration remote keys. Seperate by ,")
	goflag.Parse()
}

func (app *App) initLogger() {
	logger.InitLogger()
}
