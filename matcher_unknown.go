package golexer

type UnknownMatcher struct {
}

func (self *UnknownMatcher) Match(tz *Tokenizer) *Token {

	begin := tz.Index()

	tz.ConsumeOne()

	return NewToken(self, tz.StringRange(begin, tz.Index()))
}
