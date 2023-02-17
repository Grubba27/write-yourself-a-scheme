package lexer

import "unicode"

type tokenKind uint

const (
	// LPAREN RPAREN
	syntaxToken tokenKind = iota
	// numbers
	integerToken
	// + , fn
	identifierToken
)

type token struct {
	value    string
	kind     tokenKind
	location int
}

func (t token) debug(source []rune) {
	// debug
}

func getIdentifierToken(source []rune, cursor int) (int, *token) {
	originalCursor := cursor
	var acc []rune
	for cursor < len(source) {
		r := source[cursor]
		if !unicode.IsSpace(r) {
			cursor++
			acc = append(acc, r)
			continue
		}
		break

	}
	if len(acc) == 0 {
		return originalCursor, nil
	}

	return cursor, &token{
		value:    string(acc),
		kind:     identifierToken,
		location: originalCursor,
	}

}

func getIntegerToken(source []rune, cursor int) (int, *token) {
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

	return cursor, &token{
		value:    string(acc),
		kind:     integerToken,
		location: originalCursor,
	}

}
func getSyntaxToken(source []rune, cursor int) (int, *token) {
	if source[cursor] == '(' || source[cursor] == ')' {
		return cursor + 1, &token{
			value:    string([]rune{source[cursor]}),
			kind:     syntaxToken,
			location: cursor,
		}
	}
	return cursor, nil
}

func eatWhitespace(source []rune, cursor int) int {
	for cursor < len(source) {
		if unicode.IsSpace(source[cursor]) {
			cursor++
			continue
		}
		break
	}
	return cursor
}

func Lexer(raw string) []token {
	source := []rune(raw)
	var tokens []token
	var t *token
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
