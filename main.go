package main

import (
	"fmt"
	"os"
	"os/user"

	"funlang/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is 'BasicLang' Programming language by Sifatul",
		user.Username)
	fmt.Println("Feel free to type in commands")
	repl.Start(os.Stdin, os.Stdout)
}
