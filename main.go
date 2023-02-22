package main

import (
	"fmt"
	"os"
	l "write-yourself-a-scheme/lexer"
	"write-yourself-a-scheme/parser"
	"write-yourself-a-scheme/walker"
)

func main() {
	app, err := os.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}
	tokens := l.Lex(string(app))

	var parserIndex int
	ast := parser.Ast{
		parser.Value{
			Kind: parser.LiteralKind,
			Literal: &l.Token{
				Value: "begin",
				Kind:  l.IdentifierToken,
			},
		},
	}

	for parserIndex < len(tokens) {
		childAst, next := parser.Parse(tokens, parserIndex)
		ast = append(ast, parser.Value{
			Kind: parser.ListKind,
			List: &childAst,
		})
		parserIndex = next
	}
	ctx := map[string]any{}
	walker.Initialize()
	value := walker.EvaluateValue(ast, ctx)
	fmt.Println(ast.Pretty())
	fmt.Println(value)
}
