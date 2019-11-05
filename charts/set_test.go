package charts_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/patrickhuber/gearly/charts"
	"github.com/patrickhuber/gearly/grammars"
)

var _ = Describe("Set", func() {
	var (
		set charts.Set
	)
	BeforeEach(func() {
		set = charts.NewSet()
	})
	Describe("Scans", func() {})
	Describe("Completions", func() {
		It("returns empty list", func() {
			completions := set.Completions()
			Expect(completions).ToNot(BeNil())
			Expect(len(completions)).To(Equal(0))
		})
	})
	Describe("Predictions", func() {})
	Describe("Add", func() {
		Context("when completed", func() {
			It("adds to copmleted list", func() {
				lhs := grammars.NewNonTerminal("S")
				rhs := grammars.NewLiteralTerminal("s")
				production := grammars.NewProduction(lhs, rhs)
				dottedRule := grammars.NewDottedRule(production, 1)
				state := charts.NewState(dottedRule, 0)
				value := set.Add(state)
				Expect(value).To(BeTrue())
				completions := set.Completions()
				Expect(len(completions)).To(Equal(1))
			})
		})
	})
})
