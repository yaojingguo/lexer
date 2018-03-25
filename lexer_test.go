package lexer

import "testing"

func TestSingleToken(t *testing.T) {
	testData := []struct {
		cal      string
		expected int
	}{
		{":=", ASSIGN},
		{"+", PLUS},
		{"-", MINUS},
		{"*", TIMES},
		{"/", DIV},
		{"(", LPAREN},
		{")", RPAREN},
		{"count1", ID},
		{"0", NUMBER},
		{".1", NUMBER},
		{"1.", NUMBER},
		{"1.1", NUMBER},
		{"read", READ},
		{"write", WRITE},
	}
	for i, d := range testData {
		s := MakeScanner(d.cal)
		token := s.Lex()
		if token != d {
			t.Errorf("%d: %q: expected %d, but found %d", i, d.sql, d.expected, token)
		}
	}
}
