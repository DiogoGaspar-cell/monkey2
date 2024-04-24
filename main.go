package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/diogo-gaspar23/monkey2/repl"
)

func main() {
	currUser, err := user.Current()

	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! This is the Monkey programming language!\n", currUser.Username)
	fmt.Printf("Feel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}