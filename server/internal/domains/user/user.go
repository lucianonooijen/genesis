package user

import (
	"context"

	"go.uber.org/zap"

	"git.bytecode.nl/bytecode/genesis/internal/data/database"
	"git.bytecode.nl/bytecode/genesis/internal/entities"
	"git.bytecode.nl/bytecode/genesis/internal/interactors"
)

// CreateUser creates a user in the database, sends an email and returns a JWT.
func CreateUser(s *interactors.Services, newUser entities.NewUserRequest) (*entities.JwtResponse, error) {
	log := s.BaseLogger.Named("domains/user/CreateUser")
	ctx := context.TODO()

	// Hash password
	log.Debug("hashing password")

	hashedPassword, err := s.PassHash.PlainTextToHash(newUser.Password)
	if err != nil {
		return nil, err
	}

	// Create user
	log.Debug("saving user to database")

	userToDB := database.CreateUserParams{
		Email:        newUser.Email,
		PasswordHash: hashedPassword,
		FirstName:    newUser.FirstName,
	}

	err = s.Database.CreateUser(ctx, userToDB)
	if err != nil {
		return nil, err
	}

	// Fetch new user
	log.Debug("fetching user from database", zap.String("email", newUser.Email))

	user, err := s.Database.GetUserByEmail(ctx, newUser.Email)
	if err != nil {
		return nil, err
	}

	// Send confirmation email
	log.Debug("sending confirmation email")

	if err = s.Mailer.SendAccountCreated(user.Email, user.FirstName); err != nil {
		return nil, err
	}

	// Generate JWT
	log.Debug("generating JWT")

	jwt, err := s.JWT.CreateJWT(string(user.ID))
	if err != nil {
		return nil, err
	}

	// Return
	log.Debug("create user done")

	return &entities.JwtResponse{JWT: jwt}, nil
}

// Login takes a login request and sends generated a JWT if the password matches the password hash in the database.
func Login(s *interactors.Services, loginData entities.LoginRequest) (*entities.JwtResponse, error) {
	log := s.BaseLogger.Named("domains/user/LoginUser")
	ctx := context.TODO()

	// Fetch user from DB
	user, err := s.Database.GetUserByEmail(ctx, loginData.Email)
	if err != nil {
		return nil, err
	}

	// Check password hash
	err = s.PassHash.ComparePassToHash(loginData.Password, user.PasswordHash)
	if err != nil {
		return nil, err
	}

	// Generate JWT
	log.Debug("generating JWT")

	jwt, err := s.JWT.CreateJWT(string(user.ID))
	if err != nil {
		return nil, err
	}

	// Return data
	log.Debug("create user done")

	return &entities.JwtResponse{JWT: jwt}, nil
}
