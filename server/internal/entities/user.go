package entities

// NewUserRequest contains the data for creating a new user.
type NewUserRequest struct {
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required"`
	FirstName string `json:"firstName" validate:"required,alphanum"`
}

// NewUserResponse it the received response when account creation is successful.
type NewUserResponse struct {
	JWT string `json:"jwt"`
}
