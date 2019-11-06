package grammars

// LiteralTerminal is a string literal terminal
type LiteralTerminal interface {
	Terminal
	Literal() []rune
}

type literalTerminal struct {
	literal   []rune
	tokenType TokenType
}

// NewLiteralTerminal creates a new literal terminal from the rune array
func NewLiteralTerminal(literal []rune) LiteralTerminal {
	tokenType := NewTokenType(string(literal))
	return &literalTerminal{
		literal:   literal,
		tokenType: tokenType,
	}
}

// NewLiteralTerminalFromString creates a new literal terminal from the string
func NewLiteralTerminalFromString(literal string) LiteralTerminal {
	return NewLiteralTerminal([]rune(literal))
}

func (l *literalTerminal) Literal() []rune {
	return l.literal
}

func (l *literalTerminal) Equal(sym Symbol) bool {
	other, ok := sym.(LiteralTerminal)
	if !ok {
		return false
	}

	if len(l.literal) != len(other.Literal()) {
		return false
	}
	for i := 0; i < len(l.literal); i++ {
		if l.literal[i] != other.Literal()[i] {
			return false
		}
	}
	return true
}

func (l *literalTerminal) String() string {
	return string(l.literal)
}

func (l *literalTerminal) SymbolType() SymbolType {
	return SymbolTypeTerminal
}

func (l *literalTerminal) CanApply(ch rune) bool {
	if len(l.literal) == 0 {
		return false
	}
	return l.literal[0] == ch
}

func (l *literalTerminal) TokenType() TokenType {
	return l.tokenType
}
