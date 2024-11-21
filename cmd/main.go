package main

import "fmt"

type TokenType int

const (
	PLUS TokenType = iota
	MINUS
)

func main() {
	var test TokenType = PLUS

	fmt.Println(test)
}
