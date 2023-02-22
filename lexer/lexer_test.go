package lexer

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_getIntegerToken(t *testing.T) {
	tests := []struct {
		source         string
		cursor         int
		expectedValue  string
		expectedCursor int
	}{
		{
			"foo 123",
			4,
			"123",
			7,
		},
		{
			"foo 12 3",
			4,
			"12",
			6,
		},
		{
			"foo 12a 3",
			4,
			"12",
			6,
		},
	}
	for i, test := range tests {
		cursor, token := getIntegerToken([]rune(test.source), test.cursor)
		if cursor != test.expectedCursor {
			t.Errorf("Expected Number(%d) is not same as"+
				" actual number (%d)", cursor, test.expectedCursor)
		}

		if token.Value != test.expectedValue {
			t.Errorf("Expected value(%s) is not same as"+
				" actual number (%s)", token.Value, test.expectedValue)
		}

		if token.Kind != IntegerToken {
			t.Errorf("Expected interger token in test number: %d", i)
		}
	}
}

func Test_getIdentifierToken(t *testing.T) {
	tests := []struct {
		source         string
		cursor         int
		expectedValue  string
		expectedCursor int
	}{
		{
			"123 ab + ",
			4,
			"ab",
			6,
		},
		{
			"123 ab123 + ",
			4,
			"ab123",
			9,
		},
	}
	for i, test := range tests {
		cursor, token := getIdentifierToken([]rune(test.source), test.cursor)
		if cursor != test.expectedCursor {
			t.Errorf("Expected Number(%d) is not same as"+
				" actual number (%d)", cursor, test.expectedCursor)
		}

		if token.Value != test.expectedValue {
			t.Errorf("Expected value(%s) is not same as"+
				" actual number (%s)", token.Value, test.expectedValue)
		}

		if token.Kind != IdentifierToken {
			t.Errorf("Expected interger token in test number: %d", i)
		}
	}
}

func Test_Lexer(t *testing.T) {
	tests := []struct {
		source string
		tokens []Token
	}{
		{
			" ( add 13 2 )",
			[]Token{
				{
					Value:    "(",
					Kind:     SyntaxToken,
					Location: 1,
				},
				{
					Value:    "add",
					Kind:     IdentifierToken,
					Location: 3,
				},
				{
					Value:    "13",
					Kind:     IntegerToken,
					Location: 7,
				},

				{
					Value:    "2",
					Kind:     IntegerToken,
					Location: 10,
				},
				{
					Value:    ")",
					Kind:     SyntaxToken,
					Location: 12,
				},
			},
		},
	}

	for _, test := range tests {
		tokens := Lex(test.source)

		if !reflect.DeepEqual(test.tokens, tokens) {
			fmt.Printf("expected: %#v \n", test.tokens)
			fmt.Printf("got: %#v \n", tokens)
			t.Errorf("Something is wrong with the lexer")
		}
	}
}
