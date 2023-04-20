package server

import (
	"fmt"

	"github.com/duongquoctue/insync/core/coreapp"
	"github.com/duongquoctue/insync/core/drivers/mysql"
)

func Init(app *coreapp.App) {
	err := mysql.NewConnection(nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Established MySQL Connection")

	// Serve it when ready
	app.ServeDMS()
}
