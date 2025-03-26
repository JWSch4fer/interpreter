package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/JWSch4fer/interpreter/lexer"
	"github.com/JWSch4fer/interpreter/parser"
)

const PROMPT = ">>"

const ERRORSEP = "\x1b[38;5;208m" + `
====================================================================
` + "\x1b[0m"

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
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

		io.WriteString(out, program.String())
		io.WriteString(out, "\n")
	}
}

func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, ERRORSEP)
	io.WriteString(out, "Something Has Interrupted Internal Execution; Error Thrown.\n")
	io.WriteString(out, "Check Syntax...\n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
	io.WriteString(out, ERRORSEP)
}
