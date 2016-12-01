package golexer

import "unicode"

// 标识符
type IdentifierMatcher struct {
	baseMatcher
}

func (self *IdentifierMatcher) Match(tz *Tokenizer) (*Token, error) {

	if !unicode.IsLetter(tz.Current()) && tz.Current() != '_' {
		return nil, nil
	}

	begin := tz.Index()

	for {

		tz.ConsumeOne()

		if !unicode.IsLetter(tz.Current()) && tz.Current() != '_' {
			break
		}

	}

	return NewToken(self, tz, tz.StringRange(begin, tz.index), ""), nil
}

func NewIdentifierMatcher(id int) TokenMatcher {
	return &IdentifierMatcher{
		baseMatcher{id},
	}
}
