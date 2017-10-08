package endpoints

import (
	"encoding/json"
	"net/http"

	"github.com/arraying/Arraybot-API/files"
)

// GetLanguages handles the GET request for the endpoint.
func GetLanguages(writer http.ResponseWriter, request *http.Request) {
	writer = setContent(writer)
	writer.WriteHeader(http.StatusOK)
	languages := LanguagesResponse{
		Languages: files.GetLanguages(),
	}
	if err := json.NewEncoder(writer).Encode(languages); err != nil {
		panic(err)
	}
}
