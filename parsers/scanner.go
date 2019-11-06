package parsers

import (
	"bufio"
	"io"

	"github.com/patrickhuber/gearly/grammars"
)

// Scanner defines a text scanner interface that communicates with a parser to parse input
type Scanner interface {
	Read() (bool, error)
	Position() int
	Line() int
	Column() int
	EndOfStream() bool
	RunToEnd() (bool, error)
}

type scanner struct {
	position    int
	line        int
	column      int
	parser      Parser
	reader      *bufio.Reader
	accumulator Accumulator
}

// NewScanner creates a new scanner for the parser and io reader
func NewScanner(parser Parser, reader io.Reader) Scanner {
	return &scanner{
		position:    0,
		line:        0,
		column:      0,
		parser:      parser,
		reader:      bufio.NewReader(reader),
		accumulator: NewAccumulator(),
	}
}

func (s *scanner) Read() (bool, error) {
	if s.EndOfStream() {
		return false, nil
	}

	ch, err := s.readRune()
	if err != nil {
		return false, err
	}

	s.updatePositionMetrics(ch)
	s.accumulator.Accumulate(ch)

	terminal := grammars.NewLiteralTerminal([]rune{' '})
	token := NewLiteralLexeme(terminal, s.position, s.accumulator.Span(s.position, 1))
	return s.parser.Pulse(token), nil
}

func (s *scanner) readRune() (rune, error) {
	ch, _, err := s.reader.ReadRune()
	return ch, err
}

func (s *scanner) updatePositionMetrics(ch rune) {
	s.position++
	if ch == '\n' {
		s.column = 0
		s.line++
	} else {
		s.column++
	}
}

func (s *scanner) Position() int {
	return s.position
}

func (s *scanner) Line() int {
	return s.line
}

func (s *scanner) Column() int {
	return s.column
}

func (s *scanner) EndOfStream() bool {
	_, err := s.reader.Peek(1)
	return err != nil
}

func (s *scanner) RunToEnd() (bool, error) {
	for !s.EndOfStream() {
		ok, err := s.Read()
		if err != nil {
			return false, err
		}
		if !ok {
			return false, nil
		}
	}
	return s.parser.IsAccepted(), nil
}
