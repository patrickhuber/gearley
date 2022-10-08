package charts_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/patrickhuber/gearly/charts"
	"github.com/patrickhuber/gearly/grammars"
)

var _ = Describe("State", func() {
	var (
		state charts.State
	)
	BeforeEach(func() {
		lhs := grammars.NewNonTerminal("S")
		rhs := []grammars.Symbol{
			grammars.NewLiteralTerminalFromString("s"),
		}
		production := grammars.NewProduction(lhs, rhs...)
		dottedRule := grammars.NewDottedRule(production, 0)
		state = charts.NewState(dottedRule, 0)
	})
	Describe("DottedRule", func() {
		It("returns dotted rule", func() {
			dottedRule := state.DottedRule()
			Expect(dottedRule).ToNot(BeNil())
			Expect(dottedRule.Position()).To(Equal(0))
			Expect(dottedRule.IsComplete()).To(BeFalse())
		})
	})
	Describe("Origin", func() {
		It("returns origin", func() {
			origin := state.Origin()
			Expect(origin).To(Equal(0))
		})
	})
	Describe("Equal", func() {
		Context("when equal", func() {
			It("returns true", func() {
				lhs := grammars.NewNonTerminal("S")
				rhs := []grammars.Symbol{
					grammars.NewLiteralTerminalFromString("s"),
				}
				production := grammars.NewProduction(lhs, rhs...)
				dottedRule := grammars.NewDottedRule(production, 0)
				expected := charts.NewState(dottedRule, 0)

				Expect(state.Equal(expected)).To(BeTrue())
			})
		})
	})
})
