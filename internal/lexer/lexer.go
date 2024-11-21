package lexer

// The Lexer struct
type Lexer struct {
	source, fileName                            string
	currentPosition, currentLine, currentColumn int
}

// create lexer and return
func createLexer(source, fileName string) *Lexer {
	var lexer *Lexer = new(Lexer)
	lexer.source = source
	lexer.fileName = fileName
	lexer.currentPosition = 0
	lexer.currentLine = 1
	lexer.currentColumn = 0

	return lexer
}
