package grammars_test

import (
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/patrickhuber/gearly/grammars"
)

var _ = Describe("Production", func() {
	var (
		production grammars.Production
	)
	BeforeEach(func() {
		leftHandSide := grammars.NewNonTerminal("S")
		rightHandSide := []grammars.Symbol{
			grammars.NewNonTerminal("S"),
			grammars.NewLiteralTerminalFromString("s"),
		}
		production = grammars.NewProduction(leftHandSide, rightHandSide...)
	})
	Describe("LeftHandSide", func() {
		It("returns the left hand side", func() {
			symbol := production.LeftHandSide()
			Expect(symbol).ToNot(BeNil())
			Expect(symbol.Name()).To(Equal("S"))
		})
	})
	Describe("RightHandSide", func() {
		It("returns the right hand side", func() {
			rhs := production.RightHandSide()
			Expect(rhs).ToNot(BeNil())
			Expect(len(rhs)).To(Equal(2))
			Expect(rhs[0]).ToNot(BeNil())
			Expect(rhs[1]).ToNot(BeNil())
		})
	})
	Describe("Equal", func() {
		Context("When Equal", func() {
			It("returns true", func() {
				leftHandSide := grammars.NewNonTerminal("S")
				rightHandSide := []grammars.Symbol{
					grammars.NewNonTerminal("S"),
					grammars.NewLiteralTerminalFromString("s"),
				}
				p := grammars.NewProduction(leftHandSide, rightHandSide...)
				Expect(production.Equal(p)).To(BeTrue())
			})
		})
		Context("When Has Different Right Hand Side", func() {
			It("returns false", func() {
				leftHandSide := grammars.NewNonTerminal("S")
				rightHandSide := []grammars.Symbol{
					grammars.NewNonTerminal("S"),
					grammars.NewLiteralTerminalFromString("s"),
					grammars.NewLiteralTerminalFromString("v"),
				}
				p := grammars.NewProduction(leftHandSide, rightHandSide...)
				Expect(production.Equal(p)).To(BeFalse())
			})
		})
		Context("When Has Different Left Hand Side", func() {
			It("returns false", func() {
				leftHandSide := grammars.NewNonTerminal("M")
				rightHandSide := []grammars.Symbol{
					grammars.NewNonTerminal("S"),
					grammars.NewLiteralTerminalFromString("s"),
				}
				p := grammars.NewProduction(leftHandSide, rightHandSide...)
				Expect(production.Equal(p)).To(BeFalse())
			})
		})
	})
	Describe("String", func() {
		It("returns the production as a string", func() {
			Expect(fmt.Sprintf("%v", production)).To(Equal("S -> S s"))
		})
	})
})
