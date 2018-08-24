package token

// Type is the type of token in Monkey source code.
type Type string

// Token represents a token in Monkey source code.
type Token struct {
	Type    Type
	Literal string
}

// Token Types
const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers and literals
	IDENT = "IDENT"
	INT   = "INT"

	ASSIGN = ":"
	PLUS   = "+"
	ONTO   = "->"

	// Delimeters
	COMMA = ","

	LPAREN = "("
	RPAREN = ")"

	// Keywords
	THEN = "THEN"
	GO   = "|>"
)

// New returns a new token.
func New(tokenType Type, ch byte) Token {
	return Token{Type: tokenType, Literal: string(ch)}
}
