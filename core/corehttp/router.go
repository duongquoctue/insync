package corehttp

import "net/http"
import "github.com/julienschmidt/httprouter"

type Routes []Route

type Route struct {
	Name string

	Method   string
	BasePath string
	Pattern  string

	Handler http.Handler
}

func Serve(addr string, routes Routes) error {
	router := httprouter.New()

	for _, route := range routes {
		path := route.BasePath + route.Pattern
		router.Handler(route.Method, path, route.Handler)
	}

	err := http.ListenAndServe(addr, router)
	if err != nil {
		return err
	}

	return nil
}
