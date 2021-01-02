package token

type TokenType string

const (
	// Special types
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers + literal
	IDENT = "IDENT" // add, foobar, x, y, ...
	INT   = "INT"   // 1234

	// Operators
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

type Token struct {
	Type    TokenType
	Literal string
}

var keywords = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok { // tok is the value if it is found; ok is false if the value is not found
		return tok
	}
	return IDENT
}
