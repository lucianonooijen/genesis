package responses

// ErrorBody is the JSON body that will be returned for non-2XX responses
type ErrorBody struct {
	// Title is the human-readable error
	Title string `json:"title,omitempty"`

	// Detail contains details about the error
	Detail string `json:"detail,omitempty"`

	// Status contains the HTTP status code
	Status int `json:"status,omitempty"`

	// RawError contains the raw error data
	RawError interface{} `json:"rawError,omitempty"`
}
