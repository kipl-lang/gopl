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
		var tmpToken *token.Token = lexer.scanToken()

		if firstToken == nil {
			firstToken = tmpToken
		} else {
			token.GetLastToken(firstToken).Next = tmpToken
		}

		if token.GetLastToken(firstToken).TokenType != token.TOKEN_EOF {
			break
		}
	}
}

func (lexer *Lexer) scanToken() *token.Token {

	if lexer.isAtEnd() {
		return token.MakeToken(token.TOKEN_EOF, "end", lexer.fileName, lexer.currentLine, lexer.currentColumn)
	}

	var char byte = lexer.advance()

	switch char {
	case '+':

	}
}

func (lexer *Lexer) peek() byte {
	return lexer.source[lexer.currentPosition]
}

func (lexer *Lexer) nextPeek() byte {
	return lexer.source[lexer.currentPosition+1]
}

func (lexer *Lexer) advance() byte {
	var char byte = lexer.source[lexer.currentPosition]
	lexer.currentPosition++
	return char
}

func (lexer *Lexer) isAtEnd() bool {
	return lexer.currentPosition == uint32(len(lexer.source))
}

func (lexer *Lexer) isMatch(char byte) bool {
	if lexer.isAtEnd() || lexer.peek() != char {
		return false
	}

	lexer.advance()
	return true
}
