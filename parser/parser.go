package parser

import (
	"fmt"
	"strconv"

	"funlang/ast"
	"funlang/lexer"
	"funlang/token"
)

const (
	_ int = iota
	LOWEST
	EQUALS      // ==
	LESSGREATER // > or <
	SUM         // +
	PRODUCT     // *
	PREFIX      // -X or !X
	CALL        // myFunction(X)
)

var precedences = map[token.TokenType]int{
	token.EQUAL:     EQUALS,
	token.NOT_EQUAL: EQUALS,
	token.LT:        LESSGREATER,
	token.GT:        LESSGREATER,
	token.PLUS:      SUM,
	token.MINUS:     SUM,
	token.SLASH:     PRODUCT,
	token.ASTERISK:  PRODUCT,
}

type (
	prefixParseFn func() ast.Expression
	infixParseFn  func(ast.Expression) ast.Expression
)

type Parser struct {
	l         *lexer.Lexer
	curToken  token.Token
	peekToken token.Token // the immediate next token or EOR/nil
	errors    []string    // list of all the encountered errors in the program during parsing

	prefixParseFns map[token.TokenType]prefixParseFn
	infixParseFns  map[token.TokenType]infixParseFn
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l, errors: []string{}}

	p.prefixParseFns = make(map[token.TokenType]prefixParseFn)
	p.registerPrefix(token.IDENT, p.parseIdentifier)
	p.registerPrefix(token.INT, p.parseIntegerLiteral)
	p.registerPrefix(token.BANG, p.parsePrefixExpression)
	p.registerPrefix(token.MINUS, p.parsePrefixExpression)

	p.infixParseFns = make(map[token.TokenType]infixParseFn)
	p.registerInfix(token.PLUS, p.parseInfixExpression)
	p.registerInfix(token.MINUS, p.parseInfixExpression)
	p.registerInfix(token.SLASH, p.parseInfixExpression)
	p.registerInfix(token.ASTERISK, p.parseInfixExpression)
	p.registerInfix(token.EQUAL, p.parseInfixExpression)
	p.registerInfix(token.NOT_EQUAL, p.parseInfixExpression)
	p.registerInfix(token.LT, p.parseInfixExpression)
	p.registerInfix(token.GT, p.parseInfixExpression)

	// calling the nextToken() twice will first update the peekToken later update the curToken
	p.nextToken()
	p.nextToken()
	return p
}

// go through all the tokens and build the structure of the program (ast)
func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{Statements: []ast.Statement{}}
	for p.curToken.Type != token.EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}
	return program
}

// get all the errors during parsing process
func (p *Parser) Errors() []string {
	return p.errors
}

// moves the point to the next available token
func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

// validates the next token's token type and if the next token is not the
// expected token then it creates a new parsing error
func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	}
	p.peekError(t)
	return false
}

// validates the TokenType of the curToken
func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.curToken.Type == t
}

// validates the TokenType of the peekToken which is the immediate next token of the current token
func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

// get the precedence for next token
func (p *Parser) peekPrecedence() int {
	if p, ok := precedences[p.peekToken.Type]; ok {
		return p
	}
	return LOWEST
}

// get the precedence for current token
func (p *Parser) curPrecedence() int {
	if p, ok := precedences[p.curToken.Type]; ok {
		return p
	}
	return LOWEST
}

// adds a new prefix parsing func to the Prefix.prefixParseFns table
func (p *Parser) registerPrefix(t token.TokenType, fn prefixParseFn) {
	p.prefixParseFns[t] = fn
}

// adds a new infix parsing func to the Prefix.infixParseFns table
func (p *Parser) registerInfix(t token.TokenType, fn infixParseFn) {
	p.infixParseFns[t] = fn
}

// creates a new error when the peeking/next token is not of the expected type
func (p *Parser) peekError(t token.TokenType) {
	msg := fmt.Sprintf("expected next token to be '%s' but got '%s'",
		t, p.peekToken.Type)
	p.errors = append(p.errors, msg)
}

// create a new errors for unknown prefix parsing
func (p *Parser) noPrefixFnErr(t token.TokenType) {
	msg := fmt.Sprintf("no prefix parse function found for %s", t)
	p.errors = append(p.errors, msg)
}

// parse the statements from the current token till the next semicolon
func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.LET:
		return p.parseLetStatement()
	case token.RETURN:
		return p.parseReturnStatement()
	default:
		return p.parseExpressionStatement()
	}
}

// parse a let statement
func (p *Parser) parseLetStatement() *ast.LetStatement {
	stmt := &ast.LetStatement{Token: p.curToken}

	if !p.expectPeek(token.IDENT) {
		return nil
	}

	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	if !p.expectPeek(token.ASSIGN) {
		return nil
	}

	// TODO: skip parsing the expressions till we encounter semicolon
	// this will be implemented later
	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

// parse a return statement
func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	stmt := &ast.ReturnStatement{Token: p.curToken}

	// move to token that's after the return keyword
	// unlike LetStatement return statement don't always need an expression
	// because programs may have empty returns to stop a function execution
	p.nextToken()

	// TODO: skipping parsing the expressions until we hit a semicolon
	// this will be implemented later
	for p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}
	return stmt
}

// parse an expression statement
func (p *Parser) parseExpressionStatement() *ast.ExpressionStatement {
	stmt := &ast.ExpressionStatement{Token: p.curToken}
	stmt.Expression = p.parseExpression(LOWEST)
	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}
	return stmt
}

// func parse an identifier node and return an expression
// since identifiers are expressions too
func (p *Parser) parseIdentifier() ast.Expression {
	return &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
}

// parse and create expressions from integers
func (p *Parser) parseIntegerLiteral() ast.Expression {
	il := &ast.IntegerLiteral{Token: p.curToken}
	if val, err := strconv.ParseInt(p.curToken.Literal, 0, 64); err != nil {
		msg := fmt.Sprintf("could not parse %q as integer", p.curToken.Literal)
		p.errors = append(p.errors, msg)
	} else {
		il.Value = val
	}
	return il
}

// parse an expression using a prefixParseFn
func (p *Parser) parseExpression(precedence int) ast.Expression {
	prefix := p.prefixParseFns[p.curToken.Type]
	if prefix == nil {
		p.noPrefixFnErr(p.curToken.Type)
		return nil
	}
	leftExp := prefix()

	for !p.peekTokenIs(token.SEMICOLON) && precedence < p.peekPrecedence() {
		infix := p.infixParseFns[p.peekToken.Type]
		if infix == nil {
			break
		}
		p.nextToken()
		leftExp = infix(leftExp)
	}

	return leftExp
}

// parse a prefix expression e.g. it will parse '!true' by creating
// an expression for '!' and another one for 'true'
// a simple 'varname' is also parsed using this function
func (p *Parser) parsePrefixExpression() ast.Expression {
	pe := &ast.PrefixExpression{
		Token:    p.curToken,
		Operator: p.curToken.Literal,
	}

	p.nextToken()
	pe.Right = p.parseExpression(PREFIX)

	return pe
}

// parse a infix expression and all it's tokens
func (p *Parser) parseInfixExpression(left ast.Expression) ast.Expression {
	exp := &ast.InfixExpression{
		Token:    p.curToken,
		Operator: p.curToken.Literal,
		Left:     left,
	}
	precedence := p.curPrecedence()

	p.nextToken()
	exp.Right = p.parseExpression(precedence)

	return exp
}
