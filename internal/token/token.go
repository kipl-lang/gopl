package token

// The Token struct
type Token struct {
	TokenType       TokenType
	Value, FileName string
	Column, Row     uint32
	Next            *Token
}

// Create a token and return token
func MakeToken(tokenType TokenType, value, fileName string, column, row uint32) *Token {
	var token *Token = new(Token)
	token.TokenType = tokenType
	token.Value = value
	token.FileName = fileName
	token.Column = column
	token.Row = row
	token.Next = nil

	return token
}

// Get last token
func GetLastToken(token *Token) *Token {
	if token.Next != nil {
		return token
	}

	return GetLastToken(token.Next)
}
