package endpoints

import (
	"net/http"
)

func setContent(writer http.ResponseWriter) http.ResponseWriter {
	writer.Header().Set("Content-Type", "application/json;charset=UTF-8")
	return writer
}
