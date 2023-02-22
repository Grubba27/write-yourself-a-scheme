package parser

import (
	"os"
	l "write-yourself-a-scheme/lexer"
)

type valueKind uint

const (
	LiteralKind valueKind = iota
	ListKind
)

type Value struct {
	Kind    valueKind
	Literal *l.Token
	List    *Ast
}

func (v Value) pretty() string {
	if v.Kind == LiteralKind {
		return v.Literal.Value
	}
	return v.List.Pretty()
}

type Ast []Value

func (ast Ast) Pretty() string {
	p := "("
	for _, value := range ast {
		p += value.pretty()
		p += " "
	}
	return p + ")"
}

// Parse for example: "(+ 13 (- 12 1)"
// Parse(["(", "+", "13", "(", "-", "12", "1", ")", ")"]):
//
//	should produce: ast{
//	  value{
//	    kind: Literal,
//	    Literal: "+",
//	  },
//	  value{
//	    kind: Literal,
//	    Literal: "13",
//	  },
//	  value{
//	    kind: List,
//	    List: ast {
//	      value {
//	        kind: Literal,
//	        Literal: "-",
//	      },
//	      value {
//	        kind: Literal,
//	        Literal: "12",
//	      },
//	      value {
//	        kind: Literal,
//	        Literal: "1",
//	      },
//	    }
//	  }
//	}
func Parse(tokens []l.Token, index int) (Ast, int) {
	var a Ast

	token := tokens[index]
	if !(token.Kind == l.SyntaxToken && token.Value == "(") {
		panic("Should be an open paran")
	}

	index++

	for index < len(tokens) {
		token := tokens[index]
		if token.Kind == l.SyntaxToken && token.Value == "(" {
			child, next := Parse(tokens, index)
			a = append(a, Value{
				Kind: ListKind,
				List: &child,
			})
			index = next
			continue
		}

		if token.Kind == l.SyntaxToken && token.Value == ")" {
			return a, index + 1
		}

		a = append(a, Value{
			Kind:    LiteralKind,
			Literal: &token,
		})
		index++
	}

	if tokens[index-1].Kind == l.SyntaxToken &&
		tokens[index-1].Value != ")" {
		os.Exit(1)
	}

	return a, index
}
