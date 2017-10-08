package endpoints

import (
	"net/http"
)

// Endpoint is the endpoint object of an API endpoint.
type Endpoint struct {
	Name    string
	Method  string
	Pattern string
	Handler http.HandlerFunc
}

const (
	apiBase      = "/api/"
	languageBase = "languages/"
)

// Endpoints is a collection of all endpoints.
var Endpoints = []Endpoint{
	Endpoint{
		Name:    "Get Languages",
		Method:  "GET",
		Pattern: apiBase + languageBase,
		Handler: GetLanguages,
	},
	Endpoint{
		Name:    "Get Language",
		Method:  "GET",
		Pattern: apiBase + languageBase + "{language}",
		Handler: GetLanguage,
	},
	Endpoint{
		Name:    "Patch Language",
		Method:  "PATCH",
		Pattern: apiBase + languageBase + "{language}",
		Handler: PatchLanguage,
	},
}
