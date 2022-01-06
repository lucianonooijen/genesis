package user

import (
	"git.bytecode.nl/bytecode/genesis/internal/data/database"
	"git.bytecode.nl/bytecode/genesis/internal/interactors"
)

func jwtForUser(s *interactors.Services, user *database.User) (string, error) {
	uniqueUserIdentifyer := string(user.ID)

	return s.JWT.CreateJWT(uniqueUserIdentifyer)
}
