package endpoints

import (
	"fmt"
	"net/http"
)

// Index shows the index page.
func Index(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(writer, "Welcome to the API.")
}
