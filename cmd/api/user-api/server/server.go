package server

import (
	"net/http"

	"github.com/duongquoctue/insync/cmd/api/user-api/handlers"
	"github.com/duongquoctue/insync/core/coreapp"
	"github.com/duongquoctue/insync/core/corehttp"
	"github.com/duongquoctue/insync/pkg/codename"
	insync_service "github.com/duongquoctue/insync/pkg/proto/user/service"
)

type server struct {
	UserClient insync_service.UserClient
}

func Init(app *coreapp.App) {
	connMap := app.InitGRPCConnection([]string{
		codename.DMSUser.CodeName,
	})
	s := &server{
		UserClient: insync_service.NewUserClient(connMap[codename.DMSUser.CodeName]),
	}

	r := s.InitRoutes(app.HTTPBasePath)
	app.ServeAPI(r)
}

func (s *server) InitRoutes(path string) corehttp.Routes {
	return corehttp.Routes{
		{
			Name:     "Create new user",
			Method:   http.MethodPost,
			BasePath: path,
			Pattern:  "",
			Handler:  &handlers.CreateUserHandler{},
		},
		{
			Name:     "Get new user",
			Method:   http.MethodGet,
			BasePath: path,
			Pattern:  "",
			Handler: &handlers.GetUserHandler{
				UserClient: s.UserClient,
			},
		},
	}
}
