package coreapp

import (
	"fmt"

	"github.com/duongquoctue/insync/core/coregrpc"
	"github.com/duongquoctue/insync/core/corehttp"
	"github.com/duongquoctue/insync/core/logger"
	"google.golang.org/grpc"
)

type CoreAppType string

const (
	TypeAPI      CoreAppType = "api"
	TypeDMS      CoreAppType = "dms"
	TypeConsumer CoreAppType = "consumer"
)

type App struct {
	Name string
	Type CoreAppType

	ConfigLoaded        bool
	ConfigRemoteAddress string
	ConfigRemoteKeys    string

	HTTPHost     string
	HTTPPort     int64
	HTTPBasePath string
}

func NewAPIApp(name string, httpBasePath string) *App {
	app := &App{
		Name:         name,
		HTTPBasePath: httpBasePath,
		Type:         TypeAPI,
	}

	app.loadCommon()
	return app
}

func NewDMSApp(name string) *App {
	app := &App{
		Name: name,
		Type: TypeDMS,
	}

	app.loadCommon()
	return app
}

func (app *App) InitGRPCConnection(codes []string) map[string]*grpc.ClientConn {
	if !app.ConfigLoaded {
		app.loadFlags()
		app.ConfigLoaded = true
	}

	addr := fmt.Sprintf("%v:%v", app.HTTPHost, app.HTTPPort)

	return coregrpc.Dial(addr, codes)
}

func (app *App) ServeAPI(routes corehttp.Routes) {
	if !app.ConfigLoaded {
		app.loadFlags()
		app.ConfigLoaded = true
	}

	addr := fmt.Sprintf("%v:%v", app.HTTPHost, app.HTTPPort)

	log := logger.NewLogger()
	log.Infof("serving HTTP server %v at %v", app.Name, addr)

	err := corehttp.Serve(addr, routes)
	if err != nil {
		log.Fatalf("can't start http server, err: %v", err)
	}
}

func (app *App) ServeDMS() {
	if !app.ConfigLoaded {
		app.loadFlags()
		app.ConfigLoaded = true
	}
}
