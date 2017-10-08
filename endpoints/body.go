package endpoints

// LanguagesResponse is the response object when getting all languages.
type LanguagesResponse struct {
	Languages []string `json:"l"`
}

// RegularResponse is a generic response object.
type RegularResponse struct {
	Success bool   `json:"s"`
	Message string `json:"m"`
}

// PatchPayload is the Payload when updating messages.
type PatchPayload struct {
	Key     string `json:"k"`
	Message string `json:"v"`
}
