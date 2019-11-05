package grammars

import "strings"

type nonTerminal struct {
	name string
}

// NonTerminal represents a production name in a grammar
type NonTerminal interface {
	Symbol
	Name() string
}

// NewNonTerminal creates a new NonTerminal with the given name
func NewNonTerminal(name string) NonTerminal {
	return &nonTerminal{
		name: name,
	}
}

func (s *nonTerminal) Name() string {
	return s.name
}

func (s *nonTerminal) Equal(sym Symbol) bool {
	other, ok := sym.(NonTerminal)
	if !ok {
		return false
	}
	return strings.Compare(s.name, other.Name()) == 0
}

func (s *nonTerminal) String() string {
	return s.name
}

func (s *nonTerminal) SymbolType() SymbolType {
	return SymbolTypeNonTerminal
}
