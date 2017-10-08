package files

const (
	// ConfigName is the configuration file name.
	ConfigName = "config.json"
)

var (
	// Configuration is the config object.
	Configuration Config
)

// Config is the configuration.
type Config struct {
	Port          string `json:"port"`
	LanguagesPath string `json:"languages_path"`
	Languages     []struct {
		Language string   `json:"language"`
		Tokens   []string `json:"tokens"`
	} `json:"languages"`
}
