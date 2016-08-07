package golexer

type Tokenizer struct {
	src   []rune
	index int
	line  int
}

func (self *Tokenizer) Current() rune {

	if self.EOF() {
		return 0
	}

	return self.src[self.index]
}

func (self *Tokenizer) Index() int {
	return self.index
}

func (self *Tokenizer) Line() int {
	return self.line
}

func (self *Tokenizer) Peek(offset int) rune {

	if self.index+offset >= len(self.src) {
		return 0
	}

	return self.src[self.index+offset]
}

func (self *Tokenizer) ConsumeOne() {

	self.index++
}

func (self *Tokenizer) ConsumeMulti(count int) {

	self.index += count
}

func (self *Tokenizer) EOF() bool {
	return self.index >= len(self.src)
}

func (self *Tokenizer) increaseLine() {
	self.line++
}

func (self *Tokenizer) StringRange(begin, end int) string {

	if begin < 0 {
		begin = 0
	}

	if end > len(self.src) {
		end = len(self.src)
	}

	return string(self.src[begin:end])
}

func NewTokenizer(s string) *Tokenizer {

	return &Tokenizer{
		src:  []rune(s),
		line: 1,
	}
}
