package endpoints

import (
	"encoding/json"
	"github.com/arraying/Arraybot-API/files"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

const (
	languageKey   = "k"
	languageValue = "v"
)

func setContent(writer http.ResponseWriter) http.ResponseWriter {
	writer.Header().Set("Content-Type", "application/json;charset=UTF-8")
	return writer
}

func respond(writer http.ResponseWriter, status int, body interface{}) {
	writer.WriteHeader(status)
	err := json.NewEncoder(writer).Encode(body)
	if err != nil && status != http.StatusNoContent {
		log.Println("Could not encode body,", err)
	}
}

func update(writer http.ResponseWriter, language string, body []byte) {
	kv := make(map[string]string)
	if err := json.Unmarshal(body, &kv); err != nil {
		log.Println("Could not decode body,", err)
		respond(writer, http.StatusUnprocessableEntity, RegularResponse{
			Success: false,
			Message: "Malformed JSON",
		})
		return
	}
	key, hasKey := kv[languageKey]
	value, hasValue := kv[languageValue]
	if !(hasKey && hasValue) {
		respond(writer, http.StatusBadRequest, RegularResponse{
			Success: false,
			Message: "Must contain JSON keys " + languageKey + ", " + languageValue + ".",
		})
		return
	}
	name := strings.Replace(files.GetLanguageName(), "{language}", language, -1)
	content, err := ioutil.ReadFile(name)
	if err != nil {
		respond(writer, http.StatusInternalServerError, RegularResponse{
			Success: false,
			Message: "Could not read language",
		})
		return
	}
	languageJSON := make(map[string]string)
	json.Unmarshal(content, &languageJSON)
	languageJSON[key] = value
	newContent, err := json.Marshal(languageJSON)
	if err != nil {
		respond(writer, http.StatusInternalServerError, RegularResponse{
			Success: false,
			Message: "Could not marshal new language",
		})
		return
	}
	err = ioutil.WriteFile(name, newContent, 0600)
	respond(writer, http.StatusNoContent, nil)
}
