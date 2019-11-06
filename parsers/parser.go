package parsers

// Parser defines an early parser
type Parser interface {
	Pulse(token Token) bool
	IsAccepted() bool
	Location() int
}
