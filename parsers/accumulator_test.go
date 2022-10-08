package parsers_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/patrickhuber/gearly/parsers"
)

var _ = Describe("Accumulator", func() {
	It("can add", func() {
		a := parsers.NewAccumulator()
		Expect(a.Length()).To(Equal(0))
		a.Accumulate('a')
		Expect(a.Length()).To(Equal(1))
	})
	It("returns capture", func() {
		a := parsers.NewAccumulator()
		a.Accumulate('a')
		a.Accumulate('b')
		capture := a.Capture()
		Expect(len(capture)).To(Equal(2))
		Expect(capture[0]).To(Equal('a'))
		Expect(capture[1]).To(Equal('b'))
	})
	It("returns span", func() {
		a := parsers.NewAccumulator()
		a.Accumulate('a')
		a.Accumulate('b')
		a.Accumulate('c')
		s := a.Span(1, 1)
		Expect(s.Length()).To(Equal(1))
		Expect(s.RuneAt(0)).To(Equal('b'))
	})
})
