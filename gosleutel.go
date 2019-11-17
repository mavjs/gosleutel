package gosleutel

import (
	"fmt"
	"golang.org/x/crypto/scrypt"
	"strings"
)

const SALT = "$sleutel#%s#sleutel$"

func createSalt(acc string) []byte {
	salt := []byte(fmt.Sprintf(SALT, strings.ToLower(acc)))
	return salt
}

func CreatePass(password string, account string) ([]byte, error) {
	hash, err := scrypt.Key([]byte(password), createSalt(account), 16, 8, 16, 32)
	if err != nil {
		return hash, err
	}
	return hash, nil
}
