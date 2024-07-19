package tokenizer

type Token struct {
	Val  string
	Prev *Token
	Next *Token
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
			if curr.Val != "" {
				curr = moveForward(curr, "")
			}
			break
		case ";":
			if curr.Val != "" {
				curr = moveForward(curr, "")
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
	eofToken := &Token{Val: "EOF", Prev: curr, Next: nil}
	curr.Next = eofToken
	return head
}

func moveForward(curr *Token, val string) *Token {
	t := &Token{Val: val, Prev: curr, Next: nil}
	curr.Next = t
	return t
}
