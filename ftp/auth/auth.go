package auth

import (
	"bytes"
	"crypto/sha256"
	"errors"
)

type hash [sha256.Size]byte

const salt = "progfay/ftp-server"

var usernamePasswordHashMap = map[string]hash{
	"progfay": toHash("password"),
}

func toHash(str string) hash {
	return sha256.Sum256([]byte(salt + str))
}

var (
	UserNotExists     = errors.New("user not exists")
	IncorrectPassword = errors.New("incorrect password")
)

func Verify(username, password string) error {
	passwordHash, exists := usernamePasswordHashMap[username]
	if !exists {
		return UserNotExists
	}

	h := toHash(password)
	if !bytes.Equal(passwordHash[:], h[:]) {
		return IncorrectPassword
	}

	return nil
}
