package lexer

import "unicode"

type tokenKind uint

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
}

func (t Token) debug(source []rune) {
	// debug
}

func getIdentifierToken(source []rune, cursor int) (int, *Token) {
	originalCursor := cursor
	var acc []rune
	for cursor < len(source) {
		r := source[cursor]
		if unicode.IsSpace(r) {
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

func getIntegerToken(source []rune, cursor int) (int, *Token) {
	originalCursor := cursor
	var acc []rune
	for cursor < len(source) {
		r := source[cursor]
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
	}

}
func getSyntaxToken(source []rune, cursor int) (int, *Token) {
	if source[cursor] == '(' || source[cursor] == ')' {
		return cursor + 1, &Token{
			Value:    string([]rune{source[cursor]}),
			Kind:     SyntaxToken,
			Location: cursor,
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

func Lex(raw string) []Token {
	source := []rune(raw)
	var tokens []Token
	var t *Token
	cursor := 0
	for cursor < len(source) {
		cursor = eatWhitespace(source, cursor)
		if cursor == len(source) {
			break
		}

		cursor, t = getSyntaxToken(source, cursor)
		if t != nil {
			tokens = append(tokens, *t)
			continue
		}

		cursor, t = getIntegerToken(source, cursor)
		if t != nil {
			tokens = append(tokens, *t)
			continue
		}

		cursor, t = getIdentifierToken(source, cursor)
		if t != nil {
			tokens = append(tokens, *t)
			continue
		}
		// err
		panic("There is some err")
	}
	return tokens
}
