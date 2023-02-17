package parser

import (
	"os"
	l "write-yourself-a-scheme/lexer"
)

type valueKind uint

const (
	literalKind valueKind = iota
	listKind
)

type value struct {
	kind    valueKind
	literal *l.Token
	list    *Ast
}

func (v value) pretty() string {
	if v.kind == literalKind {
		return v.literal.Value
	}
	return v.list.pretty()
}

type Ast []value

func (ast Ast) pretty() string {
	p := "("
	for _, value := range ast {
		p += value.pretty()
		p += " "
	}
	return p + ")"
}

func parse(tokens []l.Token, index int) (Ast, int) {
	var a Ast

	token := tokens[index]
	if !(token.Kind == l.SyntaxToken && token.Value == "(") {
		panic("Should be an open paran")
	}

	index++

	for index < len(tokens) {
		token := tokens[index]
		if token.Kind == l.SyntaxToken && token.Value == "(" {
			child, next := parse(tokens, index)
			a = append(a, value{
				kind: listKind,
				list: &child,
			})
			index = next
			continue
		}

		if token.Kind == l.SyntaxToken && token.Value == ")" {
			return a, index + 1
		}

		a = append(a, value{
			kind:    literalKind,
			literal: &token,
		})
		index++
	}

	if tokens[index-1].Kind == l.SyntaxToken &&
		tokens[index-1].Value != ")" {
		os.Exit(1)
	}

	return a, index
}
