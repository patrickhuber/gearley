package grammars

import "fmt"

type production struct {
	leftHandSide  NonTerminal
	rightHandSide []Symbol
}

// Production represents a grammar production
type Production interface {
	LeftHandSide() NonTerminal
	RightHandSide() []Symbol
	Equal(other Production) bool
	IsEmpty() bool
}

// NewProduction creates a new production with the given name and body
func NewProduction(leftHandSide NonTerminal, rightHandSide ...Symbol) Production {
	return &production{
		leftHandSide:  leftHandSide,
		rightHandSide: rightHandSide,
	}
}

func (p *production) LeftHandSide() NonTerminal {
	return p.leftHandSide
}

func (p *production) RightHandSide() []Symbol {
	return p.rightHandSide
}

func (p *production) Equal(other Production) bool {
	if p.LeftHandSide().Name() != other.LeftHandSide().Name() {
		return false
	}
	if len(p.RightHandSide()) != len(other.RightHandSide()) {
		return false
	}
	for i, s := range p.RightHandSide() {
		o := other.RightHandSide()[i]
		if !s.Equal(o) {
			return false
		}
	}
	return true
}

func (p *production) IsEmpty() bool {
	return len(p.rightHandSide) == 0
}

func (p *production) String() string {
	value := fmt.Sprintf("%v ->", p.leftHandSide)
	for _, s := range p.rightHandSide {
		value += fmt.Sprintf(" %v", s)
	}
	return value
}
