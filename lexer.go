package lexer

type Scanner struct {
	in  string
	pos int
}

const (
	ILLEGAL_TOKEN = iota
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
)

func MakeScanner(str string) Scanner {
	return Scanner{
		str,
		0}
}

func (s *Scanner) Lex() int {
	for {
		curChar := s.in[s.pos]
		switch curChar {
		case '+':
			s.pos++
			return PLUS
		case '-':
			s.pos++
			return MINUS
		case '*':
			s.pos++
			return TIMES
		case '/':
			s.pos++
			return DIV
		case '(':
			s.pos++
			return LPAREN
		case ')':
			s.pos++
			return RPAREN
		case ':':
			s.pos++
			if s.in[s.pos] == '=' {
				s.pos++
				return ASSIGN
			} else {
				return ILLEGAL_TOKEN
			}
		case '.':
			s.pos++
			if !isDigit(curChar) {
				return ILLEGAL_TOKEN
			}
			peek = s.pos
			for ; isDigit(s.in[peek]); peek++ {
			}
			s.pos = peek
		case '/':
			if curChar == '/' {
				for s.pos++; s.in[s.pos] != '\n'; s.pos++ {
				}
				continue
			} else if curChar == '*' {
				for s.pos++; !(s.in[pos] == '*' && s.in[pos+1] == '/'); s.pos++ {
				}
				continue
			}
			return DIV
		default:
			if isLetter(curChar) {
				peek := s.pos
				for ; isLetter(s.in[s.peek]) || isDigit(peekChar); peek++ {
				}
				word := s.in[pos:peek]
				pos = peek
				switch word {
				case "read":
					return READ
				case "write":
					return WRITE
				default:
					return ID
				}
			}
		}
	}
}

func isLetter(r rune) bool {
	return ('a' <= r && r <= 'z') || ('A' <= r && r <= 'Z')
}

func isDigit(r rune) bool {
	return '0' <= r && r <= '9'
}
