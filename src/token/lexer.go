package token

// Type of a token.
type Type int

// Token types
const (
	EOF Type = iota
	Alpha
	Symbol
	white
	special
)

// Token consisting of a type and a value
type Token struct {
	Type
	V string
}

func (t Token) String() string {
	if t.Type == EOF {
		return "EOF"
	}
	return t.V
}

// Stream with Peek and Next methods.
// Empty string is used for EOF.
type Stream struct {
	toks chan Token
	cur  Token
}

// NewStream tokenizes the input string.
func NewStream(input string) *Stream {
	toks := make(chan Token, 100)
	go lexAll(input, toks)
	return &Stream{
		toks: toks,
		cur:  <-toks,
	}
}

// Peek the next token without consuming it.
func (s *Stream) Peek() Token {
	return s.cur
}

// Next consumes and returns the next token in the stream.
func (s *Stream) Next() Token {
	cur := s.cur
	s.cur = <-s.toks
	return cur
}

func lexAll(input string, output chan<- Token) {
	var tokStart, i int
	var state Type
	for {
		// get type of current rune
		var curType Type
		if i == len(input) {
			curType = EOF
		} else {
			curType = runeType(input[i])
		}

		// emit currrent token if we hit a type boundary
		if curType != state {
			if state == Alpha || state == Symbol {
				output <- Token{state, input[tokStart:i]}
			}
			tokStart = i
		}

		switch curType {
		case EOF:
			close(output)
			return
		case special:
			output <- Token{Symbol, input[i : i+1]}
		}
		state = curType
		i++
	}
}

// runeType classifies the given byte as one of several types.
func runeType(c byte) Type {
	if c == '_' || (c >= '0' && c <= '9') || (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') {
		return Alpha
	}
	switch c {
	case ' ', '\t', '\r', '\n':
		return white
	case '(', ')':
		return special
	default:
		return Symbol
	}
}
