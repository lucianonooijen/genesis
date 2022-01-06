package entities

// NewUserRequest contains the data for creating a new user.
type NewUserRequest struct {
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required"`
	FirstName string `json:"firstName" validate:"required,alphanum"`
}

// JwtResponse it the received response when account creation is successful.
type JwtResponse struct {
	JWT string `json:"jwt"`
}

// LoginRequest contains the data for login requests.
type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}
