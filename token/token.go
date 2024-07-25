package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

// Tokens
const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers + literals
	IDENT  = "IDENT"  // identifier var and function names
	INT    = "INT"    // 12345...
	STRING = "STRING" // abcde...

	// Operators
	ASSIGN    = "="
	PLUS      = "+"
	MINUS     = "-"
	SLASH     = "/"
	REMAINDER = "%"
	ASTERISK  = "*"
	BANG      = "!"
	LT        = "<"
	GT        = ">"
	EQUAL     = "=="
	NOT_EQUAL = "!="
	LT_EQUAL  = "<="
	GT_EQUAL  = ">="

	// Delimeters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	FUNCTION = "fn"
	LET      = "let"
	RETURN   = "return"
	FOR      = "for"
	BREAK    = "break"
	CONTINUE = "continue"
	IF       = "if"
	ELSE     = "else"
	TRUE     = "true"
	FALSE    = "false"
)

var keywords = map[string]TokenType{
	"fn":       FUNCTION,
	"let":      LET,
	"return":   RETURN,
	"for":      FOR,
	"if":       IF,
	"else":     ELSE,
	"break":    BREAK,
	"continue": CONTINUE,
	"true":     TRUE,
	"false":    FALSE,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
