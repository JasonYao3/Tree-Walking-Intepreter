package ast

import (
	"go_interpreter/token"
	"testing"
)

func TestString(t *testing.T) {
	program := &Program{
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

	if program.String() != "let myVar = anotherVar;" {
		t.Errorf("program.String() wrong. got=%q", program.String())
	}
}

func TestLetStatementString(t *testing.T) {
	letStatement := &LetStatement{
		Token: token.Token{Type: token.LET, Literal: "let"},
		Name: &Identifier{
			Token: token.Token{Type: token.IDENT, Literal: "myVar"},
			Value: "myVar",
		},
		Value: &Identifier{
			Token: token.Token{Type: token.IDENT, Literal: "anotherVar"},
			Value: "anotherVar",
		},
	}

	if letStatement.String() != "let myVar = anotherVar;" {
		t.Errorf("program.String() wrong. got=%q", letStatement.String())
	}
}

func TestReturnStatementString(t *testing.T) {
	returnStatement := &ReturnStatement{
		Token: token.Token{Type: token.LET, Literal: "return"},
		ReturnValue: &Identifier{
			Token: token.Token{Type: token.IDENT, Literal: "64"},
			Value: "64",
		},
	}

	if returnStatement.String() != "return 64;" {
		t.Errorf("program.String() wrong. got=%q", returnStatement.String())
	}
}
