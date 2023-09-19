package lexer

import (
	"testing"
)

func TestLex(t *testing.T) {
	// !TODO add a proper test and its file
	lexed := lex("test.BAS")

	if lexed == "" {
		t.Errorf("failed")
	}
}
