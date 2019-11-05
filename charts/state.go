package charts

import (
	"fmt"

	"github.com/patrickhuber/gearly/grammars"
)

type state struct {
	dottedRule grammars.DottedRule
	origin     int
}

type State interface {
	DottedRule() grammars.DottedRule
	Origin() int
	Equal(s State) bool
}

func NewState(dottedRule grammars.DottedRule, origin int) State {
	return &state{
		dottedRule: dottedRule,
		origin:     origin,
	}
}

func NextState(s State) State {
	dottedRule := s.DottedRule()
	newDottedRule := grammars.NewDottedRule(dottedRule.Production(), dottedRule.Position()+1)
	return &state{
		dottedRule: newDottedRule,
		origin:     s.Origin(),
	}
}

func (s *state) Origin() int {
	return s.origin
}

func (s *state) DottedRule() grammars.DottedRule {
	return s.dottedRule
}

func (s *state) Equal(other State) bool {
	return s.origin == other.Origin() &&
		s.dottedRule.Equal(other.DottedRule())
}

func (s *state) String() string {
	return fmt.Sprintf("%s, %d", s.dottedRule, s.origin)
}
