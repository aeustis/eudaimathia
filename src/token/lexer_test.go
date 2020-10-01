package token_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/eudaimathia/src/token"
)

func TestPeek(t *testing.T) {
	s := token.NewStream("a-)")
	assert.Equal(t, token.Token{token.Alpha, "a"}, s.Next())

	hyphen := token.Token{token.Symbol, "-"}
	assert.Equal(t, hyphen, s.Peek())
	assert.Equal(t, hyphen, s.Peek())
	assert.Equal(t, hyphen, s.Next())

	paren := token.Token{token.Symbol, ")"}
	assert.Equal(t, paren, s.Peek())
	assert.Equal(t, paren, s.Next())

	eof := token.Token{token.EOF, ""}
	assert.Equal(t, eof, s.Peek())
	assert.Equal(t, eof, s.Next())
	assert.Equal(t, eof, s.Next())
}

func TestStream(t *testing.T) {
	tests := []struct {
		input    string
		expected []string
	}{
		{
			input:    "a",
			expected: []string{"a"},
		},
		{
			input:    "foo \n \r\n\tbar",
			expected: []string{"foo", "bar"},
		},
		{
			input:    "+-",
			expected: []string{"+-"},
		},
		{
			input:    "+ -",
			expected: []string{"+", "-"},
		},
		{
			input:    "a/B",
			expected: []string{"a", "/", "B"},
		},
		{
			input:    "((",
			expected: []string{"(", "("},
		},
		{
			input:    "))--",
			expected: []string{")", ")", "--"},
		},
		{
			input:    "foo1+_+2",
			expected: []string{"foo1", "+", "_", "+", "2"},
		},
		{
			input:    "a_1@15",
			expected: []string{"a_1", "@", "15"},
		},
		{
			input:    "019aAbBzZ_",
			expected: []string{"019aAbBzZ_"},
		},
		{
			input:    "A->B",
			expected: []string{"A", "->", "B"},
		},
		{
			input:    "A->(B->C)",
			expected: []string{"A", "->", "(", "B", "->", "C", ")"},
		},
	}
	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			s := token.NewStream(test.input)
			for _, exp := range test.expected {
				assert.Equal(t, exp, s.Next().V)
			}
			eof := token.Token{token.EOF, ""}
			assert.Equal(t, eof, s.Next())
		})
	}
}
