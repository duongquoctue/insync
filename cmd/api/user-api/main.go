package main

import (
	"github.com/duongquoctue/insync/cmd/api/user-api/server"
	"github.com/duongquoctue/insync/core/coreapp"
	"github.com/duongquoctue/insync/pkg/codename"
)

func main() {
	app := coreapp.NewAPIApp(
		codename.APIUser.CodeName,
		codename.APIUser.HTTPBasePath,
	)

	server.Init(app)
}
