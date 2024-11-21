package token

type Token struct {
	value       string
	fileName    string
	column, row uint32
	next        *Token
}
