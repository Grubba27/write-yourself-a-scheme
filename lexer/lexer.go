package lexer

import (
	"fmt"
	"os"
	"unicode"
)

type tokenKind uint

type LexContext struct {
	Source         []rune
	SourceFileName string
}

const (
	// LPAREN RPAREN
	SyntaxToken tokenKind = iota
	// numbers
	IntegerToken
	// + , fn
	IdentifierToken
)

type Token struct {
	Value    string
	Kind     tokenKind
	Location int
	Lc       LexContext
}

func (t Token) Debug(description string) {
	// 1. Grab the entire line from the source code where the token is at
	// 2. Print the entire line
	// 3. Print a marker to the column where the token is at
	// 4. Print the error/debug description

	var tokenLine []rune
	var tokenLineNumber int
	var tokenColumn int
	var inTokenLine bool
	var i int

	for i < len(t.Lc.Source) {
		r := t.Lc.Source[i]

		tokenLineNumber++
		if i < t.Location {
			tokenColumn++
		}

		tokenLine = append(tokenLine, r)

		if r == '\n' {
			// Got to the end of the line that the token is in.
			if inTokenLine {
				// Now outside the loop, `tokenLine`
				// will contain the entire source code
				// line where the token was. And
				// `tokenColumn` will be the column
				// number of the token.
				break
			}

			tokenColumn = 1
			tokenLine = nil
		}

		if i == t.Location {
			inTokenLine = true
		}

		i++
	}

	fmt.Printf("%s [at line %d, column %d in file %s]\n",
		description, tokenLineNumber, tokenColumn, t.Lc.SourceFileName)
	fmt.Println(string(tokenLine))

	// WILL NOT IF THERE ARE TABS OR OTHER WEIRD CHARACTERS
	for tokenColumn >= 1 {
		fmt.Printf(" ")
		tokenColumn--
	}
	fmt.Println("^ near here")
}

func (lc LexContext) getIdentifierToken(cursor int) (int, *Token) {
	originalCursor := cursor
	var acc []rune
	for cursor < len(lc.Source) {
		r := lc.Source[cursor]
		if unicode.IsSpace(r) || r == ')' {
			break
		}
		cursor++
		acc = append(acc, r)
	}
	if len(acc) == 0 {
		return originalCursor, nil
	}

	return cursor, &Token{
		Value:    string(acc),
		Kind:     IdentifierToken,
		Location: originalCursor,
	}

}

func (lc LexContext) getIntegerToken(cursor int) (int, *Token) {
	originalCursor := cursor
	var acc []rune
	for cursor < len(lc.Source) {
		r := lc.Source[cursor]
		if r >= '0' && r <= '9' {
			cursor++
			acc = append(acc, r)
			continue
		}
		break
	}
	if len(acc) == 0 {
		return originalCursor, nil
	}

	return cursor, &Token{
		Value:    string(acc),
		Kind:     IntegerToken,
		Location: originalCursor,
		Lc:       lc,
	}

}
func (lc LexContext) getSyntaxToken(cursor int) (int, *Token) {
	if lc.Source[cursor] == '(' || lc.Source[cursor] == ')' {
		return cursor + 1, &Token{
			Value:    string([]rune{lc.Source[cursor]}),
			Kind:     SyntaxToken,
			Location: cursor,
			Lc:       lc,
		}
	}
	return cursor, nil
}

func eatWhitespace(source []rune, cursor int) int {
	for cursor < len(source) {
		if !unicode.IsSpace(source[cursor]) {
			break
		}
		cursor++
	}
	return cursor
}

func (lc LexContext) Lex() []Token {
	var tokens []Token
	var t *Token
	cursor := 0
	for cursor < len(lc.Source) {
		cursor = eatWhitespace(lc.Source, cursor)
		if cursor == len(lc.Source) {
			break
		}

		cursor, t = lc.getSyntaxToken(cursor)
		if t != nil {
			tokens = append(tokens, *t)
			continue
		}

		cursor, t = lc.getIntegerToken(cursor)
		if t != nil {
			tokens = append(tokens, *t)
			continue
		}

		cursor, t = lc.getIdentifierToken(cursor)
		if t != nil {
			tokens = append(tokens, *t)
			continue
		}
		// err
		panic("There is some err")
	}
	return tokens
}

func New(file string) LexContext {
	program, err := os.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}

	return LexContext{
		SourceFileName: file,
		Source:         []rune(string(program)),
	}
}
