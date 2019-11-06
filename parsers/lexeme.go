package parsers

type Lexeme interface {
	Token
	Scan() bool
	IsAccepted() bool
}
