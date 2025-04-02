package lexer

import (
	"strings"

	"github.com/JWSch4fer/interpreter/token"
)

type Lexer struct {
	input        string
	position     int  // current position points to current char
	readPosition int  // current reading position after current char
	ch           byte // current char under examination
}

/*
if you want to change the interpreter to consider unicode and UTF-8
switch ch from byte to rune. Note: you must also update how char are
read because they can be multiple bytes now.
*/

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

// update the current character being interpreted
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhiteSpace()

	switch l.ch {
	case '"':
		tok.Type = token.STRING
		tok.Literal = l.readString()
	case '=':
		// two char token case for ==
		if l.peekChar() == '=' {
			ch := l.ch   // current =
			l.readChar() // next =
			tok = token.Token{Type: token.EQ, Literal: string(ch) + string(l.ch)}
		} else {
			tok = newToken(token.ASSIGN, l.ch)
		}
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '-':
		tok = newToken(token.MINUS, l.ch)
		// two char token case for !=
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.NOT_EQ, Literal: string(ch) + string(l.ch)}
		} else {

			tok = newToken(token.BANG, l.ch)
		}
	case '/':
		if l.peekChar() == '/' {
			return l.readComment() // define a comment for this language
		} else {
			tok = newToken(token.SLASH, l.ch)
		}
	case '*':
		tok = newToken(token.ASTERISK, l.ch)
	case '<':
		tok = newToken(token.LT, l.ch)
	case '>':
		tok = newToken(token.GT, l.ch)

	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case ':':
		tok = newToken(token.COLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case '[':
		tok = newToken(token.LBRACKET, l.ch)
	case ']':
		tok = newToken(token.RBRACKET, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)

	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default: // look for tokens that are more than one character
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			literal := l.readNumber()
			if strings.Contains(literal, ".") {
				tok.Type = token.FLOAT
			} else {
				tok.Type = token.INT
			}
			tok.Literal = literal
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}

	l.readChar()
	return tok
}

// helper function to read tokens that are multiple characters
func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}
func isLetter(ch byte) bool {
	// we're enforcing ascii only
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

// we don't interpret whitespace
func (l *Lexer) skipWhiteSpace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

// capture all characters between "
func (l *Lexer) readString() string {
	position := l.position + 1
	for {
		l.readChar()
		if l.ch == '"' || l.ch == 0 {
			break
		}
	}
	return l.input[position:l.position]
}

func (l *Lexer) readComment() token.Token {
	// deal with the initial //
	l.readChar()
	l.readChar()

	start := l.position // the start of comment text
	for {
		// check if the current and next char for the closing delimiter
		if l.ch == '/' && l.peekChar() == '/' {
			commentText := l.input[start:l.position]
			//move forward twice
			l.readChar()
			l.readChar()
			// return comment token for ast parsing
			// we are storing comments as an ast node
			// but this allows us to use consistent error parsing
			return token.Token{Type: token.COMMENT, Literal: commentText}
		}
		if l.ch == 0 { // we return 0 if we reach token.EOF
			return token.Token{Type: token.COMMENT, Literal: "Unclosed comment!?!"}
		}
		l.readChar()
	}
}

// helper function to read numbers
// we are only allowing integers and floats for now input with Latin digits
// might add binary, hex later
func (l *Lexer) readNumber() string {
	position := l.position
	hasDot := false
	for isDigit(l.ch) || (l.ch == '.' && !hasDot) {
		if l.ch == '.' {
			hasDot = true
		}
		l.readChar()
	}
	return l.input[position:l.position]
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

// helper function to look ahead one character for tokens like ==
func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}
