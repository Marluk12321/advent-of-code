package data

type TokenType int

const (
	OPEN_BRACKET TokenType = iota
	CLOSED_BRACKET
	INT
	COMMA
)

type Token struct {
	tokenType TokenType
	text      string
}

type IntBuffer []rune

func (buffer *IntBuffer) flush() Token {
	token := Token{
		tokenType: INT,
		text:      string(*buffer),
	}
	(*buffer) = (*buffer)[:0]
	return token
}

var TOKEN_TYPE = map[rune]TokenType{
	'[': OPEN_BRACKET,
	']': CLOSED_BRACKET,
	',': COMMA,
}

func tokenize(text string) []Token {
	var tokens []Token
	var intBuffer IntBuffer
	for _, char := range text {
		if char >= '0' && char <= '9' {
			intBuffer = append(intBuffer, char)
			continue
		}
		if len(intBuffer) > 0 {
			tokens = append(tokens, intBuffer.flush())
		}
		tokenType, exists := TOKEN_TYPE[char]
		if !exists {
			panic(char)
		}
		tokens = append(tokens, Token{
			tokenType: tokenType,
			text:      string(char),
		})
	}
	return tokens
}
