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

	ast, _ := parser.Parse(tokens, 0)
	ctx := map[string]any{}
	walker.Initialize()
	value := walker.EvaluateValue(ast, ctx)
	fmt.Println(ast.Pretty())
	fmt.Println(value)
}
