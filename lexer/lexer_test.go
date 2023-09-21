package lexer_test

import (
	"os"
	"testing"

	"github.com/vaurkhorov/close-enough-basic/lexer"
)

func TestLex(t *testing.T) {
	lexed := lexer.Lex("../test_files/lexer_test.BAS")

	check_bytes, check_read_err := os.ReadFile("../test_files/lexer_test_result.txt")
	check_str := string(check_bytes)

	if check_read_err != nil {
		t.Errorf("Could not read 'lexer_test_result.txt'")
	} else if lexed != check_str {
		t.Errorf("lexed output does not match check")
	}
}
