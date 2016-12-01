package golexer

import (
	"unicode"
)

// 操作符，分隔符，关键字
type KeywordMatcher struct {
	baseMatcher
	word []rune
}

func isKeyword(r rune) bool {
	return unicode.IsLetter(r) ||
		r == '_'
}

func (self *KeywordMatcher) Match(tz *Tokenizer) (*Token, error) {

	if (tz.Count() - tz.Index()) < len(self.word) {
		return nil, nil
	}

	for i, c := range self.word {

		if !isKeyword(c) {
			return nil, nil
		}

		if tz.Peek(i) != c {
			return nil, nil
		}

	}

	pc := tz.Peek(len(self.word))
	if isKeyword(pc) {
		return nil, nil
	}

	tz.ConsumeMulti(len(self.word))

	return NewToken(self, tz, string(self.word), ""), nil
}

func NewKeywordMatcher(id int, word string) TokenMatcher {
	self := &KeywordMatcher{
		baseMatcher: baseMatcher{id},
		word:        []rune(word),
	}

	for _, c := range self.word {
		if !isKeyword(c) {
			panic("not keyword")
		}
	}

	return self
}
