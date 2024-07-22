package mathparser

import (
	"fmt"
	"testing"
)

func TestTokenization(t *testing.T) {
	inputStr := "2+2\n3/3\n5-7\n6*8"
	tokens := ParseTokens(inputStr)
	fmt.Println(tokens)

	equations := BuildEquations(tokens)
	fmt.Println(equations)

	for _, e := range equations {
		res := ExecOperation(e)
		fmt.Printf("%s %s %s = %d\n", e.Left.Literal, e.Action.Literal, e.Right.Literal, res)
	}
}
