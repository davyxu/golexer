package golexer

// 行结束
type LineEndMatcher struct {
	baseMatcher
}

func (self *LineEndMatcher) Match(tz *Tokenizer) (*Token, error) {

	var count int
	for {

		c := tz.Peek(count)

		if c == '\n' {
			tz.increaseLine()
		} else if c == '\r' {

		} else {
			break
		}

		count++

	}

	if count == 0 {
		return nil, nil
	}

	tz.ConsumeMulti(count)

	return NewToken(self, tz, "\r", ""), nil
}

func NewLineEndMatcher(id int) TokenMatcher {
	return &LineEndMatcher{
		baseMatcher{id},
	}
}
