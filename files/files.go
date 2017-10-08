package files

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"regexp"
)

var regex = regexp.MustCompile("\\t|\\r|\\n")

// DoesFileExist checks if the specified file exists.
func DoesFileExist(path string, dir bool) bool {
	stat, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	isDir := stat.IsDir()
	if dir && !isDir {
		return false
	}
	return true
}

// LoadConfig loads the config file.
func LoadConfig() error {
	out, err := ioutil.ReadFile(ConfigName)
	if err != nil {
		return err
	}
	err = json.Unmarshal(out, &Configuration)
	return err
}

// GetLanguages gets a collection of languages.
func GetLanguages() (languages []string) {
	languages = make([]string, 0)
	for _, language := range Configuration.Languages {
		languages = append(languages, language.Language)
	}
	return
}

// GetLanguageName gets a generic language path.
func GetLanguageName() string {
	return Configuration.LanguagesPath + "/{language}.json"
}

// IsLanguage checks whether a provided string is a valid language.
func IsLanguage(language string) bool {
	for _, entry := range GetLanguages() {
		if language == entry {
			return true
		}
	}
	return false
}

// IsValidToken checks if the provided token is valid.
func IsValidToken(language, token string) bool {
	for _, entry := range Configuration.Languages {
		if entry.Language == language {
			for _, languageToken := range entry.Tokens {
				if token == languageToken {
					return true
				}
			}
		}
	}
	return false
}

// FormatJSON formats all JSON and removes excessive whitespace.
func FormatJSON(input []byte) string {
	buffer := new(bytes.Buffer)
	err := json.Compact(buffer, input)
	if err != nil {
		log.Fatalln("Error compacting JSON,", err)
	}
	return buffer.String()
}
