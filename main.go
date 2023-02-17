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
	tokens := lexer.Lexer(string(app))
	fmt.Println(tokens)
}
