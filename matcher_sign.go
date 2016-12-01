package golexer

// 操作符，分隔符，关键字
type SignMatcher struct {
	baseMatcher
	word []rune
}

func (self *SignMatcher) Match(tz *Tokenizer) (*Token, error) {

	if (tz.Count() - tz.Index()) < len(self.word) {
		return nil, nil
	}

	for i, c := range self.word {

		if tz.Peek(i) != c {
			return nil, nil
		}

	}

	tz.ConsumeMulti(len(self.word))

	return NewToken(self, tz, string(self.word), ""), nil
}

func NewSignMatcher(id int, word string) TokenMatcher {
	return &SignMatcher{
		baseMatcher: baseMatcher{id},
		word:        []rune(word),
	}
}
