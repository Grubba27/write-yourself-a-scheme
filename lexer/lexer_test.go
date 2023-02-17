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

		if token.value != test.expectedValue {
			t.Errorf("Expected value(%s) is not same as"+
				" actual number (%s)", token.value, test.expectedValue)
		}

		if token.kind != integerToken {
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

		if token.value != test.expectedValue {
			t.Errorf("Expected value(%s) is not same as"+
				" actual number (%s)", token.value, test.expectedValue)
		}

		if token.kind != identifierToken {
			t.Errorf("Expected interger token in test number: %d", i)
		}
	}
}

func Test_Lexer(t *testing.T) {
	tests := []struct {
		source string
		tokens []token
	}{
		{
			" ( add 13 2 )",
			[]token{
				{
					value:    "(",
					kind:     syntaxToken,
					location: 1,
				},
				{
					value:    "add",
					kind:     identifierToken,
					location: 3,
				},
				{
					value:    "13",
					kind:     integerToken,
					location: 7,
				},

				{
					value:    "2",
					kind:     integerToken,
					location: 10,
				},
				{
					value:    ")",
					kind:     syntaxToken,
					location: 12,
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
