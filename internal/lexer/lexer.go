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
func CreateLexer(source, fileName string) *Lexer {
	var lexer *Lexer = new(Lexer)
	lexer.source = source
	lexer.fileName = fileName
	lexer.currentPosition = 0
	lexer.currentLine = 1
	lexer.currentColumn = 0

	return lexer
}

func (lexer *Lexer) Scanner() *token.Token {
	var firstToken *token.Token = nil

	for {
		var tmpToken *token.Token = lexer.scanToken()

		if firstToken == nil {
			firstToken = tmpToken
		} else {
			token.GetLastToken(firstToken).Next = tmpToken
		}

		if token.GetLastToken(firstToken).TokenType == token.TOKEN_EOF {
			break
		}
	}

	return firstToken
}

func (lexer *Lexer) scanToken() *token.Token {

	lexer.skipWhiteSpace()

	if lexer.isAtEnd() {
		return token.MakeToken(token.TOKEN_EOF, "end", lexer.fileName, lexer.currentLine, lexer.currentColumn)
	}

	var char byte = lexer.advance()

	switch char {
	case '+':
		if lexer.isMatch('=') {
			return token.MakeToken(token.TOKEN_PLUS_EQUAL, "+=", lexer.fileName, lexer.currentLine, lexer.currentColumn)
		}
		return token.MakeToken(token.TOKEN_PLUS, "+", lexer.fileName, lexer.currentLine, lexer.currentColumn)
	default:
		return token.MakeToken(token.TOKEN_ERROR, "Unexpexted Token", lexer.fileName, lexer.currentLine, lexer.currentColumn)
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
	lexer.currentColumn++
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

func (lexer *Lexer) skipWhiteSpace() {
	for !lexer.isAtEnd() {
		var char byte = lexer.peek()

		switch char {
		case ' ':
			fallthrough
		case '\t':
			fallthrough
		case '\r':
			lexer.advance()
		case '\n':
			lexer.currentLine++
			lexer.currentColumn = 0
			lexer.advance()
		default:
			return
		}
	}
}
