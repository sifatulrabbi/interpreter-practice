package parser

import (
	"fmt"
	"testing"

	"funlang/ast"
	"funlang/lexer"
)

func TestLetStatements(t *testing.T) {
	input := `let x = 5;
    let y = 10;
    let foobar = 838383;`
	// input = `let x 5;
	//    let = 10;
	//    let 838383;`

	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	if program == nil {
		t.Fatal("ParseProgram() returned nil")
	}

	checkParserErrors(t, p)

	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contina 3 statements got=%d",
			len(program.Statements))
	}
	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}
	for i, tt := range tests {
		s := program.Statements[i]
		if !testLetStatement(t, s, tt.expectedIdentifier) {
			break
		}
	}
}

func TestReturnStatements(t *testing.T) {
	input := `return 5;
    return 10;
    return aValue`

	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	if program == nil {
		t.Fatal("ParseProgram() returned nil")
	}

	checkParserErrors(t, p)

	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. got=%d",
			len(program.Statements))
	}

	for _, stmt := range program.Statements {
		returnStmt, ok := stmt.(*ast.ReturnStatement)
		if !ok {
			t.Errorf("stmt not *ast.returnStatement. got=%T", stmt)
			continue
		}
		if returnStmt.TokenLiteral() != "return" {
			t.Errorf("returnStmt.TokenLiteral not 'return', got %q",
				returnStmt.TokenLiteral())
		}
	}
}

func TestIdentifierExpression(t *testing.T) {
	input := `foobar;`

	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	if program == nil {
		t.Fatal("ParseProgram() returned nil")
	}

	checkParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("Program does not have the right amount statements. expected 1 got=%d\n",
			len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T\n",
			program.Statements[0])
	}

	ident, ok := stmt.Expression.(*ast.Identifier)
	if !ok {
		t.Fatalf("expression is not *ast.Identifier. got=%T\n",
			stmt.Expression)
	}
	if ident.Value != "foobar" {
		t.Errorf("ident.Value not %s. got=%s", "foobar", ident.Value)
	}
	if ident.TokenLiteral() != "foobar" {
		t.Errorf("ident.TokenLiteral not %s. got=%s", "foobar",
			ident.TokenLiteral())
	}
}

func TestIntegerLiteralExpressions(t *testing.T) {
	input := "5;"
	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	if program == nil {
		t.Fatal("ParseProgram() returned nil")
	}

	checkParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("program does not have right amount of statements. expect 1, got=%d\n",
			len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not of *ast.ExpressionStatement. got=%T\n", stmt)
	}
	literal, ok := stmt.Expression.(*ast.IntegerLiteral)
	if !ok {
		t.Fatalf("stmt.Expression is not of *ast.IntegerLiteral. got=%T\n", stmt)
	}
	if literal.Value != 5 {
		t.Errorf("literal.Value not %d. got=%d\n", 5, literal.Value)
	}
	if literal.TokenLiteral() != "5" {
		t.Errorf("literal.TokenLiteral not %s. got=%s\n",
			"5", literal.TokenLiteral())
	}
}

func TestParsingPrefixOperators(t *testing.T) {
	prefixTests := []struct {
		input    string
		operator string
		value    int64
	}{
		{"!5", "!", 5},
		{"-15", "-", 15},
	}

	for _, tt := range prefixTests {
		l := lexer.New(tt.input)
		p := New(l)
		program := p.ParseProgram()
		if program == nil {
			t.Fatalf("p.ParseProgram() returned nil")
			return
		}
		checkParserErrors(t, p)

		if len(program.Statements) != 1 {
			t.Fatalf("incorrect amount of Statements. expect=1 got=%d\n", len(program.Statements))
		}

		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
		if !ok {
			t.Fatalf("program.Statements[0] is not *ast.ExpressionStatement. got=%T\n",
				program.Statements[0])
		}

		exp, ok := stmt.Expression.(*ast.PrefixExpression)
		if !ok {
			t.Fatalf("stmt is not *ast.PrefixExpression. got=%T\n", stmt)
		}
		if exp.Operator != tt.operator {
			t.Fatalf("exp.Operator is not %s. got=%q\n",
				tt.operator, exp.Operator)
		}
		if !testIntegerLiteral(t, exp.Right, tt.value) {
			return
		}
	}
}

func TestParsingInfixExpressions(t *testing.T) {
	infixTests := []struct {
		input      string
		leftValue  int64
		operator   string
		rightValue int64
	}{
		{"5 + 5;", 5, "+", 5},
		{"5 - 5;", 5, "-", 5},
		{"5 * 5;", 5, "*", 5},
		{"5 / 5;", 5, "/", 5},
		{"5 > 5;", 5, ">", 5},
		{"5 < 5;", 5, "<", 5},
		{"5 == 5;", 5, "==", 5},
		{"5 != 5;", 5, "!=", 5},
	}

	for _, tt := range infixTests {
		l := lexer.New(tt.input)
		p := New(l)
		program := p.ParseProgram()
		if program == nil {
			t.Fatalf("p.ParseProgram() returned nil")
			t.FailNow()
			return
		}
		checkParserErrors(t, p)
		if len(program.Statements) != 1 {
			t.Fatalf("program.Statements does not contain accurate amount of statements. expect=%d got=%d",
				1, len(program.Statements))
		}

		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
		if !ok {
			t.Fatalf("program.Statements[0] is not *ast.ExpressionStatement. got=%T",
				program.Statements[0])
		}
		exp, ok := stmt.Expression.(*ast.InfixExpression)
		if !ok {
			t.Fatalf("stmt.Expression is not *ast.InfixExpression. got=%T",
				stmt.Expression)
		}
		if !testIntegerLiteral(t, exp.Left, tt.leftValue) {
			return
		}
		if exp.Operator != tt.operator {
			t.Fatalf("exp.Operator is not '%s'. got=%q", tt.operator, exp.Operator)
		}
		if !testIntegerLiteral(t, exp.Right, tt.rightValue) {
		}
	}
}

func TestOperatorPrecedenceParsing(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			"-a * b",
			"((-a) * b)",
		},
		{
			"!-a",
			"(!(-a))",
		},
		{
			"a + b + c",
			"((a + b) + c)",
		},
		{
			"a + b - c",
			"((a + b) - c)",
		},
		{
			"a * b * c",
			"((a * b) * c)",
		},
		{
			"a * b / c",
			"((a * b) / c)",
		},
		{
			"a + b / c",
			"(a + (b / c))",
		},
		{
			"a + b * c + d / e - f",
			"(((a + (b * c)) + (d / e)) - f)",
		},
		{
			"3 + 4; -5 * 5",
			"(3 + 4)((-5) * 5)",
		},
		{
			"5 > 4 == 3 < 4",
			"((5 > 4) == (3 < 4))",
		},
		{
			"5 < 4 != 3 > 4",
			"((5 < 4) != (3 > 4))",
		},
		{
			"3 + 4 * 5 == 3 * 1 + 4 * 5",
			"((3 + (4 * 5)) == ((3 * 1) + (4 * 5)))",
		},
		{
			"3 + 4 * 5 == 3 * 1 + 4 * 5",
			"((3 + (4 * 5)) == ((3 * 1) + (4 * 5)))",
		},
	}

	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := New(l)
		program := p.ParseProgram()
		checkParserErrors(t, p)

		actual := program.String()
		if actual != tt.expected {
			t.Errorf("expected=%q, got=%q", tt.expected, actual)
		}
	}
}

func testLetStatement(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "let" {
		t.Errorf("s.TokenLietral() is not 'let'. got=%q", s.TokenLiteral())
		return false
	}

	letStmt, ok := s.(*ast.LetStatement)

	if !ok {
		t.Errorf("s is not *ast.LetStatement. got=%T", s)
		return false
	}
	if letStmt.Name.Value != name {
		t.Errorf("letStmt.Name.Value is not '%s'. got='%s'",
			name, letStmt.Name.Value)
		return false
	}
	if letStmt.Name.TokenLiteral() != name {
		t.Errorf("letStmt.Name.Value not '%s'. got='%s'",
			name, letStmt.Name.TokenLiteral())
		return false
	}

	return true
}

func checkParserErrors(t *testing.T, p *Parser) {
	errors := p.Errors()
	if len(errors) < 1 {
		return
	}
	t.Errorf("parser has %d errors\n", len(errors))
	for _, v := range errors {
		t.Errorf("parser error: %q\n", v)
	}
	t.FailNow()
}

func testIntegerLiteral(t *testing.T, il ast.Expression, v int64) bool {
	integer, ok := il.(*ast.IntegerLiteral)
	if !ok {
		t.Errorf("il is not *ast.IntegerLiteral. got=%T\n", il)
		return false
	}
	if integer.Value != v {
		t.Errorf("integer.Value is not %d. got=%d\n", v, integer.Value)
	}
	if integer.TokenLiteral() != fmt.Sprintf("%d", v) {
		t.Errorf("integer.TokenLiteral() is not %d. got=%s\n", v, integer.TokenLiteral())
	}
	return true
}
