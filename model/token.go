package model

type TokenType int

const (
	LeftParentheses TokenType = iota
	RightParentheses
	LeftSquareBracket
	RightSquareBracket
	LeftBracket
	RightBracket
	None
)

func (t TokenType) String() string {
	return [...]string{"(", ")", "[", "]", "{", "}", ""}[t]
}
