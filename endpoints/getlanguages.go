package endpoints

import (
	"encoding/json"
	"net/http"

	"github.com/arraying/Arraybot-API/files"
	"github.com/arraying/Arraybot-API/ratelimits"
)

// GetLanguages handles the GET request for the endpoint.
func GetLanguages(writer http.ResponseWriter, request *http.Request) {
	writer = setContent(writer)
	if ratelimit, exists := ratelimits.Ratelimits[files.APIGetLanguages]; exists {
		if ratelimited := ratelimit.Handle(writer, request); ratelimited {
			return
		}
	}
	writer.WriteHeader(http.StatusOK)
	languages := LanguagesResponse{
		Languages: files.GetLanguages(),
	}
	if err := json.NewEncoder(writer).Encode(languages); err != nil {
		panic(err)
	}
}
