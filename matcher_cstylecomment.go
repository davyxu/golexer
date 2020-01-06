package golexer

import "reflect"

// #开头的行注释
type CStyleCommentMatcher struct {
	baseMatcher
}

func (self *CStyleCommentMatcher) String() string {
	return reflect.TypeOf(self).Elem().Name()
}

func (self *CStyleCommentMatcher) Match(tz *Tokenizer) (Token, error) {
	if tz.Current() != '/' || tz.Peek(1) != '/' {
		return EmptyToken, nil
	}

	tz.ConsumeOne()
	tz.ConsumeOne()

	begin := tz.Index()

	for {

		if tz.Current() == '\n' || tz.Current() == '\r' || tz.Current() == 0 {
			break
		}

		tz.ConsumeOne()

	}

	return NewToken(self, tz, tz.StringRange(begin, tz.index), ""), nil
}

func NewCStyleCommentMatcher(id int) TokenMatcher {
	return &CStyleCommentMatcher{
		baseMatcher{id},
	}
}
