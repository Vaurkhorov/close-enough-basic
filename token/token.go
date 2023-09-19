package token

const (
	EOFToken = iota
	CRLF

	Number

	// Identifiers
	Variable
	Function

	// Binary Operators
	Plus
	Minus
	Divide
	Multiply
	Modulo
	Assignment
)

type Token struct {
	Type  int
	Value string
}
