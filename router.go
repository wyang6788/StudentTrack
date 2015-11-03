package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
    router := mux.NewRouter().StrictSlash(true)
    for _, route := range routes {
        var handler http.Handler
        handler = route.HandlerFunc
        handler = Logger(handler, route.Name)

				if route.Name == "Get" {
					router.
            Methods(route.Method).
            Path(route.Pattern).
            Name(route.Name).
						Queries("name","{name}").
            Handler(handler)
				} else{
					router.
            Methods(route.Method).
            Path(route.Pattern).
            Name(route.Name).
            Handler(handler)
				}
    }
    return router
}

