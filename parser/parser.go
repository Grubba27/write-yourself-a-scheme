package parser

import l "write-yourself-a-scheme/lexer"

type ast struct {
	car *ast // first
	cdr *ast // ...rest
}

func parse(tokens []l.Token) ast {

}
