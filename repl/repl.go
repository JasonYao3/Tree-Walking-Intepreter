package repl

import (
	"bufio"
	"fmt"
	"go_interpreter/compiler"

	// "go_interpreter/evaluator"
	// "go_interpreter/object"
	"go_interpreter/lexer"
	"go_interpreter/parser"
	"go_interpreter/vm"
	"io"
)

/*
the REPL reads input, sends it to the interpreter for evaluation, prints the result/output of the
interpreter and starts again. Read, Evaluate, Print, Loop.
*/

const PROMPT = ">> "

// TODO: Initialize the lexer with an io.Reader and the filename.
func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	// env := object.NewEnvironment()
	// macroEnv := object.NewEnvironment()

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		// evaluator.DefineMacros(program, macroEnv)
		// expanded := evaluator.ExpandedMacros(program, macroEnv)
		// evaluated := evaluator.Eval(expanded, env)
		// if evaluated != nil {
		// 	io.WriteString(out, evaluated.Inspect())
		// 	io.WriteString(out, "\n")
		// }

		comp := compiler.New()
		err := comp.Compile(program)
		if err != nil {
			fmt.Fprintf(out, "Compilation failed:\n %s\n", err)
			continue
		}

		machine := vm.New(comp.Bytecode())
		err = machine.Run()
		if err != nil {
			fmt.Fprintf(out, "Executing bytecode failed:\n %s\n", err)
			continue
		}

		for {
		lastPopped := machine.LastPoppedStackElem()
		io.WriteString(out, lastPopped.Inspect())
		io.WriteString(out, "\n")
		}

	}
}

func printParserErrors(out io.Writer, errors []string) {
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
