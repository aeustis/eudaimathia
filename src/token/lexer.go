package token

// Stream with Peek and Next methods.
// Empty string is used for EOF.
type Stream struct {
	toks chan string
	cur  string
}

// NewStream tokenizes the input string.
func NewStream(input string) *Stream {
	toks := make(chan string, 100)
	go lexAll(input, toks)
	return &Stream{
		toks: toks,
		cur:  <-toks,
	}
}

// Peek the next token without consuming it.
func (s *Stream) Peek() string {
	return s.cur
}

// Next consumes and returns the next token in the stream.
func (s *Stream) Next() string {
	cur := s.cur
	s.cur = <-s.toks
	return cur
}

// rune types
const (
	white = iota
	alpha
	symbol
	special
	eof
)

func lexAll(input string, output chan<- string) {
	var tokStart, i, state int
	for {
		// get type of current rune
		curType := 0
		if i == len(input) {
			curType = eof
		} else {
			curType = runeType(input[i])
		}

		// emit currrent token if we hit a type boundary
		if curType != state {
			if state == alpha || state == symbol {
				output <- input[tokStart:i]
			}
			tokStart = i
		}

		switch curType {
		case eof:
			close(output)
			return
		case special:
			output <- input[i : i+1]
		}
		state = curType
		i++
	}
}

func runeType(c byte) int {
	if c == '_' || (c >= '0' && c <= '9') || (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') {
		return alpha
	}
	switch c {
	case ' ', '\t', '\r', '\n':
		return white
	case '(', ')':
		return special
	default:
		return symbol
	}
}
