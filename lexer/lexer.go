package lexer

import (
	"funlang/token"
)

type Lexer struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	currentChar  byte // current char under examination
}

func New(input string) *Lexer {
	l := &Lexer{input: input, position: 0, readPosition: 0, currentChar: 0}
	l.readChar()
	return l
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token
	l.skipWhitespace()
	switch l.currentChar {
	case '=':
		if l.peekChar() == '=' {
			ch := l.currentChar
			l.readChar()
			tok = newToken(token.EQUAL, ch, l.currentChar)
		} else {
			tok = newToken(token.ASSIGN, l.currentChar)
		}
		break
	case ';':
		tok = newToken(token.SEMICOLON, l.currentChar)
		break
	case '(':
		tok = newToken(token.LPAREN, l.currentChar)
		break
	case ')':
		tok = newToken(token.RPAREN, l.currentChar)
		break
	case ',':
		tok = newToken(token.COMMA, l.currentChar)
		break
	case '+':
		tok = newToken(token.PLUS, l.currentChar)
		break
	case '-':
		tok = newToken(token.MINUS, l.currentChar)
		break
	case '{':
		tok = newToken(token.LBRACE, l.currentChar)
		break
	case '}':
		tok = newToken(token.RBRACE, l.currentChar)
		break
	case '/':
		tok = newToken(token.SLASH, l.currentChar)
		break
	case '%':
		tok = newToken(token.REMAINDER, l.currentChar)
		break
	case '!':
		if l.peekChar() == '=' {
			ch := l.currentChar
			l.readChar()
			tok = newToken(token.NOT_EQUAL, ch, l.currentChar)
		} else {
			tok = newToken(token.BANG, l.currentChar)
		}
		break
	case '*':
		tok = newToken(token.ASTERISK, l.currentChar)
		break
	case '<':
		if l.peekChar() == '=' {
			ch := l.currentChar
			l.readChar()
			tok = newToken(token.LT_EQUAL, ch, l.currentChar)
		} else {
			tok = newToken(token.LT, l.currentChar)
		}
		break
	case '>':
		if l.peekChar() == '=' {
			ch := l.currentChar
			l.readChar()
			tok = newToken(token.GT_EQUAL, ch, l.currentChar)
		} else {
			tok = newToken(token.GT, l.currentChar)
		}
		break
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
		break
	default:
		if isLetter(l.currentChar) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		}
		if isDigit(l.currentChar) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
		}
		tok = newToken(token.ILLEGAL, l.currentChar)
		break
	}
	l.readChar()
	return tok
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		// 0 is the ASCII code for 'NUL' character
		l.currentChar = 0
	} else {
		l.currentChar = l.input[l.readPosition]
	}
	// since we got a char we will now forward to the next position
	l.position = l.readPosition
	l.readPosition++
}

func (l *Lexer) readIdentifier() string {
	pos := l.position
	for isLetter(l.currentChar) {
		l.readChar()
	}
	return l.input[pos:l.position]
}

func (l *Lexer) readNumber() string {
	pos := l.position
	for isDigit(l.currentChar) {
		l.readChar()
	}
	return l.input[pos:l.position]
}

func (l *Lexer) skipWhitespace() {
	for l.currentChar == ' ' || l.currentChar == '\n' || l.currentChar == '\r' || l.currentChar == '\t' {
		l.readChar()
	}
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

func newToken(tokenType token.TokenType, ch ...byte) token.Token {
	str := ""
	for _, v := range ch {
		str += string(v)
	}
	return token.Token{Type: tokenType, Literal: str}
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
