package lexer

import (
	"testing"

	"../token"
)

func TestNexToken(t *testing.T) {
	input := `=+(){},;`

	tests := []struct {
		expedtedType    token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expedtedType {
			t.Fatalf("tests[%d] - tokentype wrong.expected=%q, got=%q", i, tt.expedtedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong.expected=%q, got=%q", i, tt.expectedLiteral, tok.Literal)
		}
	}
}
