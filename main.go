package main

import (
	"fmt"
	"os"
	"write-yourself-a-scheme/lexer"
)

func main() {
	app, err := os.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}
	tokens := lexer.Lex(string(app))
	fmt.Println(tokens)
}
