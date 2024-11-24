package main

import (
	"fmt"
	"gopl/internal/lexer"
	"gopl/internal/token"
)

func main() {

	var lex *lexer.Lexer = lexer.CreateLexer("+", "kerem.gopl")

	var fToken *token.Token = lex.Scanner()

	fmt.Println(fToken.Column)
}
