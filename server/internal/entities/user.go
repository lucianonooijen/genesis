package entities

import "github.com/google/uuid"

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

// PasswordResetStartRequest contains the required data to reset the password for a user account.
type PasswordResetStartRequest struct {
	Email string `json:"email" validate:"required,email"`
}

// PasswordResetCompleteRequest contains the data to complete the password reset for a user account.
type PasswordResetCompleteRequest struct {
	ResetToken uuid.UUID `json:"resetToken" validate:"required"`
	Password   string    `json:"password" validate:"required"`
}

// UserProfile contains the profile for a user.
type UserProfile struct {
	FirstName string `json:"firstName" validate:"required,alphanum"`
}

// DeleteAccountRequest is the data required to delete the user's account.
type DeleteAccountRequest struct {
	Password string `json:"password" validate:"required"`
}
