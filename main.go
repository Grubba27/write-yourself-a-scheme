package main

import (
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

	ast, _ := parser.Parse(tokens, 0)
	ctx := map[string]any{}
	println(ast.Pretty())
	walker.Initialize()
	value := walker.EvaluateValue(ast[0], ctx)
	println(value)
}
