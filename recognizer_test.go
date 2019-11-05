package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/patrickhuber/gearly/grammars"
)

var _ = Describe("Recognizer", func() {
	Describe("Match", func() {
		It("recognizes input", func() {

			start := grammars.NewNonTerminal("S")
			a := grammars.NewNonTerminal("A")
			b := grammars.NewNonTerminal("B")
			letterA := grammars.NewLiteralTerminal("a")
			letterB := grammars.NewLiteralTerminal("b")

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
			recognizer := NewRecognizer(grammar)
			ok, err := recognizer.MatchString("abaab")
			Expect(ok).To(BeTrue())
			Expect(err).To(BeNil())
		})
	})
})
