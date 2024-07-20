package tokenizer

import (
	"testing"
)

func TestTokenizationProcess(t *testing.T) {
	input := "let a = 1;"
	expectedTokens := map[string]bool{
		"let": false,
		"a":   false,
		"=":   false,
		"1":   false,
		";":   false,
		"EOF": false,
	}
	token := Tokenize(input)
	for curr := token; curr != nil; curr = curr.Next {
		if _, ok := expectedTokens[curr.Val]; !ok {
			t.Errorf("'%s' should not be in the token list. Prev: %s\n", curr.Val, curr.Prev.Val)
			t.FailNow()
		}
		expectedTokens[curr.Val] = true
	}
	for k, v := range expectedTokens {
		if !v {
			t.Errorf("'%s' was not found\n", k)
			t.FailNow()
		}
	}
}

func TestTokenizationWithComplexInput(t *testing.T) {
	input := `fn add() {
        return a + b;
    };`
	expectedTokens := map[string]bool{
		"fn":     false,
		"add":    false,
		"(":      false,
		")":      false,
		"{":      false,
		"return": false,
		"a":      false,
		"+":      false,
		"b":      false,
		";":      false,
		"}":      false,
		"EOF":    false,
	}
	token := Tokenize(input)
	for curr := token; curr != nil; curr = curr.Next {
		if _, ok := expectedTokens[curr.Val]; !ok {
			t.Errorf("'%s' should not be in the token list. Prev: %s\n", curr.Val, curr.Prev.Val)
			t.FailNow()
		}
		expectedTokens[curr.Val] = true
	}
	for k, v := range expectedTokens {
		if !v {
			t.Errorf("'%s' was not found\n", k)
			t.FailNow()
		}
	}
}
