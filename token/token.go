package token

type TokenType string

// string is not the most performant but it is pragmatic
type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL" //something the interpreter can't interpret
	EOF     = "EOF"
	EXIT    = "EXIT"
	COMMENT = "COMMENT"

	//identifier and literals
	IDENT  = "IDENT"  // add, x, y, foo ...
	INT    = "INT"    // 1234567890
	FLOAT  = "FLOAT"  // 3.14159
	STRING = "STRING" // "blah blah"

	//OPERATORS
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"

	LT     = "<"
	GT     = ">"
	EQ     = "=="
	NOT_EQ = "!="

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"
	COLON     = ":"

	LPAREN   = "("
	RPAREN   = ")"
	LBRACE   = "{"
	RBRACE   = "}"
	LBRACKET = "["
	RBRACKET = "]"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
	NULL     = "NULL"
)

var keywords = map[string]TokenType{
	"df":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
	"exit":   EXIT,
	"NULL":   NULL,
}

// define language specified keywords versus user defined variables
func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
