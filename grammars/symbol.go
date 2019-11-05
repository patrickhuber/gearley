package grammars

// SymbolType defines the type of symbol
type SymbolType string

const (
	// SymbolTypeTerminal is a terminal symbol type
	SymbolTypeTerminal SymbolType = "Terminal"
	// SymbolTypeNonTerminal is a non terminal symbol type
	SymbolTypeNonTerminal SymbolType = "NonTerminal"
)

// Symbol represents a base type for terminals and non terminals
type Symbol interface {
	Equal(s Symbol) bool
	SymbolType() SymbolType
}
