package lexer

import "github.com/MikhailWahib/Ra/token"

const (
	NUL     = 0
	NEWLINE = '\n'
)

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {
	// Brackets and braces
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case '[':
		tok = newToken(token.LBRACKET, l.ch)
	case ']':
		tok = newToken(token.RBRACKET, l.ch)

	// Comparison operators
	case '<':
		tok = newToken(token.LT, l.ch)
	case '>':
		tok = newToken(token.GT, l.ch)
	case '=':
		tok = l.makeCompoundToken('=', token.ASSIGN, token.EQ)
	case '!':
		tok = l.makeCompoundToken('=', token.BANG, token.NOT_EQ)

	// Arithmetic operators
	case '+':
		tok = l.makeCompoundToken('=', token.PLUS, token.PLUS_ASSIGN)
	case '-':
		tok = l.makeCompoundToken('=', token.MINUS, token.MINUS_ASSIGN)
	case '*':
		tok = l.makeCompoundToken('=', token.ASTERISK, token.ASTERISK_ASSIGN)
	case '/':
		tok = l.makeCompoundToken('=', token.SLASH, token.SLASH_ASSIGN)

	// Delimiters
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case ':':
		tok = newToken(token.COLON, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)

	// Special cases
	case '"':
		tok.Type = token.STRING
		tok.Literal = l.readString()
	case '#':
		l.skipComment()
		return l.NextToken()
	case NUL:
		tok.Literal = ""
		tok.Type = token.EOF

	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}

	l.readChar()
	return tok
}

func (l *Lexer) makeCompoundToken(nextChar byte, singleType, compoundType token.TokenType) token.Token {
	if l.peekChar() == nextChar {
		ch := l.ch
		l.readChar()
		return token.Token{Type: compoundType, Literal: string(ch) + string(l.ch)}
	}
	return newToken(singleType, l.ch)
}

func (l *Lexer) skipComment() {
	l.skipWhitespace()
	for l.ch != NEWLINE && l.ch != NUL {
		l.readChar()
	}
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readString() string {
	position := l.position + 1

	for {
		l.readChar()
		if l.ch == '"' || l.ch == 0 {
			break
		}
	}

	return l.input[position:l.position]
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}
