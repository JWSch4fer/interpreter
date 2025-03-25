package ast

import (
	"testing"

	"github.com/JWSch4fer/interpreter/token"
)

func TestString(t *testing.T) {

	//construct abstract tree by hand for testing
	Program := &Program{
		Statements: []Statement{
			&LetStatement{
				Token: token.Token{Type: token.LET, Literal: "let"},
				Name: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "myVar"},
					Value: "myVar",
				},
				Value: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "anotherVar"},
					Value: "anotherVar",
				},
			},
		},
	}

	if Program.String() != "let myVar = anotherVar;" {
		t.Errorf("program.String() is wrong. got=%q", Program.String())
	}
}
