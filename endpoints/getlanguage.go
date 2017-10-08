package endpoints

import (
	"encoding/json"
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
	var response RegularResponse
	if !files.IsLanguage(language) {
		writer.WriteHeader(http.StatusBadRequest)
		response = RegularResponse{
			Success: false,
			Message: "Language does not exist",
		}
	} else {
		raw := files.GetLanguageName()
		name := strings.Replace(raw, "{language}", language, -1)
		if !files.DoesFileExist(name, false) {
			writer.WriteHeader(http.StatusInternalServerError)
			response = RegularResponse{
				Success: false,
				Message: "Language file does not exist",
			}
		} else {
			content, err := ioutil.ReadFile(name)
			if err != nil {
				writer.WriteHeader(http.StatusInternalServerError)
				response = RegularResponse{
					Success: false,
					Message: "Could not read language",
				}
			} else {
				response = RegularResponse{
					Success: true,
					Message: files.FormatJSON(string(content)),
				}
			}
		}
	}
	if err := json.NewEncoder(writer).Encode(response); err != nil {
		panic(err)
	}
}
