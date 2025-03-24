package token

type ToKenType string

// string is not the most performant but it is pragmatic
type Token struct {
	Type    ToKenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL" //something the interpreter can't interpret
	EOF     = "EOF"

	//identifier and literals
	IDENT = "IDENT" // add, x, y, foo ...
	INT   = "INT"   //init

	//OPERATORS
	ASSIGN = "="
	PLUS   = "+"

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
)
