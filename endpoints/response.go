package endpoints

// LanguagesResponse is the response object when getting all languages.
type LanguagesResponse struct {
	Languages []string `json:"languages"`
}

// RegularResponse is a generic response object.
type RegularResponse struct {
	Success bool   `json:"s"`
	Message string `json:"m"`
}
