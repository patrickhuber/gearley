package grammars

import "fmt"

// DottedRule is a tuple of parse production and parse position
type DottedRule interface {
	Production() Production
	Position() int
	PreDotSymbol() Symbol
	PostDotSymbol() Symbol
	IsComplete() bool
	Equal(other DottedRule) bool
}

type dottedRule struct {
	production    Production
	position      int
	preDotSymbol  Symbol
	postDotSymbol Symbol
}

// NewDottedRule creates a new dotted rule with the production and position
func NewDottedRule(production Production, position int) DottedRule {
	preDotSymbol := getPreDotSymbol(production, position)
	postDotSymbol := getPostDotSymbol(production, position)

	return &dottedRule{
		production:    production,
		position:      position,
		preDotSymbol:  preDotSymbol,
		postDotSymbol: postDotSymbol,
	}
}

func getPreDotSymbol(production Production, position int) Symbol {
	if position == 0 || production.IsEmpty() {
		return nil
	}
	return production.RightHandSide()[position-1]
}

func getPostDotSymbol(production Production, position int) Symbol {
	if position >= len(production.RightHandSide()) {
		return nil
	}
	return production.RightHandSide()[position]
}

func (d *dottedRule) Production() Production {
	return d.production
}

func (d *dottedRule) Position() int {
	return d.position
}

func (d *dottedRule) PostDotSymbol() Symbol {
	return d.postDotSymbol
}

func (d *dottedRule) PreDotSymbol() Symbol {
	return d.preDotSymbol
}

func (d *dottedRule) Equal(other DottedRule) bool {
	return d.position == other.Position() &&
		d.production.Equal(other.Production())
}

func (d *dottedRule) String() string {
	p := d.production
	value := fmt.Sprintf("%v ->", p.LeftHandSide())
	for i, sym := range p.RightHandSide() {
		delim := " "
		if i == d.position {
			delim = "*"
		}
		value += fmt.Sprintf("%s%v", delim, sym)
	}
	if d.position == len(p.RightHandSide()) {
		value += "*"
	}
	return value
}

func (d *dottedRule) IsComplete() bool {
	return d.position == len(d.production.RightHandSide())
}
