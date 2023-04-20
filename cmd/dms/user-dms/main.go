package main

import (
	"github.com/duongquoctue/insync/cmd/dms/user-dms/server"
	"github.com/duongquoctue/insync/core/coreapp"
	"github.com/duongquoctue/insync/pkg/codename"
)

func main() {
	app := coreapp.NewDMSApp(
		codename.DMSUser.CodeName,
	)

	server.Init(app)
}
