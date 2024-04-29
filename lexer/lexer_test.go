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

!-/*5;
5 < 10 > 5;

if (5 < 10) {
	return true;
} else {
	return false;
}

10 == 10;
10 != 9;
`

	expectedFilename := "test.go"

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
		expectedFilename string
		expectedLine int
		expectedCol int
	}{
		{token.LET, "let", expectedFilename, 1, 1},
		{token.IDENTIFIER, "five", expectedFilename, 1, 5},
		{token.ASSIGN, "=", expectedFilename, 1, 10},
		{token.INT, "5", expectedFilename, 1, 12},
		{token.TERMINATOR, ";", expectedFilename, 1, 13},

		{token.LET, "let", expectedFilename, 2, 1},
		{token.IDENTIFIER, "ten", expectedFilename, 2, 5},
		{token.ASSIGN, "=", expectedFilename, 2, 9},
		{token.INT, "10", expectedFilename, 2, 11},
		{token.TERMINATOR, ";", expectedFilename, 2, 13},

		{token.LET, "let", expectedFilename, 4, 1},
		{token.IDENTIFIER, "add", expectedFilename, 4, 5},
		{token.ASSIGN, "=", expectedFilename, 4, 9},
		{token.FUNCTION, "fn", expectedFilename, 4, 11},
		{token.LPAREN, "(", expectedFilename, 4, 13},
		{token.IDENTIFIER, "x", expectedFilename, 4, 14},
		{token.COMMA, ",", expectedFilename, 4, 15},
		{token.IDENTIFIER, "y", expectedFilename, 4, 17},
		{token.RPAREN, ")", expectedFilename, 4, 18},
		{token.LBRACE, "{", expectedFilename, 4, 20},
		{token.IDENTIFIER, "x", expectedFilename, 5, 3},
		{token.PLUS, "+", expectedFilename, 5, 5},
		{token.IDENTIFIER, "y", expectedFilename, 5, 7},
		{token.TERMINATOR, ";", expectedFilename, 5, 8},
		{token.RBRACE, "}", expectedFilename, 6, 1},
		{token.TERMINATOR, ";", expectedFilename, 6, 2},

		{token.LET, "let", expectedFilename, 8, 1},
		{token.IDENTIFIER, "lessThanOrEqualTo", expectedFilename, 8, 5},
		{token.ASSIGN, "=", expectedFilename, 8, 23},
		{token.FUNCTION, "fn", expectedFilename, 8, 25},
		{token.LPAREN, "(", expectedFilename, 8, 27},
		{token.IDENTIFIER, "x", expectedFilename, 8, 28},
		{token.COMMA, ",", expectedFilename, 8, 29},
		{token.IDENTIFIER, "y", expectedFilename, 8, 31},
		{token.RPAREN, ")", expectedFilename, 8, 32},
		{token.LBRACE, "{", expectedFilename, 8, 34},
		{token.IDENTIFIER, "x", expectedFilename, 9, 3},
		{token.LTE, "<=", expectedFilename, 9, 5},
		{token.IDENTIFIER, "y", expectedFilename, 9, 8},
		{token.TERMINATOR, ";", expectedFilename, 9, 9},
		{token.RBRACE, "}", expectedFilename, 10, 1},
		{token.TERMINATOR, ";", expectedFilename, 10, 2},

		{token.LET, "let", expectedFilename, 12, 1},
		{token.IDENTIFIER, "result", expectedFilename, 12, 5},
		{token.ASSIGN, "=", expectedFilename, 12, 12},
		{token.IDENTIFIER, "add", expectedFilename, 12, 14},
		{token.LPAREN, "(", expectedFilename, 12, 17},
		{token.IDENTIFIER, "five", expectedFilename, 12, 18},
		{token.COMMA, ",", expectedFilename, 12, 22},
		{token.IDENTIFIER, "ten", expectedFilename, 12, 24},
		{token.RPAREN, ")", expectedFilename, 12, 27},
		{token.TERMINATOR, ";", expectedFilename, 12, 28},

		{token.LET, "let", expectedFilename, 14, 1},
		{token.IDENTIFIER, "🤖", expectedFilename, 14, 5},
		{token.ASSIGN, "=", expectedFilename, 14, 7},
		{token.FLOAT, "3.14", expectedFilename, 14, 9},
		{token.TERMINATOR, ";", expectedFilename, 14, 13},

		{token.BANG, "!", expectedFilename, 16, 1},
		{token.MINUS, "-", expectedFilename, 16, 2},
		{token.SLASH, "/", expectedFilename, 16, 3},
		{token.ASTERISK, "*", expectedFilename, 16, 4},
		{token.INT, "5", expectedFilename, 16, 5},
		{token.TERMINATOR, ";", expectedFilename, 16, 6},

		{token.INT, "5", expectedFilename, 17, 1},
		{token.LT, "<", expectedFilename, 17, 3},
		{token.INT, "10", expectedFilename, 17, 5},
		{token.GT, ">", expectedFilename, 17, 8},
		{token.INT, "5", expectedFilename, 17, 10},
		{token.TERMINATOR, ";", expectedFilename, 17, 11},

		{token.IF, "if", expectedFilename, 19, 1},
		{token.LPAREN, "(", expectedFilename, 19, 4},
		{token.INT, "5", expectedFilename, 19, 5},
		{token.LT, "<", expectedFilename, 19, 7},
		{token.INT, "10", expectedFilename, 19, 9},
		{token.RPAREN, ")", expectedFilename, 19, 11},
		{token.LBRACE, "{", expectedFilename, 19, 13},

		{token.RETURN, "return", expectedFilename, 20, 2},
		{token.TRUE, "true", expectedFilename, 20, 9},
		{token.TERMINATOR, ";", expectedFilename, 20, 13},

		{token.RBRACE, "}", expectedFilename, 21, 1},
		{token.ELSE, "else", expectedFilename, 21, 3},
		{token.LBRACE, "{", expectedFilename, 21, 8},
		
		{token.RETURN, "return", expectedFilename, 22, 2},
		{token.FALSE, "false", expectedFilename, 22, 9},
		{token.TERMINATOR, ";", expectedFilename, 22, 14},
		
		{token.RBRACE, "}", expectedFilename, 23, 1},

		{token.INT, "10", expectedFilename, 25, 1},
		{token.EQ, "==", expectedFilename, 25, 4},
		{token.INT, "10", expectedFilename, 25, 7},
		{token.TERMINATOR, ";", expectedFilename, 25, 9},

		{token.INT, "10", expectedFilename, 26, 1},
		{token.NOT_EQ, "!=", expectedFilename, 26, 4},
		{token.INT, "9", expectedFilename, 26, 7},
		{token.TERMINATOR, ";", expectedFilename, 26, 8},

		{token.EOF, "", expectedFilename, 27, 1},
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
