package lexer

import (
	"log"
	"strings"
	"unicode"

	"github.com/DiogoGaspar-cell/monkey2/token"
)

type Lexer struct {
	filename string
	reader   *strings.Reader
	line     int
	position int
	ch       rune
}

func New(filename string, input string) *Lexer{
	r := strings.NewReader(input)

	l := &Lexer{
		filename: filename,
		reader: r,
		line: 1,
		position: -1,
	}
	
	l.readRune()

	return l
}

func (l *Lexer) readRune() {
	r, _, err := l.reader.ReadRune()

	if err != nil {
		log.Printf("Error: %v", err)
		l.ch = 0
		return
	}

	l.ch = r
	l.position += 1
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {
	case '=':
		if l.peekRune() == '=' {
			ch := l.ch
			l.readRune()
			literal := string(ch) + string(l.ch)
			tok = token.Token{
				Type: token.EQ,
				Literal: literal,
				Filename: l.filename,
				Line: l.line,
				Col: l.position,
			}
		} else {
			tok = newToken(token.ASSIGN, l.ch, l.filename, l.line, l.position)
		}
	case '+':
		tok = newToken(token.PLUS, l.ch, l.filename, l.line, l.position)
	case '-':
		tok = newToken(token.MINUS, l.ch, l.filename, l.line, l.position)
	case '!':
		if l.peekRune() == '=' {
			ch := l.ch
			l.readRune()
			literal := string(ch) + string(l.ch)
			tok = token.Token{
				Type: token.NOT_EQ,
				Literal: literal,
				Filename: l.filename,
				Line: l.line,
				Col: l.position,
			}
		} else {
			tok = newToken(token.BANG, l.ch, l.filename, l.line, l.position)
		}
	case '/':
		tok = newToken(token.SLASH, l.ch, l.filename, l.line, l.position)
	case '*':
		tok = newToken(token.ASTERISK, l.ch, l.filename, l.line, l.position)
	case '<':
		if l.peekRune() == '=' {
			ch := l.ch
			l.readRune()
			literal := string(ch) + string(l.ch)
			tok = token.Token{
				Type: token.LTE, 
				Literal: literal,
				Filename: l.filename,
				Line: l.line,
				Col: l.position,
			}
		} else {
			tok = newToken(token.LT, l.ch, l.filename, l.line, l.position)
		}
	case '>':
		if l.peekRune() == '=' {
			ch := l.ch
			l.readRune()
			literal := string(ch) + string(l.ch)
			tok = token.Token{
				Type: token.GTE, 
				Literal: literal,
				Filename: l.filename,
				Line: l.line,
				Col: l.position,
			}
		} else {
			tok = newToken(token.GT, l.ch, l.filename, l.line, l.position)
		}
	case ';':
		tok = newToken(token.TERMINATOR, l.ch, l.filename, l.line, l.position)
	case ',':
		tok = newToken(token.COMMA, l.ch, l.filename, l.line, l.position)
	case '{':
		tok = newToken(token.LBRACE, l.ch, l.filename, l.line, l.position)
	case '}':
		tok = newToken(token.RBRACE, l.ch, l.filename, l.line, l.position)
	case '(':
		tok = newToken(token.LPAREN, l.ch, l.filename, l.line, l.position)
	case ')':
		tok = newToken(token.RPAREN, l.ch, l.filename, l.line, l.position)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			tok.Col = l.position
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)

			tok.Filename, tok.Line = l.filename, l.line
			
			return tok
		} else if isDigit(l.ch) {
			tok.Filename, tok.Line, tok.Col = l.filename, l.line, l.position
	
			tok.Literal, tok.Type = l.readNumber()

			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch, l.filename, l.line, l.position)
		}
	}

	l.readRune()

	return tok
}

func isLetter(ch rune) bool {
	return unicode.IsLetter(ch) || ch == '_'
}

func isDigit(ch rune) bool {
	return unicode.IsDigit(ch)
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		if (l.ch == '\n') {
			l.line++
			l.position = -1
		}
		l.readRune()
	}
}

func (l *Lexer) peekRune() rune {
	r, _, err := l.reader.ReadRune()

	if err != nil {
		log.Printf("Error: %v", err)
		return 0
	}
	
	err = l.reader.UnreadRune()

	if err != nil {
		log.Printf("Error: %v", err)
		return 0
	}

	return r
}

func (l *Lexer) readIdentifier() string {
	var identifier string

	for isLetter(l.ch) {
		identifier += string(l.ch)
		l.readRune()
	}

	return identifier
}

func (l *Lexer) readNumber() (string, token.TokenType) {
	var number string

	for isDigit(l.ch) {
		number += string(l.ch)
		l.readRune()
	}

	if l.ch != '.' && !isDigit(l.peekRune()) {
		return number, token.INT
	}

	number += string(l.ch)
	l.readRune()
	
	for isDigit(l.ch) {
		number += string(l.ch)
		l.readRune()
	}

	return number, token.FLOAT
}

func newToken(tokenType token.TokenType, ch rune, filename string, line int, col int) token.Token {
	return token.Token{
		Type: tokenType,
		Literal: string(ch),
		Filename: filename,
		Line: line,
		Col: col,
	}
}
