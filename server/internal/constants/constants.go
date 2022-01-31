package constants

const (
	// BasePathAPI is the base used for the base of all requests.
	// Example: [server url]:[port]/[BasePathAPI]/[all other urls]
	// CAN ONLY BE CHANGED WITH A MAJOR RELEASE!
	BasePathAPI = "/v1"
	// APIStaticPath is the extension after BasePathAPI where static files are served.
	APIStaticPath = "/static"
)
