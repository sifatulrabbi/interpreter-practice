package main

import (
	"fmt"
	"log"
	"os"

	"funlang/inputparser"
	"funlang/tokenizer"
)

func main() {
	var (
		args    = os.Args[1:]
		content string
		token   *tokenizer.Token
	)
	if len(args) < 1 {
		content = inputparser.ParseCLI()
	} else if len(args) == 1 {
		content = inputparser.ParseFile(args[0])
	} else {
		log.Panicln("Invalid command")
	}
	token = tokenizer.Tokenize(content)
	printTokenList(token)
}

func printTokenList(head *tokenizer.Token) {
	for curr := head; curr != nil; curr = curr.Next {
		fmt.Printf("%s ", curr.Val)
	}
	fmt.Println()
}
