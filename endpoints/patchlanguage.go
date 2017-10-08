package endpoints

import (
	//"encoding/json"
	"github.com/arraying/Arraybot-API/files"
	"github.com/gorilla/mux"
	"io"
	"io/ioutil"
	"net/http"
)

// PatchLanguage handles the PATCH request for the endpoint.
func PatchLanguage(writer http.ResponseWriter, request *http.Request) {
	writer = setContent(writer)
	language := mux.Vars(request)["language"]
	if !files.IsLanguage(language) {
		respond(writer, http.StatusBadRequest, RegularResponse{
			Success: false,
			Message: "Language does not exist",
		})
		return
	}
	token := request.Header.Get("Authorization")
	if token == "" {
		respond(writer, http.StatusUnauthorized, RegularResponse{
			Success: false,
			Message: "You must provde the auth token",
		})
		return
	}
	if !files.IsValidToken(language, token) {
		respond(writer, http.StatusForbidden, RegularResponse{
			Success: false,
			Message: "Access denied",
		})
		return
	}
	body, err := ioutil.ReadAll(io.LimitReader(request.Body, 1048576))
	if err != nil {
		respond(writer, http.StatusInternalServerError, RegularResponse{
			Success: false,
			Message: "Error loading body",
		})
		return
	}
	if err = request.Body.Close(); err != nil {
		respond(writer, http.StatusInternalServerError, RegularResponse{
			Success: false,
			Message: "Error closing body",
		})
		return
	}
	update(writer, language, body)
}
