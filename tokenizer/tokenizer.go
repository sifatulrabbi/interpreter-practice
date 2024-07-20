package tokenizer

import (
	"fmt"
	"log"
)

const (
	LET         = "LET"
	SEMICOLON   = "SEMICOLON"
	ASSIGN      = "ASSIGN"
	BLOCK_START = "BLOCK_START"
	BLOCK_END   = "BLOCK_END"
	FUNCTION    = "FUNCTION"
	ADD         = "ADD"
	DIVIDE      = "DIVIDE"
	REMAINDER   = "REMAINDER"
	SUBTRACT    = "SUBTRACT"
	EQUAL       = "EQUAL"
	RETURN      = "RETURN"
	IF          = "IF"
	ELSE        = "ELSE"
	BREAK       = "BREAK"
	CONTINUE    = "CONTINUE"
)

var KnownTokens = map[string]string{
	"let":      "LET",
	";":        "SEMICOLON",
	"=":        "ASSIGN",
	"{":        "BLOCK_START",
	"}":        "BLOCK_END",
	"fn":       "FUNCTION",
	"+":        "ADD",
	"/":        "DIVIDE",
	"%":        "REMAINDER",
	"-":        "SUBTRACT",
	"==":       "EQUAL",
	"return":   "RETURN",
	"if":       "IF",
	"else":     "ELSE",
	"break":    "BREAK",
	"continue": "CONTINUE",
}

type Token struct {
	Val  string
	Prev *Token
	Next *Token
}

type Declaration struct {
	Type  string
	Name  string
	Value *Token
}

func (d Declaration) String() string {
	str := fmt.Sprintf("%s %s = ", d.Type, d.Name)
	for curr := d.Value; curr != nil; curr = curr.Next {
		str += curr.Val
	}
	return str
}

func (t Token) Next2() *Token {
	if t.Next != nil {
		return t.Next.Next
	}
	return nil
}

func (t Token) Prev2() *Token {
	if t.Prev != nil {
		return t.Prev.Prev
	}
	return nil
}

func (t Token) IsEOF() bool {
	if t.Val == "\n" || t.Val == "\n\r" {
		return true
	}
	return false
}

func Tokenize(str string) *Token {
	head := &Token{Val: "", Prev: nil, Next: nil}
	curr := head

	for i := 0; i < len(str); i++ {
		v := string(str[i])
		switch v {
		case " ":
			if curr.Val != "" && i+1 < len(str) {
				curr = moveForward(curr, "")
			}
			break
		case ";":
			if curr.Val != "" {
				curr = moveForward(curr, ";")
			}
			break
		case "{":
			curr = moveForward(curr, "{")
			break
		case "}":
			curr = moveForward(curr, "}")
			break
		case "(":
			curr = moveForward(curr, "(")
			break
		case ")":
			curr = moveForward(curr, ")")
			break
		case "[":
			curr = moveForward(curr, "[")
			break
		case "]":
			curr = moveForward(curr, "]")
			break
		case "\n":
			break
		case "\n\r":
			break
		default:
			curr.Val += v
			break
		}
	}
	if curr.Val == "" {
		curr.Val = "EOF"
	} else {
		eofToken := &Token{Val: "EOF", Prev: curr, Next: nil}
		curr.Next = eofToken
	}
	return head
}

func ValidateSyntax(token *Token) error {
	var err error = nil
	for curr := token; curr != nil; curr = curr.Next {
		switch curr.Val {
		case LET:
			curr = validateLetDeclaration(curr)
			break
		case SEMICOLON:
			break
		case ASSIGN:
			break
		case BLOCK_START:
			break
		case BLOCK_END:
			break
		case FUNCTION:
			break
		case ADD:
			break
		case DIVIDE:
			break
		case REMAINDER:
			break
		case SUBTRACT:
			break
		case EQUAL:
			break
		case RETURN:
			break
		default:
			break
		}
	}
	return err
}

func moveForward(curr *Token, val string) *Token {
	if curr.Val == "" {
		curr.Val = val
		return curr
	}
	t := &Token{Val: val, Prev: curr, Next: nil}
	curr.Next = t
	return t
}

func validateLetDeclaration(start *Token) *Token {
	var (
		curr = start
		d    = Declaration{}
	)
	if curr.Val != LET {
		return curr
	}
	d.Type = curr.Val
	curr = curr.Next
	if curr == nil {
		log.Panicln("Syntax error: Expecting a name after let declaration but found nothing")
	}
	return nil
}

func validateFunctionDeclaration(curr *Token) *Token {
	return nil
}
