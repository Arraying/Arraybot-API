package main

// Imports
import (
	"log"
	"net/http"
	"strings"

	"github.com/arraying/Arraybot-API/files"
	"github.com/arraying/Arraybot-API/handler"
)

// The main function.
func main() {
	log.Println("Starting up the Arraybot API...")
	if !files.DoesFileExist(files.ConfigName, false) {
		log.Fatalf("Configuration file \"%v\" does not exist.\n", files.ConfigName)
	}
	if err := files.LoadConfig(); err != nil {
		log.Fatalln("Error loading the config,", err)
	}
	if !files.DoesFileExist(files.Configuration.LanguagesPath, true) {
		log.Fatalf("Language directory \"%v\" (relative) does not exist.\n", files.Configuration.LanguagesPath)
	}
	log.Println("Validating the languages...")
	for _, language := range files.GetLanguages() {
		raw := files.GetLanguageName()
		name := strings.Replace(raw, "{language}", language, -1)
		if !files.DoesFileExist(name, false) {
			log.Fatalf("The language \"%v\" specified in \"%v\" does not exist in the directory \"%v\"", language, files.ConfigName, files.Configuration.LanguagesPath)
		}
	}
	log.Println("Configuration OK, loading server...")
	router := handler.NewRouter()
	log.Fatal(http.ListenAndServe(files.Configuration.Port, router))
}
