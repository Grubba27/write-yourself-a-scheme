package parser

import l "write-yourself-a-scheme/lexer"

type valueKind uint

const (
	literalKind valueKind = iota
	listKind
)

type value struct {
	kind    valueKind
	literal *l.Token
	list    Ast
}

type Ast []value

func parse(tokens []l.Token, index int) (Ast, int) {
	var ast Ast

	for index < len(tokens) {
		token := tokens[index]

		if token.Kind == l.SyntaxToken && token.Value == "(" {
			child, next := parse(tokens, index+1)
			ast = append(ast, value{
				kind: listKind,
				list: child,
			})
			index = next
			continue
		}

		if token.Kind == l.SyntaxToken && token.Value == ")" {
			return ast, index + 1
		}

		ast = append(ast, value{
			kind:    literalKind,
			literal: &token,
		})
		index++
	}

	return ast, index
}
