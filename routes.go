package main

import "net/http"

type Route struct{
	Name	string
	Method	string
	Pattern  string
	HandlerFunc  http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"List",
		"GET",
		"/Student/listall",
		List,
	},
	Route{
		"Get",
		"GET",
		"/Student/getstudent",
		Get,
	},
	Route{
		"Post",
		"POST",
		"/Student",
		Post,
	},
	Route{
		"Update",
		"PUT",
		"/Student",
		Update,		
	},
	Route{
		"Delete",
		"DELETE",
		"/Student/{year:[0-9]+}",
		Delete,
	},
}
