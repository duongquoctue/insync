package main

import (
	"github.com/duongquoctue/insync/cmd/api/session-api/server"
	"github.com/duongquoctue/insync/core/coreapp"
	"github.com/duongquoctue/insync/pkg/codename"
)

func main() {
	app := coreapp.NewAPIApp(
		codename.APISession.CodeName,
		codename.APISession.HTTPBasePath,
	)

	server.Init(app)
}
