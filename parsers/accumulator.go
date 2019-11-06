package parsers

type ReadonlyAccumulatorSpan interface {
	Start() int
	Length() int
	Capture() []rune
	RuneAt(index int) rune
	Peek() rune
}

const EOF = rune(-1)

type AccumulatorSpan interface {
	ReadonlyAccumulatorSpan
	Grow() bool
}

type accumulatorSpan struct {
	accumulator Accumulator
	start       int
	length      int
}

func (s *accumulatorSpan) Start() int {
	return s.start
}

func (s *accumulatorSpan) Length() int {
	return s.length
}

func (s *accumulatorSpan) Capture() []rune {
	return s.accumulator.Capture()[s.start:s.length]
}

func (s *accumulatorSpan) Grow() bool {
	if s.Peek() == EOF {
		return false
	}
	s.length++
	return true
}

func (s *accumulatorSpan) Peek() rune {
	if s.length >= s.accumulator.Length() {
		return EOF
	}
	return s.accumulator.RuneAt(s.length)
}

func (s *accumulatorSpan) RuneAt(index int) rune {
	return s.accumulator.RuneAt(index)
}

type Accumulator interface {
	Accumulate(ch rune)
	Span(start, length int) AccumulatorSpan
	Length() int
	Capture() []rune
	RuneAt(index int) rune
}

func NewAccumulator() Accumulator {
	return &accumulator{}
}

type accumulator struct {
	builder []rune
}

func (a *accumulator) Accumulate(ch rune) {
	a.builder = append(a.builder, ch)
}

func (a *accumulator) Length() int {
	return len(a.builder)
}

func (a *accumulator) Capture() []rune {
	return a.builder
}

func (a *accumulator) String() string {
	return string(a.builder)
}

func (a *accumulator) RuneAt(index int) rune {
	return a.builder[index]
}

func (a *accumulator) Span(start, length int) AccumulatorSpan {
	return &accumulatorSpan{
		accumulator: a,
		start:       start,
		length:      length,
	}
}
