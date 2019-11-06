package grammars

// Terminal represents a grammar terminal
type Terminal interface {
	Symbol
	CanApply(ch rune) bool
	TokenType() TokenType
}
