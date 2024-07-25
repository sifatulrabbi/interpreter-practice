package repl

import (
	"bufio"
	"fmt"
	"io"

	"funlang/lexer"
	"funlang/token"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	sc := bufio.NewScanner(in)
	for {
		fmt.Printf(PROMPT)
		if !sc.Scan() {
			return
		}
		l := lexer.New(sc.Text())
		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
