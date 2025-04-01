package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/JWSch4fer/interpreter/evaluate"
	"github.com/JWSch4fer/interpreter/lexer"
	"github.com/JWSch4fer/interpreter/object"
	"github.com/JWSch4fer/interpreter/parser"
	"github.com/JWSch4fer/interpreter/repl"
)

func main() {
	// get current user
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	// If a file path is provided as an argument, read and execute
	if len(os.Args) > 1 {
		filePath := os.Args[1]
		content, err := os.ReadFile(filePath)
		if err != nil {
			fmt.Printf("Error reading file %s: %s\n", filePath, err)
		}

		// Process file with the interpreter
		l := lexer.New(string(content))
		p := parser.New(l)
		program := p.ParseProgram()

		// Check for parser error
		if len(p.Errors()) != 0 {
			fmt.Println("parser errors:")
			for _, msg := range p.Errors() {
				fmt.Println("\t", msg)
			}
			os.Exit(1)
		}

		//Create Environment and evaluate the program
		env := object.NewEnvironment()
		result := evaluate.Eval(program, env)
		/*
			TODO: need to update this right now we just print the last thing
				that was run by the program. So if the last ast node evaluated
				was a comment we print the comment, if it was a boolean we print
				true/false, etc. need to implement a print function
		*/

		// fmt.Println(result.Type())
		fmt.Println(result.Inspect())
	} else {
		// interactive mode

		fmt.Printf("Hello %s\n", user.Username)
		fmt.Println("starting interpreter...")
		repl.Start(os.Stdin, os.Stdout)
	}
}
