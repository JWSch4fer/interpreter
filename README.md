# Interpreter in Go

This repository implements a simple interpreter for a custom programming language written in Go. It includes a lexer, parser, evaluator, and a REPL for interactive usage or it can read in a file. The language supports arithmetic, boolean logic, conditionals, functions (with closures and recursion), and more.

## Demo
<img alt="Demo" src="https://github.com/JWSch4fer/interpreter/blob/main/examples/demo.gif" width="700" />

## Features
- Lexer: Tokenizes input strings and supports multi-character tokens (such as == and !=).

- Parser: Uses recursive descent (Pratt parsing) to construct an Abstract Syntax Tree (AST) from tokens.

- Evaluator: Executes the AST, supporting arithmetic, boolean operations, conditionals, function definitions, and function calls.

- REPL: Interactive shell for testing code snippets.

- File Execution: For larger coding tasks (example available).

- Error Handling: Detects and reports both syntax and runtime errors.

- Testing: Comprehensive test suite covering lexer, parser, evaluator, and AST construction.

## Installation
Ensure you have Go installed (version 1.24 or above).


Clone the repo then execute the following:
```
go build .
```

## Usage
An example is available in example/01.sh


To run the interpreter in interactive mode:
You will see a prompt (>>) where you can enter code. For example:

```sh
./interpreter

starting interpreter...
// Recursion //
>>let counter = df(x) { if (x > 100) {return x} else {counter(x + 1)}}
>>counter(0)
101

// define functions //
>> let add = df(x, y) { x + y; };
>> add(5, 3);
8

// support floats //
>> let x = 3;
>> let y = 5;
>> x / y;
0

>> let y = 5.0;
>> x / y;
0.600000

// Closures are also supported //
>> let newAdder = df(x) { df(y) { x + y }; };
>> let addTwo = newAdder(2);
>> addTwo(2.5);
4.500000

//strings with concatenation //
>>let myString = "hello ";
>>myString + "world"
hello world

>>myString + "3.0"
hello 3.0
>>len(myString + "3.0")
9

>>myString + 3.0
Error: type mismatch: STRING + FLOAT

// arrays with builtin functions //
>>let x = [1,2,3,4];
let x = map(df(a){a / 0.1}, x);
>>x
[10.000000, 20.000000, 30.000000, 40.000000]

// Hash is also available //
>>let p =  [{"first": 10000, "second": 777}, {"name": "Bob", "age": 28}];
>>p[1]["name"]
Bob

>>exit
```

## Running Tests
A comprehensive test suite is provided to validate the lexer, parser, evaluator, Object, and AST implementation.

## Future Improvements
- Unicode & UTF-8 Support: Currently, the lexer processes ASCII input. Consider switching from byte to rune for full UTF-8 support.

- Enhanced Error Messages: Improve debugging capabilities with more detailed error reporting.

## Note:
This project builds on the excellent guide written by Thorsten Ball, Writing an Interpreter in Go. Highly recommend this resource to anyone interested in a deeper understanding of how languages are built.

## License
This project is licensed under the MIT License.
Feel free to adjust the content to match your project's specifics and style preferences.
