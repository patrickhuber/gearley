package grammars_test

import (
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/patrickhuber/gearly/grammars"
)

var _ = Describe("NonTerminal", func() {
	var (
		nonTerminal grammars.NonTerminal
	)
	BeforeEach(func() {
		nonTerminal = grammars.NewNonTerminal("test")
	})
	Describe("Name", func() {
		It("returns name", func() {
			Expect(nonTerminal.Name()).To(Equal("test"))
		})
	})
	Describe("Equal", func() {
		It("returns true when same", func() {
			expected := grammars.NewNonTerminal("test")
			Expect(nonTerminal.Equal(expected)).To(BeTrue())
		})
		It("returns false when different symbol", func() {
			expected := grammars.NewLiteralTerminalFromString("test")
			Expect(nonTerminal.Equal(expected)).To(BeFalse())
		})
		It("returns false when different name", func() {
			expected := grammars.NewNonTerminal("false")
			Expect(nonTerminal.Equal(expected)).To(BeFalse())
		})
	})
	Describe("String", func() {
		It("returns name", func() {
			Expect(fmt.Sprintf("%v", nonTerminal)).To(Equal("test"))
		})
	})
})
