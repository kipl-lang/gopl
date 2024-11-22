package lexer

import (
	"gopl/internal/token"
)

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

func (lexer *Lexer) scanner() {
	var firstToken *token.Token = nil

	for()
}

func (lexer *Lexer) isAtEnd() bool {
	return lexer.currentPosition == len(lexer.source)
}
