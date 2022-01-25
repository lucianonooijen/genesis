package user

import (
	"git.bytecode.nl/bytecode/genesis/server/internal/data/database"
	"git.bytecode.nl/bytecode/genesis/server/internal/interactors"
)

func jwtForUser(s *interactors.Services, user *database.User) (string, error) {
	uniqueUserIdentifyer := string(user.ID)

	return s.JWT.CreateJWT(uniqueUserIdentifyer)
}
