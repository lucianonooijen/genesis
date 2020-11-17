package interactors

// JwtUtil is the interface returned by data/jwt
type JwtUtil interface {
	CreateJWT(user string) (token string, err error)
	ValidateJWT(token string) (user string, err error)
}
