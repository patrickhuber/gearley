package parsers_test

import (
	"strings"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/patrickhuber/gearly/parsers"
)

type FakeParser struct {
	pulseCount   int
	currentCount int
}

func NewFakeParser(pulseCount int) parsers.Parser {
	return &FakeParser{
		pulseCount:   pulseCount,
		currentCount: 0,
	}
}

func (p *FakeParser) Pulse(token parsers.Token) bool {
	if p.currentCount >= p.pulseCount {
		return false
	}
	p.currentCount++
	return true
}

func (p *FakeParser) IsAccepted() bool {
	return p.currentCount >= p.pulseCount
}

func (p *FakeParser) Location() int {
	return p.currentCount
}

func NewScanner(text string) parsers.Scanner {
	reader := strings.NewReader(text)
	parser := NewFakeParser(len(text))
	return parsers.NewScanner(parser, reader)
}

var _ = Describe("Scanner", func() {
	Describe("Read", func() {
		It("reads input", func() {
			scanner := NewScanner(" ")
			result, err := scanner.Read()
			Expect(err).To(BeNil())
			Expect(result).To(BeTrue())
			Expect(scanner.EndOfStream()).To(BeTrue())
		})
	})
	Describe("Position", func() {
		It("updates position", func() {
			scanner := NewScanner(" ")
			Expect(scanner.Position()).To(Equal(0))
			result, err := scanner.Read()
			Expect(err).To(BeNil())
			Expect(result).To(BeTrue())
			Expect(scanner.Position()).To(Equal(1))
		})
	})
	Describe("Column", func() {
		Context("when single line", func() {
			It("equals position", func() {
				scanner := NewScanner("this is a string")
				for {
					result, err := scanner.Read()
					Expect(err).To(BeNil())
					if !result {
						break
					}
					Expect(scanner.Position()).To(Equal(scanner.Column()))
				}
			})
		})
		Context("when newline", func() {
			It("resets column", func() {
				scanner := NewScanner("test\nfile")
				for {
					result, err := scanner.Read()
					Expect(err).To(BeNil())
					if !result {
						break
					}
					if scanner.Position() < 5 {
						Expect(scanner.Column()).To(Equal(scanner.Position()))
					} else {
						Expect(scanner.Column()).To(Equal(scanner.Position() - 5))
					}
				}
			})
		})
	})
	Describe("Line", func() {
		Context("when single line", func() {
			It("remains zero", func() {
				scanner := NewScanner("this is a string")
				for {
					result, err := scanner.Read()
					Expect(err).To(BeNil())
					if !result {
						break
					}
					Expect(scanner.Line()).To(Equal(0))
				}
			})
		})
		Context("when multi line", func() {
			It("increments", func() {
				scanner := NewScanner("line 1\nline 2\nline 3")
				for {
					result, err := scanner.Read()
					Expect(err).To(BeNil())
					if !result {
						break
					}
					if scanner.Position() < 7 {
						Expect(scanner.Line()).To(Equal(0))
					} else if scanner.Position() < 14 {
						Expect(scanner.Line()).To(Equal(1))
					} else {
						Expect(scanner.Line()).To(Equal(2))
					}
				}
			})
		})
	})
	Describe("RunToEnd", func() {
		It("reads to end of input", func() {
			scanner := NewScanner("this\nis\na\ntest")
			result, err := scanner.RunToEnd()
			Expect(err).To(BeNil())
			Expect(result).To(BeTrue())
		})
	})
})
