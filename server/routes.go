package server

import "net/http"

type Route struct {
	Method string
	URI string
	HandlerFuc  http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"PUT",
		"/PUT",
		Put,
	},
	Route{
		"POST",
		"/POST",
		Post,
	},
	Route{
		"GET",
		"/GET",
		Get,
	},
	Route{
		"DELETE",
		"/DELETE",
		Delete,
	},
}