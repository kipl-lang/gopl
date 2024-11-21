package token

// The Token struct
type Token struct {
	tokenType       TokenType
	value, fileName string
	column, row     uint32
	next            *Token
}

// Create a token and return token
func makeToken(tokenType TokenType, value, fileName string, column, row uint32) *Token {
	var token *Token = new(Token)
	token.tokenType = tokenType
	token.value = value
	token.fileName = fileName
	token.column = column
	token.row = row
	token.next = nil

	return token
}
