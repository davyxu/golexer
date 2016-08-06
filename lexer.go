package golexer

type TokenMatcher interface {
	Match(*Tokenizer) *Token
}

type matcherMeta struct {
	m      TokenMatcher
	ignore bool
}

type Lexer struct {
	matchers []matcherMeta

	comm chan *Token
}

func (self *Lexer) AddMatcher(m TokenMatcher) {
	self.matchers = append(self.matchers, matcherMeta{
		m:      m,
		ignore: false,
	})
}

func (self *Lexer) AddIgnoreMatcher(m TokenMatcher) {
	self.matchers = append(self.matchers, matcherMeta{
		m:      m,
		ignore: true,
	})
}

func (self *Lexer) Start(src string) {

	if self.comm != nil {
		close(self.comm)
	}

	self.comm = make(chan *Token)

	go self.tokenWorker(src)
}

func (self *Lexer) Read() *Token {
	if self.comm == nil {
		return nil
	}

	return <-self.comm
}

func (self *Lexer) tokenWorker(src string) {

	tn := NewTokenizer(src)

	if len(self.matchers) > 0 {
		for !tn.EOF() {

			for _, mm := range self.matchers {

				token := mm.m.Match(tn)

				if token == nil {
					continue
				}

				if mm.ignore {
					break
				}

				self.comm <- token

			}
		}
	}

	self.comm <- nil
}

func NewLexer() *Lexer {

	return &Lexer{}

}
