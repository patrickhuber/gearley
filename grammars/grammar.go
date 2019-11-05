package grammars

type grammar struct {
	productions []Production
	start       NonTerminal
}

// Grammar defines a language grammar
type Grammar interface {
	Start() NonTerminal
	Productions() []Production
}

// NewGrammar creates a new Grammar for the start symbol and set of productions
func NewGrammar(start NonTerminal, productions ...Production) Grammar {
	return &grammar{
		start:       start,
		productions: productions,
	}
}

func (g *grammar) Start() NonTerminal {
	return g.start
}

func (g *grammar) Productions() []Production {
	return g.productions
}
