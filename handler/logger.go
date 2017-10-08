package handler

import (
	"log"
	"net/http"
	"time"
)

func Logger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		start := time.Now()
		inner.ServeHTTP(writer, request)
		log.Printf(
			"%s\t%s\t%s\t%s",
			request.Method,
			request.RequestURI,
			name,
			time.Since(start),
		)
	})
}
