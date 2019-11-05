package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"github.com/patrickhuber/gearly/charts"

	"github.com/patrickhuber/gearly/grammars"
)

type recognizer struct {
	grammar grammars.Grammar
	chart   charts.Chart
}

type Recognizer interface {
	MatchReader(reader io.Reader) (bool, error)
	MatchString(s string) (bool, error)
}

func NewRecognizer(grammar grammars.Grammar) Recognizer {
	return &recognizer{
		grammar: grammar,
	}
}

func (r *recognizer) MatchString(s string) (bool, error) {
	return r.MatchReader(strings.NewReader(s))
}

func (r *recognizer) MatchReader(reader io.Reader) (bool, error) {

	r.initialize()

	bufferedReader := bufio.NewReader(reader)

	index := 0
	for {
		ch, _, err := bufferedReader.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			return false, err
		}
		if !r.pulse(ch, index) {
			break
		}
		index++
	}

	return r.isAccepted(), nil
}

func (r *recognizer) pulse(ch rune, index int) bool {
	r.scanPass(index, ch)

	tokenRecognized := len(r.chart.Sets()) > index+1
	if !tokenRecognized {
		return false
	}

	index++
	r.reductionPass(index)
	return true
}

func (r *recognizer) scanPass(index int, ch rune) {
	set := r.chart.Sets()[index]
	for s := 0; s < len(set.Scans()); s++ {
		scan := set.Scans()[s]
		r.scan(scan, ch, index)
	}
}

func (r *recognizer) reductionPass(index int) {
	set := r.chart.Sets()[index]
	p := 0
	c := 0

	for resume := true; resume; {
		if c < len(set.Completions()) {
			completion := set.Completions()[c]
			r.complete(completion, index)
			c++
		} else if p < len(set.Predictions()) {
			prediction := set.Predictions()[p]
			r.predict(prediction, index)
			p++
		} else {
			resume = false
		}
	}
}

func (r *recognizer) isAccepted() bool {
	sets := r.chart.Sets()
	lastSet := sets[len(sets)-1]
	start := r.grammar.Start()

	for c := 0; c < len(lastSet.Completions()); c++ {
		completion := lastSet.Completions()[c]
		// must be a completion state
		if !completion.DottedRule().IsComplete() {
			continue
		}
		//  the origin should be 0
		if completion.Origin() != 0 {
			continue
		}
		// the start symbol of the grammar should equal the left hand side of the production
		production := completion.DottedRule().Production()
		if production.LeftHandSide().Equal(start) {
			return true
		}
	}
	return false
}

func (r *recognizer) initialize() {

	start := r.grammar.Start()
	r.chart = charts.NewChart()

	for p := 0; p < len(r.grammar.Productions()); p++ {

		production := r.grammar.Productions()[p]

		if production.LeftHandSide() != start {
			continue
		}

		dottedRule := grammars.NewDottedRule(production, 0)
		state := charts.NewState(dottedRule, 0)
		if !r.chart.Add(0, state) {
			continue
		}

		fmt.Printf("initialize(0): %v", state)
		fmt.Println()
	}

	r.reductionPass(0)
}

func (r *recognizer) complete(state charts.State, index int) {

	set := r.chart.Get(state.Origin())
	search := state.DottedRule().Production().LeftHandSide()

	// look for prediction state that created this node
	for p := 0; p < len(set.Predictions()); p++ {

		prediction := set.Predictions()[p]
		dottedRule := prediction.DottedRule()
		postDotSymbol := dottedRule.PostDotSymbol()

		if !postDotSymbol.Equal(search) {
			continue
		}

		nextState := charts.NextState(prediction)
		if !r.chart.Add(index, nextState) {
			continue
		}

		fmt.Printf("complete (%d): %v", index, nextState)
		fmt.Println()
	}
}

func (r *recognizer) scan(state charts.State, ch rune, index int) {

	postDotSymbol := state.DottedRule().PostDotSymbol()

	switch v := postDotSymbol.(type) {
	case grammars.LiteralTerminal:
		literal := v.Literal()
		if len(literal) == 0 {
			return
		}
		if !strings.HasPrefix(literal, string(ch)) {
			return
		}
	}

	newState := charts.NextState(state)
	if !r.chart.Add(index+1, newState) {
		return
	}

	fmt.Printf("scan (%d): %v", index+1, newState)
	fmt.Println()
}

func (r *recognizer) predict(state charts.State, index int) {

	set := r.chart.Get(index)
	postDotSymbol := state.DottedRule().PostDotSymbol()

	for p := 0; p < len(r.grammar.Productions()); p++ {

		production := r.grammar.Productions()[p]

		if !production.LeftHandSide().Equal(postDotSymbol) {
			continue
		}

		dottedRule := grammars.NewDottedRule(production, 0)
		newState := charts.NewState(dottedRule, index)

		if !set.Add(newState) {
			continue
		}

		fmt.Printf("predict (%d): %v", index, newState)
		fmt.Println()
	}
}
