package constants

const (
	// GinContextKeyRole is the role key that is added to the request.
	GinContextKeyRole = "jwt_role"

	// GinContextKeyUser is the key used in Gin to access the user in the context.
	GinContextKeyUser = "user_logged_in"

	// GinHeaderNameClientVersion is the HTTP header name for the client version.
	GinHeaderNameClientVersion = "X-Genesis-Client-Version"

	// GinHeaderNameServerVersion is the HTTP header name for the server version.
	GinHeaderNameServerVersion = "X-Genesis-Server-Version"
)
