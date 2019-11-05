package grammars

import (
	"strings"
)

// Terminal represents a grammar terminal
type Terminal interface {
	Symbol
}

// LiteralTerminal is a string literal terminal
type LiteralTerminal interface {
	Terminal
	Literal() string
}

type literalTerminal struct {
	literal string
}

// NewLiteralTerminal creates a new string literal terminal
func NewLiteralTerminal(literal string) LiteralTerminal {
	return &literalTerminal{
		literal: literal,
	}
}

func (l *literalTerminal) Literal() string {
	return l.literal
}

func (l *literalTerminal) Equal(sym Symbol) bool {
	other, ok := sym.(LiteralTerminal)
	if !ok {
		return false
	}

	return strings.Compare(l.literal, other.Literal()) == 0
}

func (l *literalTerminal) String() string {
	return l.literal
}

func (l *literalTerminal) SymbolType() SymbolType {
	return SymbolTypeTerminal
}
