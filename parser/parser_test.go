package parser

import (
	"fmt"
	"testing"
	"write-yourself-a-scheme/lexer"
)

func areEqual(expected, got Value) bool {
	if expected.Kind != got.Kind {
		println("Kinds are not equal", expected.Kind, got.Kind)
		return false
	}

	if expected.Kind == LiteralKind {
		if expected.Literal.Value != got.Literal.Value {
			fmt.Println("Literals not equal", expected.Literal, got.Literal)
			return false
		}

		return true
	}

	// is a List so recurse :)
	return compareAst(*expected.List, *got.List)
}
func compareAst(expected, got Ast) bool {
	if len(expected) != len(got) {
		return false

	}
	for i := range expected {
		e := expected[i]
		g := got[i]

		if !areEqual(e, g) {
			return false
		}
	}
	return true
}

func Test_parse(t *testing.T) {
	tests := []struct {
		input        string
		prettyOutput string
		output       Ast
	}{
		{
			"(+ 1 2)",
			"(+ 1 2 )",
			Ast{
				Value{
					Kind:    LiteralKind,
					Literal: &lexer.Token{Value: "+"},
				},
				Value{
					Kind:    LiteralKind,
					Literal: &lexer.Token{Value: "1"},
				},
				Value{
					Kind:    LiteralKind,
					Literal: &lexer.Token{Value: "2"},
				},
			},
		},
		{
			"(+ 1 (- 12 9))",
			"(+ 1 (- 12 9 ) )",
			Ast{
				Value{
					Kind:    LiteralKind,
					Literal: &lexer.Token{Value: "+"},
				},
				Value{
					Kind:    LiteralKind,
					Literal: &lexer.Token{Value: "1"},
				},
				Value{
					Kind: ListKind,
					List: &Ast{
						Value{
							Kind:    LiteralKind,
							Literal: &lexer.Token{Value: "-"},
						},
						Value{
							Kind:    LiteralKind,
							Literal: &lexer.Token{Value: "12"},
						},
						Value{
							Kind:    LiteralKind,
							Literal: &lexer.Token{Value: "9"},
						},
					},
				},
			},
		},
		{
			"(+ 1 (- 12 9) 12)",
			"(+ 1 (- 12 9 ) 12 )",
			Ast{
				Value{
					Kind:    LiteralKind,
					Literal: &lexer.Token{Value: "+"},
				},
				Value{
					Kind:    LiteralKind,
					Literal: &lexer.Token{Value: "1"},
				},
				Value{
					Kind: ListKind,
					List: &Ast{
						Value{
							Kind:    LiteralKind,
							Literal: &lexer.Token{Value: "-"},
						},
						Value{
							Kind:    LiteralKind,
							Literal: &lexer.Token{Value: "12"},
						},
						Value{
							Kind:    LiteralKind,
							Literal: &lexer.Token{Value: "9"},
						},
					},
				},
				Value{
					Kind:    LiteralKind,
					Literal: &lexer.Token{Value: "12"},
				},
			},
		},
		{
			"((+ 1 2) 1 (- 12 9) 12)",
			"((+ 1 2 ) 1 (- 12 9 ) 12 )",
			Ast{
				Value{
					Kind: ListKind,
					List: &Ast{
						Value{
							Kind:    LiteralKind,
							Literal: &lexer.Token{Value: "+"},
						},
						Value{
							Kind:    LiteralKind,
							Literal: &lexer.Token{Value: "1"},
						},
						Value{
							Kind:    LiteralKind,
							Literal: &lexer.Token{Value: "2"},
						},
					},
				},
				Value{
					Kind:    LiteralKind,
					Literal: &lexer.Token{Value: "1"},
				},
				Value{
					Kind: ListKind,
					List: &Ast{
						Value{
							Kind:    LiteralKind,
							Literal: &lexer.Token{Value: "-"},
						},
						Value{
							Kind:    LiteralKind,
							Literal: &lexer.Token{Value: "12"},
						},
						Value{
							Kind:    LiteralKind,
							Literal: &lexer.Token{Value: "9"},
						},
					},
				},
				Value{
					Kind:    LiteralKind,
					Literal: &lexer.Token{Value: "12"},
				},
			},
		},
	}

	for _, test := range tests {
		tokens := lexer.Lex(test.input)
		ast, _ := Parse(tokens, 0)
		if ast.Pretty() != test.prettyOutput {
			fmt.Printf("expected: %s \n", test.prettyOutput)
			fmt.Printf("got: %s \n", ast.Pretty())
			t.Errorf("Something is wrong with the parser")
		}
		if !compareAst(test.output, ast) {
			fmt.Printf("expected: %s \n", test.prettyOutput)
			fmt.Printf("got: %s \n", ast.Pretty())
			t.Errorf("Something is wrong with the parser")
		}
	}
}
