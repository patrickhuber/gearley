package parsers

import "github.com/patrickhuber/gearly/grammars"

type literalLexeme struct {
	position int
	literal  grammars.LiteralTerminal
	span     AccumulatorSpan
}

func NewLiteralLexeme(literal grammars.LiteralTerminal, position int, span AccumulatorSpan) Lexeme {
	return &literalLexeme{
		position: position,
		span:     span,
		literal:  literal,
	}
}

func (l *literalLexeme) Position() int {
	return l.position
}

func (l *literalLexeme) TokenType() grammars.TokenType {
	return l.literal.TokenType()
}

func (l *literalLexeme) Span() ReadonlyAccumulatorSpan {
	return l.span
}

func (l *literalLexeme) IsAccepted() bool {
	return l.span.Length() == len(l.literal.Literal())
}

func (l *literalLexeme) Scan() bool {
	peek := l.span.Peek()
	if peek == EOF {
		return false
	}
	if peek != l.literal.Literal()[l.span.Length()] {
		return false
	}
	return l.span.Grow()
}
