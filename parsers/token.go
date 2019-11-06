package parsers

import "github.com/patrickhuber/gearly/grammars"

type Token interface {
	TokenType() grammars.TokenType
	Position() int
	Span() ReadonlyAccumulatorSpan
}
