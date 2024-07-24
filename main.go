package main

import (
	"bufio"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) < 1 || args[0] == "" {
		sc := bufio.NewScanner(os.Stdin)
		for sc.Scan() {
			text := sc.Text()
			if text == "exit" {
				break
			}
		}
		return
	}
}
