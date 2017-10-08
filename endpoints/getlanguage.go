package endpoints

import (
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/arraying/Arraybot-API/files"

	"github.com/gorilla/mux"
)

// GetLanguage handles the GET request for the endpoint.
func GetLanguage(writer http.ResponseWriter, request *http.Request) {
	writer = setContent(writer)
	language := mux.Vars(request)["language"]
	if !files.IsLanguage(language) {
		respond(writer, http.StatusBadRequest, RegularResponse{
			Success: false,
			Message: "Language does not exist",
		})
		return
	}
	raw := files.GetLanguageName()
	name := strings.Replace(raw, "{language}", language, -1)
	if !files.DoesFileExist(name, false) {
		respond(writer, http.StatusInternalServerError, RegularResponse{
			Success: false,
			Message: "Language file does not exist",
		})
		return
	}
	content, err := ioutil.ReadFile(name)
	if err != nil {
		respond(writer, http.StatusInternalServerError, RegularResponse{
			Success: false,
			Message: "Could not read language",
		})
		return
	}
	respond(writer, http.StatusOK, RegularResponse{
		Success: true,
		Message: files.FormatJSON(content),
	})
}
