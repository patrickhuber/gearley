package charts

import "github.com/patrickhuber/gearly/grammars"

type set struct {
	scans       []State
	predictions []State
	completions []State
}

type Set interface {
	Scans() []State
	Completions() []State
	Predictions() []State
	Add(s State) bool
}

// NewSet creates a new set
func NewSet() Set {
	return &set{
		scans:       make([]State, 0),
		predictions: make([]State, 0),
		completions: make([]State, 0),
	}
}

func (s *set) Scans() []State {
	return s.scans
}

func (s *set) Predictions() []State {
	return s.predictions
}

func (s *set) Completions() []State {
	return s.completions
}

func (s *set) Add(state State) bool {
	if state.DottedRule().IsComplete() {
		return s.addUniqueCompletion(state)
	}

	postDotSymbol := state.DottedRule().PostDotSymbol()
	if postDotSymbol == nil {
		return false
	}

	if postDotSymbol.SymbolType() == grammars.SymbolTypeTerminal {
		return s.addUniqueScan(state)
	}

	return s.addUniquePrediction(state)
}

func (s *set) addUniqueCompletion(state State) bool {
	ok, completions := addUnique(s.completions, state)
	if !ok {
		return false
	}
	s.completions = completions
	return true
}

func (s *set) addUniqueScan(state State) bool {
	ok, scans := addUnique(s.scans, state)
	if !ok {
		return false
	}
	s.scans = scans
	return true
}

func (s *set) addUniquePrediction(state State) bool {
	ok, predictions := addUnique(s.predictions, state)
	if !ok {
		return false
	}
	s.predictions = predictions
	return true
}

func addUnique(states []State, state State) (bool, []State) {
	for _, s := range states {
		if s.Equal(state) {
			return false, states
		}
	}
	return true, append(states, state)
}
