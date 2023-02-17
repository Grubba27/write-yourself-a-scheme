package main

import (
	"os"
	l "write-yourself-a-scheme/lexer"
	"write-yourself-a-scheme/parser"
)

func main() {
	app, err := os.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}
	tokens := l.Lex(string(app))

	ast, _ := parser.Parse(tokens, 0)
	println(ast.Pretty())
}
