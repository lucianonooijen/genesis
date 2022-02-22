package responses

// ErrorBody is the JSON body that will be returned for non-2XX responses.
type ErrorBody struct {
	// Status contains the HTTP status code
	Status int `json:"status,omitempty" example:"418"`

	// Title is the human-readable error
	Title string `json:"title,omitempty" example:"Error title"`

	// Detail contains details about the error
	Detail string `json:"detail,omitempty" example:"Slightly longer error detail"`

	// RawError contains the raw error data
	RawError interface{} `json:"rawError,omitempty"`
}
