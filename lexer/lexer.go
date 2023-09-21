package lexer

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"

	"github.com/vaurkhorov/close-enough-basic/token"
)

func Lex(path string) string {
	// This is for testing

	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	lexer := NewLexer(file)

	lexed := ""

	for {
		pos, tok, lit := get_token(lexer)
		if tok == token.EOFToken {
			break
		}

		lexed += fmt.Sprintf("%d:%d\t%s\t%s\n", pos.row, pos.column, token.ConstantNames[tok], lit)
	}

	return lexed
}

type Position struct {
	row    int
	column int
}

type Lexer struct {
	position Position
	reader   *bufio.Reader
}

func NewLexer(reader io.Reader) *Lexer {
	return &Lexer{
		position: Position{
			row:    0,
			column: 0,
		},
		reader: bufio.NewReader(reader),
	}
}

func get_token(lexer *Lexer) (Position, int, string) {
	for {
		t, _, err := lexer.reader.ReadRune()

		if err != nil {
			if err == io.EOF {
				return lexer.position, token.EOFToken, ""
			}

			panic(err)
		}

		switch t {
		case '+':
			lexer.position.column++
			return lexer.position, token.Plus, "+"
		case '-':
			lexer.position.column++
			return lexer.position, token.Minus, "-"
		case '*':
			lexer.position.column++
			return lexer.position, token.Multiply, "*"
		case '/':
			lexer.position.column++
			return lexer.position, token.Divide, "/"
		case '%':
			lexer.position.column++
			return lexer.position, token.Modulo, "%"
		case '=':
			lexer.position.column++
			return lexer.position, token.Assignment, "="

		case ' ':
			continue

		case '\n':
			lexer.position.row++
			return lexer.position, token.CRLF, ";"

		default:
			if unicode.IsDigit(t) {
				pos := Position{
					row:    lexer.position.row,
					column: lexer.position.column + 1,
				}
				lexer.reader.UnreadRune()

				num := get_number(lexer)

				return pos, token.Number, num
			} else if unicode.IsLetter(t) || t == '_' {
				pos := Position{
					row:    lexer.position.row,
					column: lexer.position.column + 1,
				}
				lexer.reader.UnreadRune()

				tok, ret := get_identifier(lexer)

				return pos, tok, ret
			}
		}
	}
}

func get_number(lexer *Lexer) string {
	number := ""

	for {
		t, _, err := lexer.reader.ReadRune()

		if err != nil {
			if err == io.EOF {
				lexer.reader.UnreadRune()
				return number
			}

			panic(err)
		}

		if unicode.IsDigit(t) {
			number += fmt.Sprintf("%c", t)
			lexer.position.column++
		} else { // if err == strconv.ErrSyntax
			// fmt.Println("here.")
			lexer.reader.UnreadRune()
			break
		}
		// else {
		// 	panic(err)
		// }
	}
	return number
}

func get_identifier(lexer *Lexer) (int, string) {
	identifier := ""
	args := ""

	// get just the identifier, break if '(' is encountered
	for {
		t, _, err := lexer.reader.ReadRune()

		if err != nil {
			if err == io.EOF {
				lexer.reader.UnreadRune()
				return token.Variable, identifier
			}

			panic(err)
		}

		if unicode.IsLetter(t) {
			identifier += fmt.Sprintf("%c", t)
			lexer.position.column++
		} else if t == '(' {
			lexer.position.column++
			break
		} else {
			lexer.reader.UnreadRune()
			return token.Variable, identifier
		}
	}

	// '(' was encountered, so this has to be a function
	// now we put the arguments in a comma separated string
	for {
		t, _, err := lexer.reader.ReadRune()

		if err != nil {
			if err == io.EOF {
				panic("expected ')' before EOF")
			}

			panic(err)
		}

		if unicode.IsLetter(t) {
			args += fmt.Sprintf("%c", t)
			lexer.position.column++
		} else if unicode.IsDigit(t) {
			lexer.reader.UnreadRune()
			args += get_number(lexer)
		} else if t == rune(')') {
			lexer.position.column++
			if args == "" {
				return token.Function, identifier
			} else {
				return token.Function, fmt.Sprintf("%s:%s", identifier, args)
			}
		} else if t == ',' {
			args += ","
			lexer.position.column++
		} else if t == ' ' {
			lexer.position.column++
		} else {
			panic("expected ')'")
		}
	}
}
