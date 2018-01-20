// colin
//
// 1/18/18
//  
// 11:14 AM

package token

type TokenType string

type Token struct {
	Type TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF = "EOF"
	IDENT = "IDENT"
	INT = "INT"
	ASSIGN = "="
	PLUS = "+"
	COMMA = ","
	SEMICOLON = ";"
	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"
	FUNCTION = "FUNCTION"
	LET = "LET"
)

