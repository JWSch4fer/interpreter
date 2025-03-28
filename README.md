# Interpreter in Go

This repository implements a simple interpreter for a custom programming language written in Go. It includes a lexer, parser, evaluator, and a REPL (Read-Eval-Print Loop) for interactive usage. The language supports basic arithmetic, boolean logic, conditionals, functions (with closures and recursion), and more.

## Features
- Lexer: Tokenizes input strings and supports multi-character tokens (such as == and !=).

- Parser: Uses recursive descent (Pratt parsing) to construct an Abstract Syntax Tree (AST) from tokens.

- Evaluator: Executes the AST, supporting arithmetic, boolean operations, conditionals, function definitions, and function calls.

- REPL: Interactive shell for testing code snippets.

- Error Handling: Detects and reports both syntax and runtime errors.

- Testing: Comprehensive test suite covering lexer, parser, evaluator, and AST construction.

## Installation
Ensure you have Go installed (version 1.24 or above).

## Usage
To run the interpreter in interactive mode:
You will see a prompt (>>) where you can enter code. For example:

```
go run main/main.go

starting interpreter...
>>let counter = df(x) { if (x > 100) {return x} else {counter(x + 1)}}
>>counter(0)
101

>> let add = df(x, y) { x + y; };
>> add(5, 3);
8

>>exit
```

## Running Tests
A comprehensive test suite is provided to validate the lexer, parser, evaluator, and AST implementation.

## Future Improvements
- Unicode & UTF-8 Support: Currently, the lexer processes ASCII input. Consider switching from byte to rune for full UTF-8 support.

- Extended Data Types: Add support for floats, strings, and additional data types.

- Enhanced Error Messages: Improve debugging capabilities with more detailed error reporting.

## Note:
This project follows the excellent guide written by Thorsten Ball, Writing an Interpreter in Go. Highly recommend this resource to anyone interested in a deeper understanding of how languages are built.

## License
This project is licensed under the MIT License.
Feel free to adjust the content to match your project's specifics and style preferences.
