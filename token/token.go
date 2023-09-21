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

var ConstantNames = map[int]string{
	EOFToken:   "EOF",
	CRLF:       "CRLF",
	Number:     "Num",
	Variable:   "Var",
	Function:   "Fn",
	Plus:       "Plus",
	Minus:      "Minus",
	Divide:     "Div",
	Multiply:   "Mult",
	Modulo:     "Mod",
	Assignment: "Asgn",
}

type Token struct {
	Type  int
	Value string
}
