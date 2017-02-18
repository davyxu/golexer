package golexer

import "errors"

type Parser struct {
	lexer *Lexer

	curr *Token

	srcName string

	errFunc func(error)
}

func (self *Parser) Lexer() *Lexer {
	return self.lexer
}

func (self *Parser) Expect(id int) Token {

	if self.TokenID() != id {
		panic(errors.New("Expect " + self.lexer.MatcherString(id)))
	}

	t := *self.curr

	self.NextToken()

	return t
}

func (self *Parser) NextToken() {

	token, err := self.lexer.Read()

	if err != nil {
		panic(err)
	}

	self.curr = token
}

func (self *Parser) TokenID() int {
	return self.curr.MatcherID()
}

func (self *Parser) TokenValue() string {
	return self.curr.Value()
}

func (self *Parser) MatcherName() string {
	return self.curr.MatcherName()
}

func (self *Parser) MatcherString() string {
	return self.curr.MatcherString()
}

func (self *Parser) TokenRaw() string {

	return self.curr.Raw()
}

func (self *Parser) TokenPos() TokenPos {
	return TokenPos{
		Line:       self.curr.line,
		Col:        self.curr.index,
		SourceName: self.srcName,
	}
}

func NewParser(l *Lexer, srcName string) *Parser {

	return &Parser{
		lexer:   l,
		srcName: srcName,
	}

}
