package lexer

import "testing"
import "reflect"

func TestSingleToken(t *testing.T) {
	testData := []struct {
		cal      string
		expected Lexeme
	}{
		{":=", MakeToken(ASSIGN)},
		{"+", MakeToken(PLUS)},
		{"-", MakeToken(MINUS)},
		{"*", MakeToken(TIMES)},
		{"/", MakeToken(DIV)},
		{"(", MakeToken(LPAREN)},
		{")", MakeToken(RPAREN)},
		{"count1", MakeLexeme(ID, "count1")},
		{"0", MakeLexeme(NUMBER, "0")},
		{".1", MakeLexeme(NUMBER, ".1")},
		{"1.", MakeLexeme(NUMBER, "1.")},
		{"1.1", MakeLexeme(NUMBER, "1.1")},
		{"read", MakeToken(READ)},
		{"write", MakeToken(WRITE)},
	}
	for i, d := range testData {
		s := MakeScanner(d.cal)
		lexeme := s.Lex()
		if !reflect.DeepEqual(lexeme, d.expected) {
			t.Errorf("%d: %q: expected %v, but found %v", i, d.cal, d.expected, lexeme)
		}
	}
}

func TestComments(t *testing.T) {
	testData := []struct {
		cal      string
		expected Lexeme
	}{
		{"//line-comment\n+", MakeToken(PLUS)},
		{"/* line-1\nline-2 */+", MakeToken(PLUS)},
	}
	for i, d := range testData {
		s := MakeScanner(d.cal)
		lexeme := s.Lex()
		if !reflect.DeepEqual(lexeme, d.expected) {
			t.Errorf("%d: %q: expected %v, but found %v", i, d.cal, d.expected, lexeme)
		}
	}
}

func TestScan(t *testing.T) {
	lines := `// a simple calculatation
					a := 10 + 30
					b := a
					/* line-1
					   line-2 */`
	testData := []struct {
		cal      string
		expected []Lexeme
	}{
		{`(1 + 2) / 10`, []Lexeme{MakeToken(LPAREN), MakeLexeme(NUMBER, "1"),
			MakeToken(PLUS), MakeLexeme(NUMBER, "2"), MakeToken(RPAREN), MakeToken(DIV),
			MakeLexeme(NUMBER, "10")}},
		{`a := 10`, []Lexeme{MakeLexeme(ID, "a"), MakeToken(ASSIGN), MakeLexeme(NUMBER, "10")}},
		{lines, []Lexeme{MakeLexeme(ID, "a"), MakeToken(ASSIGN), MakeLexeme(NUMBER,
			"10"), MakeToken(PLUS), MakeLexeme(NUMBER, "30"), MakeLexeme(ID, "b"),
			MakeToken(ASSIGN), MakeLexeme(ID, "a")}},
	}
	for i, d := range testData {
		s := MakeScanner(d.cal)
		var lexemes []Lexeme
		for {
			lexeme := s.Lex()
			t.Logf("lexeme: %v\n", lexeme)
			if lexeme.token == EOF {
				break
			} else if lexeme.token == ILLEGAL_TOKEN {
				t.Fatalf("illegal token: %s", lexeme.str)
			}
			lexemes = append(lexemes, lexeme)
		}
		if !reflect.DeepEqual(lexemes, d.expected) {
			t.Errorf("%d: %q: expected %v, but found %v", i, d.cal, d.expected, lexemes)
		}
	}
}

func TestErrorHandling(t *testing.T) {
	testData := []struct {
		cal      string
		expected []Lexeme
	}{
		{":+", []Lexeme{MakeLexeme(ILLEGAL_TOKEN, ":"), MakeToken(PLUS)}},
	}
	for i, d := range testData {
		s := MakeScanner(d.cal)
		var lexemes []Lexeme
		for {
			lexeme := s.Lex()
			if lexeme.token == EOF {
				break
			}
			lexemes = append(lexemes, lexeme)
		}
		if !reflect.DeepEqual(lexemes, d.expected) {
			t.Errorf("%d: %q: expected %v, but found %v", i, d.cal, d.expected, lexemes)
		}
	}
}
