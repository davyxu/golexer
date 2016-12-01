package golexer

type TokenMatcher interface {
	Match(*Tokenizer) (*Token, error)
	ID() int
}

type baseMatcher struct {
	id int
}

func (self *baseMatcher) ID() int {
	return self.id
}
