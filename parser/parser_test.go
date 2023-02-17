package parser

import (
	"fmt"
	"testing"
	"write-yourself-a-scheme/lexer"
)

func areEqual(expected, got value) bool {
	if expected.kind != got.kind {
		println("Kinds are not equal", expected.kind, got.kind)
		return false
	}

	if expected.kind == literalKind {
		if expected.literal.Value != got.literal.Value {
			fmt.Println("Literals not equal", expected.literal, got.literal)
			return false
		}

		return true
	}

	// is a list so recurse :)
	return compareAst(*expected.list, *got.list)
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
				value{
					kind:    literalKind,
					literal: &lexer.Token{Value: "+"},
				},
				value{
					kind:    literalKind,
					literal: &lexer.Token{Value: "1"},
				},
				value{
					kind:    literalKind,
					literal: &lexer.Token{Value: "2"},
				},
			},
		},
		{
			"(+ 1 (- 12 9))",
			"(+ 1 (- 12 9 ) )",
			Ast{
				value{
					kind:    literalKind,
					literal: &lexer.Token{Value: "+"},
				},
				value{
					kind:    literalKind,
					literal: &lexer.Token{Value: "1"},
				},
				value{
					kind: listKind,
					list: &Ast{
						value{
							kind:    literalKind,
							literal: &lexer.Token{Value: "-"},
						},
						value{
							kind:    literalKind,
							literal: &lexer.Token{Value: "12"},
						},
						value{
							kind:    literalKind,
							literal: &lexer.Token{Value: "9"},
						},
					},
				},
			},
		},
		{
			"(+ 1 (- 12 9) 12)",
			"(+ 1 (- 12 9 ) 12 )",
			Ast{
				value{
					kind:    literalKind,
					literal: &lexer.Token{Value: "+"},
				},
				value{
					kind:    literalKind,
					literal: &lexer.Token{Value: "1"},
				},
				value{
					kind: listKind,
					list: &Ast{
						value{
							kind:    literalKind,
							literal: &lexer.Token{Value: "-"},
						},
						value{
							kind:    literalKind,
							literal: &lexer.Token{Value: "12"},
						},
						value{
							kind:    literalKind,
							literal: &lexer.Token{Value: "9"},
						},
					},
				},
				value{
					kind:    literalKind,
					literal: &lexer.Token{Value: "12"},
				},
			},
		},
		{
			"((+ 1 2) 1 (- 12 9) 12)",
			"((+ 1 2 ) 1 (- 12 9 ) 12 )",
			Ast{
				value{
					kind: listKind,
					list: &Ast{
						value{
							kind:    literalKind,
							literal: &lexer.Token{Value: "+"},
						},
						value{
							kind:    literalKind,
							literal: &lexer.Token{Value: "1"},
						},
						value{
							kind:    literalKind,
							literal: &lexer.Token{Value: "2"},
						},
					},
				},
				value{
					kind:    literalKind,
					literal: &lexer.Token{Value: "1"},
				},
				value{
					kind: listKind,
					list: &Ast{
						value{
							kind:    literalKind,
							literal: &lexer.Token{Value: "-"},
						},
						value{
							kind:    literalKind,
							literal: &lexer.Token{Value: "12"},
						},
						value{
							kind:    literalKind,
							literal: &lexer.Token{Value: "9"},
						},
					},
				},
				value{
					kind:    literalKind,
					literal: &lexer.Token{Value: "12"},
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
