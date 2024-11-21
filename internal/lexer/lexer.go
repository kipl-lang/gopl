package lexer

type Lexer struct {
	source, fileName                           string
	currentPosition, currentRow, currentColumn int
}
