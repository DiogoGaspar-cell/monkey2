package lexer

import (
	"testing"

	"github.com/diogo-gaspar23/monkey2/token"
)

func TestNextToken(t *testing.T) {
	input := `let five = 5;
let ten = 10;

let add = fn(x, y) {
  x + y;
};

let lessThanOrEqualTo = fn(x, y) {
  x <= y;
};

let result = add(five, ten);

let 🤖 = 3.14;
`

	expectedFilename := "test.go"

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
		expectedFilename string
		expectedLine int
		expectedCol int
	}{
		{token.LET, "let", expectedFilename, 1, 0},
		{token.IDENTIFIER, "five", expectedFilename, 1, 4},
		{token.ASSIGN, "=", expectedFilename, 1, 9},
		{token.INT, "5", expectedFilename, 1, 11},
		{token.TERMINATOR, ";", expectedFilename, 1, 12},

		{token.LET, "let", expectedFilename, 2, 0},
		{token.IDENTIFIER, "ten", expectedFilename, 2, 4},
		{token.ASSIGN, "=", expectedFilename, 2, 8},
		{token.INT, "10", expectedFilename, 2, 10},
		{token.TERMINATOR, ";", expectedFilename, 2, 12},

		{token.LET, "let", expectedFilename, 4, 0},
		{token.IDENTIFIER, "add", expectedFilename, 4, 4},
		{token.ASSIGN, "=", expectedFilename, 4, 8},
		{token.FUNCTION, "fn", expectedFilename, 4, 10},
		{token.LPAREN, "(", expectedFilename, 4, 12},
		{token.IDENTIFIER, "x", expectedFilename, 4, 13},
		{token.COMMA, ",", expectedFilename, 4, 14},
		{token.IDENTIFIER, "y", expectedFilename, 4, 16},
		{token.RPAREN, ")", expectedFilename, 4, 17},
		{token.LBRACE, "{", expectedFilename, 4, 19},
		{token.IDENTIFIER, "x", expectedFilename, 5, 2},
		{token.PLUS, "+", expectedFilename, 5, 4},
		{token.IDENTIFIER, "y", expectedFilename, 5, 6},
		{token.TERMINATOR, ";", expectedFilename, 5, 7},
		{token.RBRACE, "}", expectedFilename, 6, 0},
		{token.TERMINATOR, ";", expectedFilename, 6, 1},

		{token.LET, "let", expectedFilename, 8, 0},
		{token.IDENTIFIER, "lessThanOrEqualTo", expectedFilename, 8, 4},
		{token.ASSIGN, "=", expectedFilename, 8, 22},
		{token.FUNCTION, "fn", expectedFilename, 8, 24},
		{token.LPAREN, "(", expectedFilename, 8, 26},
		{token.IDENTIFIER, "x", expectedFilename, 8, 27},
		{token.COMMA, ",", expectedFilename, 8, 28},
		{token.IDENTIFIER, "y", expectedFilename, 8, 30},
		{token.RPAREN, ")", expectedFilename, 8, 31},
		{token.LBRACE, "{", expectedFilename, 8, 33},
		{token.IDENTIFIER, "x", expectedFilename, 9, 2},
		{token.LTE, "<=", expectedFilename, 9, 4},
		{token.IDENTIFIER, "y", expectedFilename, 9, 7},
		{token.TERMINATOR, ";", expectedFilename, 9, 8},
		{token.RBRACE, "}", expectedFilename, 10, 0},
		{token.TERMINATOR, ";", expectedFilename, 10, 1},

		{token.LET, "let", expectedFilename, 12, 0},
		{token.IDENTIFIER, "result", expectedFilename, 12, 4},
		{token.ASSIGN, "=", expectedFilename, 12, 11},
		{token.IDENTIFIER, "add", expectedFilename, 12, 13},
		{token.LPAREN, "(", expectedFilename, 12, 16},
		{token.IDENTIFIER, "five", expectedFilename, 12, 17},
		{token.COMMA, ",", expectedFilename, 12, 21},
		{token.IDENTIFIER, "ten", expectedFilename, 12, 23},
		{token.RPAREN, ")", expectedFilename, 12, 26},
		{token.TERMINATOR, ";", expectedFilename, 12, 27},

		{token.LET, "let", expectedFilename, 14, 0},
		{token.IDENTIFIER, "🤖", expectedFilename, 14, 4},
		{token.ASSIGN, "=", expectedFilename, 14, 6},
		{token.FLOAT, "3.14", expectedFilename, 14, 8},
		{token.TERMINATOR, ";", expectedFilename, 14, 12},
	}

	l := New("test.go", input)

	for i, tt := range tests {
		
		tok := l.NextToken()
		
		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}

		if tok.Filename != tt.expectedFilename {
			t.Fatalf("tests[%d] - filename wrong. expected=%q, got=%q",
				i, tt.expectedFilename, tok.Filename)
		}

		if tok.Line != tt.expectedLine {
			t.Fatalf("tests[%d] - line wrong. expected=%d, got=%d",
				i, tt.expectedLine, tok.Line)
		}

		if tok.Col != tt.expectedCol {
			t.Fatalf("tests[%d] - column wrong. expected=%d, got=%d",
				i, tt.expectedCol, tok.Col)
		}
	}
}
