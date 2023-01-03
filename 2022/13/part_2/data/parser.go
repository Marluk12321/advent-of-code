package data

import "strconv"

func parseInt(token Token) int {
	value, err := strconv.Atoi(token.text)
	if err != nil {
		panic(err)
	}
	return value
}

var (
	VALUE_OR_END = map[TokenType]bool{
		OPEN_BRACKET:   true,
		INT:            true,
		CLOSED_BRACKET: true,
	}
	COMMA_OR_END = map[TokenType]bool{
		COMMA:          true,
		CLOSED_BRACKET: true,
	}
)

func parseList(tokens []Token) ([]interface{}, []Token) {
	var values []interface{}
	expectedTypes := VALUE_OR_END
	tokenQueue := tokens[1:]
	i := 0
	for {
		token := tokenQueue[i]
		if !expectedTypes[token.tokenType] {
			panic(tokens)
		}
		switch token.tokenType {
		case CLOSED_BRACKET:
			return values, tokenQueue[i+1:]
		case OPEN_BRACKET:
			subvalues, leftovers := parseList(tokenQueue[i:])
			values = append(values, subvalues)
			expectedTypes = COMMA_OR_END
			tokenQueue = leftovers
			i = 0
			continue
		case INT:
			values = append(values, parseInt(token))
			expectedTypes = COMMA_OR_END
		case COMMA:
			expectedTypes = VALUE_OR_END
		}
		i++
	}
}

func Parse(text string) interface{} {
	tokens := tokenize(text)
	switch tokens[0].tokenType {
	case INT:
		if len(tokens) > 1 {
			panic(text)
		}
		return parseInt(tokens[0])
	case OPEN_BRACKET:
		values, leftovers := parseList(tokens)
		if len(leftovers) > 0 {
			panic(text)
		}
		return values
	default:
		panic(text)
	}
}
