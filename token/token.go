package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers + literals
	IDENT  = "IDENT"  // identifier var and function names
	INT    = "INT"    // 12345..
	STRING = "STRING" // 12345..
	BOOL   = "BOOL"   // true false

	// Operators
	ASSIGN    = "="
	PLUS      = "+"
	MINUS     = "-"
	DIVIDE    = "/"
	REMAINDER = "%"

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
	IF       = "if"
	ELSE     = "else"
)

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"return": RETURN,
	"for":    FOR,
	"if":     IF,
	"else":   ELSE,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
