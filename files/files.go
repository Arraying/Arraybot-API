package files

import (
	"encoding/json"
	"io/ioutil"
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

// FormatJSON formats all JSON and removes excessive whitespace.
func FormatJSON(input string) string {
	return regex.ReplaceAllString(input, "")
}
