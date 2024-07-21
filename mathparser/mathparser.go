package mathparser

import (
	"log"
)

type TokenType string

type Operation struct {
	Val    string
	Left   int
	Right  int
	Action string
}

type Token struct {
	Type    TokenType
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

const (
	MATH_OP TokenType = "MATH_OP"
	INTEGER TokenType = "INTEGER"
)

func ParseTokens(input string) []Token {
	tokens := []Token{}
	for i := 0; i < len(input); i++ {
		v := string(input[i])
		if shouldIgnore(v) {
			continue
		}
		tt := getTokenType(v)
		tokens = append(tokens, Token{Type: tt, Literal: v})
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

func getTokenType(s string) TokenType {
	if _, ok := operationsTabble[s]; ok {
		return MATH_OP
	}
	if _, ok := numbersTable[s]; ok {
		return INTEGER
	}
	log.Panicf("Invalid character '%s'\n", s)
	return ""
}

func shouldIgnore(s string) bool {
	ignore := false
	switch s {
	case "\n":
		ignore = true
		break
	case "\n\r":
		ignore = true
		break
	case " ":
		ignore = true
		break
	}
	return ignore
}
