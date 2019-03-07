package server

import (
	"github.com/gorilla/mux"
	"net/http"
)

func NewRouter() *mux.Router  {
	router := mux.NewRouter().StrictSlash(true)
	for _,route := range routes{
		var handler http.Handler
		
		handler = route.HandlerFuc
		//put logger here later when I get to that
		
		router.
			Methods(route.Method).
			Path(route.URI).
			Handler(handler)
	}
	return router
}