package parsers_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/patrickhuber/gearly/grammars"
	"github.com/patrickhuber/gearly/parsers"
)

var _ = Describe("EarleyParser", func() {
	Describe("Pulse", func() {
		It("returns true", func() {
			start := grammars.NewNonTerminal("S")
			a := grammars.NewNonTerminal("A")
			b := grammars.NewNonTerminal("B")
			letterA := grammars.NewLiteralTerminalFromString("a")
			letterB := grammars.NewLiteralTerminalFromString("b")

			/*
				S -> S | SS | A | B
				A -> 'a'
				B -> 'b'
			*/
			grammar := grammars.NewGrammar(
				start,
				grammars.NewProduction(start, start, start),
				grammars.NewProduction(start, a),
				grammars.NewProduction(start, b),
				grammars.NewProduction(a, letterA),
				grammars.NewProduction(b, letterB),
			)

			parser := parsers.NewEarleyParser(grammar)
			accumulator := parsers.NewAccumulator()
			text := "abaab"
			chmap := map[rune]grammars.LiteralTerminal{
				'a': letterA,
				'b': letterB,
			}
			for i, ch := range text {
				accumulator.Accumulate(ch)
				token := parsers.NewLiteralLexeme(chmap[ch], i, accumulator.Span(i, 1))
				Expect(parser.Pulse(token)).To(BeTrue())
			}
			Expect(parser.IsAccepted()).To(BeTrue())
		})
	})
})
