package parsers

import (
	"fmt"
	"log"
	"os"

	"github.com/patrickhuber/gearly/charts"
	"github.com/patrickhuber/gearly/grammars"
)

type earleyParser struct {
	chart    charts.Chart
	grammar  grammars.Grammar
	logger   *log.Logger
	location int
}

func NewEarleyParser(grammar grammars.Grammar) Parser {
	chart := charts.NewChart()
	parser := &earleyParser{
		grammar:  grammar,
		chart:    chart,
		logger:   log.New(os.Stdout, "", 0),
		location: 0,
	}
	parser.initialize()
	return parser
}

func (r *earleyParser) initialize() {

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

func (parser *earleyParser) Pulse(token Token) bool {
	parser.scanPass(parser.location, token)

	tokenRecognized := len(parser.chart.Sets()) > parser.location
	if !tokenRecognized {
		return false
	}

	parser.location++
	parser.reductionPass(parser.location)
	return true
}

func (parser *earleyParser) scanPass(index int, token Token) {
	set := parser.chart.Sets()[index]
	for s := 0; s < len(set.Scans()); s++ {
		scan := set.Scans()[s]
		parser.scan(scan, token, index)
	}
}

func (parser *earleyParser) scan(state charts.State, token Token, index int) {

	postDotSymbol := state.DottedRule().PostDotSymbol()
	terminal := postDotSymbol.(grammars.Terminal)
	if terminal.TokenType() != token.TokenType() {
		return
	}

	newState := charts.NextState(state)
	if !parser.chart.Add(index+1, newState) {
		return
	}

	parser.logger.Printf("scan (%d): %v", index+1, newState)
	parser.logger.Println()
}

func (parser *earleyParser) reductionPass(index int) {
	set := parser.chart.Sets()[index]
	p := 0
	c := 0

	for resume := true; resume; {
		if c < len(set.Completions()) {
			completion := set.Completions()[c]
			parser.complete(completion, index)
			c++
		} else if p < len(set.Predictions()) {
			prediction := set.Predictions()[p]
			parser.predict(prediction, index)
			p++
		} else {
			resume = false
		}
	}
}

func (parser *earleyParser) complete(state charts.State, index int) {

	set := parser.chart.Get(state.Origin())
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
		if !parser.chart.Add(index, nextState) {
			continue
		}

		parser.logger.Printf("complete (%d): %v", index, nextState)
		parser.logger.Println()
	}
}

func (parser *earleyParser) predict(state charts.State, index int) {

	set := parser.chart.Get(index)
	postDotSymbol := state.DottedRule().PostDotSymbol()

	for p := 0; p < len(parser.grammar.Productions()); p++ {

		production := parser.grammar.Productions()[p]

		if !production.LeftHandSide().Equal(postDotSymbol) {
			continue
		}

		dottedRule := grammars.NewDottedRule(production, 0)
		newState := charts.NewState(dottedRule, index)

		if !set.Add(newState) {
			continue
		}

		parser.logger.Printf("predict (%d): %v", index, newState)
		parser.logger.Println()
	}
}

func (parser *earleyParser) IsAccepted() bool {
	sets := parser.chart.Sets()
	lastSet := sets[len(sets)-1]
	start := parser.grammar.Start()

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

func (parser *earleyParser) Location() int {
	return parser.location
}
