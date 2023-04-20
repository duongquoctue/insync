package codename

type API struct {
	CodeName     string
	HTTPBasePath string
}

var (
	APIUser    = API{CodeName: "user-api", HTTPBasePath: "/user"}
	APISession = API{CodeName: "session-api", HTTPBasePath: "/session"}
)
