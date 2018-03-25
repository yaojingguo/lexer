package lexer

type Scanner struct {
	in  string
	pos int
}

const (
	EOF = iota
	ASSIGN
	PLUS
	MINUS
	TIMES
	DIV
	LPAREN
	RPAREN
	ID
	NUMBER
	READ
	WRITE
	ILLEGAL_TOKEN
)

func MakeScanner(str string) Scanner {
	return Scanner{
		str,
		0}
}

type Lexeme struct {
	token int
	str   string
}

func MakeLexeme(token int, str string) Lexeme {
	return Lexeme{token, str}
}

func MakeToken(token int) Lexeme {
	return Lexeme{token, ""}
}

func (s *Scanner) Lex() Lexeme {
	for {
		if s.pos >= len(s.in) {
			return MakeToken(EOF)
		}
		curChar := s.in[s.pos]
		switch curChar {
		case ' ':
			s.pos++
			continue
		case '\t':
			s.pos++
			continue
		case '\n':
			s.pos++
			continue
		case '+':
			s.pos++
			return MakeToken(PLUS)
		case '-':
			s.pos++
			return MakeToken(MINUS)
		case '*':
			s.pos++
			return MakeToken(TIMES)
		case '(':
			s.pos++
			return MakeToken(LPAREN)
		case ')':
			s.pos++
			return MakeToken(RPAREN)
		case ':':
			s.pos++
			if s.pos < len(s.in) && s.in[s.pos] == '=' {
				s.pos++
				return MakeToken(ASSIGN)
			} else {
				return MakeLexeme(ILLEGAL_TOKEN, s.in[s.pos-1:s.pos])
			}
		case '.':
			peek := s.pos + 1
			if !isDigit(s.in[peek]) {
				return MakeLexeme(ILLEGAL_TOKEN, s.in[s.pos:peek+1])
			}
			peek++
			for ; peek < len(s.in) && isDigit(s.in[peek]); peek++ {
			}
			oldPos := s.pos
			s.pos = peek
			return MakeLexeme(NUMBER, s.in[oldPos:s.pos])
		case '/':
			backslashPos := s.pos
			s.pos++
			if s.pos < len(s.in) {
				if s.in[s.pos] == '/' {
					for s.pos++; s.pos < len(s.in) && s.in[s.pos] != '\n'; s.pos++ {
					}
					if s.pos >= len(s.in) {
						return MakeLexeme(ILLEGAL_TOKEN, s.in[backslashPos:])
					} else {
						s.pos++
						continue
					}
				} else if s.in[s.pos] == '*' {
					for s.pos++; s.pos < len(s.in)-1 && !(s.in[s.pos] == '*' && s.in[s.pos+1] == '/'); s.pos++ {
					}
					if s.pos >= len(s.in)-1 {
						return MakeLexeme(ILLEGAL_TOKEN, s.in[backslashPos:])
					} else {
						s.pos = s.pos + 2
						continue
					}
				}
			}
			return MakeToken(DIV)
		default:
			if isLetter(curChar) {
				peek := s.pos + 1
				for ; peek < len(s.in) && (isLetter(s.in[peek]) || isDigit(s.in[peek])); peek++ {
				}
				word := s.in[s.pos:peek]
				s.pos = peek
				switch word {
				case "read":
					return MakeToken(READ)
				case "write":
					return MakeToken(WRITE)
				default:
					return MakeLexeme(ID, word)
				}
			} else if isDigit(curChar) {
				digitPos := s.pos
				for s.pos++; s.pos < len(s.in) && isDigit(s.in[s.pos]); s.pos++ {
				}
				if s.pos >= len(s.in) || s.in[s.pos] != '.' {
					return MakeLexeme(NUMBER, s.in[digitPos:s.pos])
				}
				for s.pos++; s.pos < len(s.in) && isDigit(s.in[s.pos]); s.pos++ {
				}
				return MakeLexeme(NUMBER, s.in[digitPos:s.pos])
			}
			return MakeLexeme(ILLEGAL_TOKEN, s.in[s.pos:s.pos+1])
		}
	}
}

func isLetter(r byte) bool {
	return ('a' <= r && r <= 'z') || ('A' <= r && r <= 'Z')
}

func isDigit(r byte) bool {
	return '0' <= r && r <= '9'
}
