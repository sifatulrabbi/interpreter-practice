package mathparser

import (
	"log"
	"slices"
)

type Operation struct {
	Val    string
	Left   int
	Right  int
	Action string
}

type Token struct {
	Type    string
	Literal string
}

var (
	operationFns = map[string]func(l, r int) int{}
	numbersTable = map[string]int{
		"0": 0,
		"1": 1,
		"2": 2,
		"3": 3,
		"4": 4,
		"5": 5,
		"6": 6,
		"7": 7,
		"8": 8,
		"9": 9,
	}
	operationsTabble = map[string]string{
		"+": "+",
		"-": "-",
		"/": "/",
		"*": "*",
		"%": "%",
	}
)

func Lexer(input string) []Token {
	tokens := []Token{}
	t := Token{}
	for i := 0; i < len(input); i++ {
		v := string(input[i])
		t.Literal += v
		if i+1 < len(input) && slices.Contains([]string{" ", "\n", "\n\r"}, string(input[i+1])) {
			tokens = append(tokens, t)
			t = Token{}
			i++
		}
	}
	return tokens
}

func init() {
	operationFns["+"] = func(l, r int) int {
		return l + r
	}

	operationFns["-"] = func(l, r int) int {
		return l - r
	}

	operationFns["*"] = func(l, r int) int {
		return l * r
	}

	operationFns["/"] = func(l, r int) int {
		return l / r
	}

	operationFns["%"] = func(l, r int) int {
		return l % r
	}

	// operationFns["**"] = func(l, r int) int {
	// 	return l ^ r
	// }
}

func ExecOperation(o Operation) int {
	fn, ok := operationFns[o.Action]
	if !ok {
		log.Panicf("Invalid or unsupported action: %s\n", o.Action)
	}
	return fn(o.Left, o.Right)
}
