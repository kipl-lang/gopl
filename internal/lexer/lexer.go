package lexer

import (
	"gopl/internal/token"
)

// The Lexer struct
type Lexer struct {
	source, fileName                            string
	currentPosition, currentLine, currentColumn uint32
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

func (lexer *Lexer) Scanner() {
	var firstToken *token.Token = nil

	for {
		if token.GetLastToken(firstToken).TokenType != token.TOKEN_EOF {
			break
		}
	}
}

func (lexer *Lexer) scanToken() *token.Token {

}

func (lexer *Lexer) peek() byte {
	return lexer.source[lexer.currentPosition]
}

func (lexer *Lexer) advance() byte {
	var char byte = lexer.source[lexer.currentPosition]
	lexer.currentPosition++
	return char
}

func (lexer *Lexer) isAtEnd() bool {
	return lexer.currentPosition == uint32(len(lexer.source))
}
