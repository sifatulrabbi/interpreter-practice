package mathparser

import (
	"fmt"
	"log"
	"strconv"
)

type TokenType string

type Equation struct {
	Result int
	Left   Token
	Right  Token
	Action Token
}

type Token struct {
	Type    TokenType
	Literal string
}

const (
	MATH_OP TokenType = "MATH_OP"
	INTEGER TokenType = "INTEGER"
)

var operationFns = map[string]func(l, r int) int{}

func ParseTokens(input string) []Token {
	var (
		tokens             = []Token{}
		buf                = ""
		prevType TokenType = ""
	)
	for i := 0; i < len(input); i++ {
		v := string(input[i])

		if shouldIgnore(v) {
			tokens = append(tokens, Token{Type: prevType, Literal: buf})
			prevType = ""
			buf = ""
			continue
		}

		tt := getTokenType(v)
		if prevType == "" {
			prevType = tt
		}

		if prevType == tt {
			buf += v
		} else {
			tokens = append(tokens, Token{Type: prevType, Literal: buf})
			prevType = tt
			buf = v
		}
	}
	if buf != "" {
		tokens = append(tokens, Token{Type: prevType, Literal: buf})
	}
	return tokens
}

func BuildEquations(tokens []Token) []Equation {
	operations := []Equation{}
	for i := 0; i < len(tokens); i++ {
		token := tokens[i]
		if token.Type != INTEGER {
			fmt.Printf("Equations should start with an integer, but found %s\n", token.Literal)
		}
		if i+2 > len(tokens)-1 {
			fmt.Printf("Incomplete math equation\n")
		}
		action := tokens[i+1]
		right := tokens[i+2]
		eq := Equation{
			Left:   token,
			Action: action,
			Right:  right,
		}
		operations = append(operations, eq)
		i += 2
	}
	return operations
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

func ExecOperation(o Equation) int {
	fn, ok := operationFns[o.Action.Literal]
	if !ok {
		log.Panicf("Invalid or unsupported action: %s\n", o.Action)
	}
	l, _ := strconv.ParseInt(o.Left.Literal, 10, 32)
	r, _ := strconv.ParseInt(o.Right.Literal, 10, 32)
	return fn(int(l), int(r))
}

func getTokenType(s string) TokenType {
	switch s {
	case "+":
		return MATH_OP
	case "-":
		return MATH_OP
	case "/":
		return MATH_OP
	case "*":
		return MATH_OP
	case "%":
		return MATH_OP
	case "0":
		return INTEGER
	case "1":
		return INTEGER
	case "2":
		return INTEGER
	case "3":
		return INTEGER
	case "4":
		return INTEGER
	case "5":
		return INTEGER
	case "6":
		return INTEGER
	case "7":
		return INTEGER
	case "8":
		return INTEGER
	case "9":
		return INTEGER
	}
	log.Panicf("Invalid character '%s'\n", s)
	return ""
}

func shouldIgnore(s string) bool {
	switch s {
	case "\n":
		return true
	case "\n\r":
		return true
	case " ":
		return true
	}
	return false
}
