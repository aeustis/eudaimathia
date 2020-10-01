package types

import (
	"strings"

	"github.com/eudaimathia/src/errs"
	"github.com/eudaimathia/src/reversevariant"
	"github.com/eudaimathia/src/token"
)

// Parse a type from the given token stream.
func (s *System) Parse(ts *token.Stream) T {
	var canon strings.Builder
	s.parse(&canon, ts)
	return s.recognize(canon.String()).urify()
}

const mapsTo = "->"

// recusively parse a token stream as a type, writing its canonical string as we go.
// Return value is the length of the canonical string written.
func (s *System) parse(w *strings.Builder, ts *token.Stream) int {
	tok := ts.Next()
	n := 0
	switch tok.V {
	case "(":
		n = s.parse(w, ts)
		if tok = ts.Next(); tok.V != ")" { // Consume right paren
			panic(errs.ParseErrorf("expected ')', got %v", tok))
		}
	default:
		m, ok := s.urMap[tok.V]
		if !ok {
			panic(errs.ParseErrorf("expected type name, got %v", tok))
		}
		n = reversevariant.WriteUint32(w, m)
	}
	tok = ts.Peek()
	if tok.V != mapsTo {
		return n
	}
	ts.Next()
	// Greedy parse for right associativity
	return n + s.parse(w, ts) + reversevariant.WriteUint32(w, uint32(n))
}

// recursively recognize a canonical string, meaning that we either identify or create
// its corresponding node in the type hierarchy.
func (s *System) recognize(in string) *fun {
	if fun := s.funMap[in]; fun != nil {
		return fun
	}
	m, k := reversevariant.ReadUint32(in)
	if k == len(in) {
		// m is the index of an ur-type
		f := &fun{ur: s.urs[m]}
		s.funMap[in] = f
		return f
	}
	// Otherwise m tells us how many bytes of "from" we have in s
	f := &fun{
		from: s.recognize(in[:m]),
		to:   s.recognize(in[m : len(in)-k]),
	}
	s.funMap[in] = f
	return f
}
