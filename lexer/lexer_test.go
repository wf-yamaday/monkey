package lexer

import (
	"testing"

	"github.com/wf-yamaday/monkey/token"
)

func TestNextToken(t *testing.T) {
	input := `=+(){},;`

	tests := []struct {
		expextedType    token.TokenType
		expextedLiteral string
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

		if tok.Type != tt.expextedType {
			t.Fatalf("test[%d] - tokentype wrong. expected=%q, got=%q", i, tt.expextedType, tok.Type)
		}

		if tok.Literal != tt.expextedLiteral {
			t.Fatalf("test[%d - literal wrong. expected=%q, got=%q", i, tt.expextedLiteral, tok.Type)
		}
	}
}
