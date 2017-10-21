package endpoints

import (
	"github.com/arraying/Arraybot-API/files"
	"net/http"
)

// Endpoint is the endpoint object of an API endpoint.
type Endpoint struct {
	Name    string
	Method  string
	Pattern string
	Handler http.HandlerFunc
}

// Endpoints is a collection of all endpoints.
var Endpoints = []Endpoint{
	Endpoint{
		Name:    files.APIIndex,
		Method:  "GET",
		Pattern: files.APIBase,
		Handler: Index,
	},
	Endpoint{
		Name:    files.APIGetLanguages,
		Method:  "GET",
		Pattern: files.APIBase + files.LanguageBase,
		Handler: GetLanguages,
	},
	Endpoint{
		Name:    files.APIGetLanguage,
		Method:  "GET",
		Pattern: files.APIBase + files.LanguageBase + "{language}",
		Handler: GetLanguage,
	},
	Endpoint{
		Name:    files.APIPatchLanguage,
		Method:  "PATCH",
		Pattern: files.APIBase + files.LanguageBase + "{language}",
		Handler: PatchLanguage,
	},
}
