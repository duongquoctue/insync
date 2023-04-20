package server

import (
	"net/http"

	"github.com/duongquoctue/insync/cmd/api/session-api/handlers"
	"github.com/duongquoctue/insync/core/coreapp"
	"github.com/duongquoctue/insync/core/corehttp"
)

type server struct {}

func Init(app *coreapp.App) {
	// connMap := app.InitGRPCConnection([]string{
	// 	codename.DMSUser.CodeName,
	// })
	s := &server{}

	r := s.InitRoutes(app.HTTPBasePath)
	app.ServeAPI(r)
}

func (s *server) InitRoutes(path string) corehttp.Routes {
	return corehttp.Routes{
		{
			Name:     "Client subscription",
			Method:   http.MethodPost,
			BasePath: path,
			Pattern:  "/subscribe",
			Handler:  &handlers.SubcribeHandler{},
		},
		{
			Name:     "Emit to client",
			Method:   http.MethodGet,
			BasePath: path,
			Pattern:  "/send-message",
			Handler: &handlers.SendMessageHandler{},
		},
	}
}
