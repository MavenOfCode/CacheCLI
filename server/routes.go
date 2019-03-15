package server

import "net/http"

type Route struct {
	Method     string
	URI        string
	HandlerFuc http.HandlerFunc
}

type Routes []Route
