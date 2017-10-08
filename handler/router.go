package handler

import (
	"github.com/arraying/Arraybot-API/endpoints"
	"github.com/gorilla/mux"
	"net/http"
)

// NewRouter creates a new router.
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, endpoint := range endpoints.Endpoints {
		var handler http.Handler
		handler = endpoint.Handler
		handler = Logger(handler, endpoint.Name)
		router.Methods(endpoint.Method).
			Path(endpoint.Pattern).
			Name(endpoint.Name).
			Handler(handler)
	}
	return router
}
