package main

import (
	"fmt"
	"github.com/atotto/clipboard"
	"github.com/mavjs/gosleutel"
	"golang.org/x/crypto/ssh/terminal"
	"os"
	"time"
)

const duration = 20

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

	terminal.Restore(int(os.Stdin.Fd()), oldState)

	hash, err := gosleutel.CreatePass(password, account)
	if err != nil {
		panic(err)
	}

	clipboard.WriteAll(fmt.Sprintf("%X", hash))

	fmt.Printf("Password copied to clipboard. Will clear in %dsecs.\n", duration)

	time.Sleep(time.Second * 20)

	clipboard.WriteAll("")

	fmt.Println("Cleared password from keyboard. Exiting application.....")

	time.Sleep(time.Second * 5)

	os.Exit(0)
}
