package responses

// ServerVersion gets overwritten with build script to git tag or commit hash
var ServerVersion = "development" // TODO: Do this differently

// ResponseBody is the body base sent for every response
type ResponseBody struct {
	Success bool        `json:"success"`
	Error   *string     `json:"error"`
	Version string      `json:"version"`
	Data    interface{} `json:"data"`
}

// TODO: Add the auth status (and if logged in, user UUID) to the response
func generateResponseBody(success bool, error *string, data interface{}) ResponseBody {
	return ResponseBody{
		Success: success,
		Error:   error,
		Version: ServerVersion,
		Data:    data,
	}
}
