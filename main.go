package main

import (
	"fmt"
	"golang.org/x/crypto/scrypt"
	"golang.org/x/crypto/ssh/terminal"
	"os"
	"strings"
)

const (
	SALT = "$sleutel#%s#sleutel$"
)

func createSalt(acc string) []byte {
	salt := []byte(fmt.Sprintf(SALT, strings.ToLower(acc)))
	return salt
}

func createPass(password string, account string) ([]byte, error) {
	hash, err := scrypt.Key([]byte(password), createSalt(account), 16, 8, 16, 32)
	if err != nil {
		return hash, err
	}
	return hash, nil
}

func main() {
	oldState, err := terminal.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		panic(err)
	}

	term := terminal.NewTerminal(os.Stdin, "")
	password, err := term.ReadPassword("master password >> ")
	if err != nil {
		panic(err)
	}
	term.SetPrompt("account >> ")
	account, err := term.ReadLine()
	if err != nil {
		panic(err)
	}

	hash, err := createPass(password, account)
	if err != nil {
		panic(err)
	}

	terminal.Restore(int(os.Stdin.Fd()), oldState)

	fmt.Printf("%X\n", hash)

}
