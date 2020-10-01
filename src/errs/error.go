package errs

import "fmt"

// ParseError is an error that occurs during parsing or tokenizing.
type ParseError string

func (e ParseError) Error() string {
	return string(e)
}

// ParseErrorf constructs a ParseError from a format string and args
func ParseErrorf(f string, args ...interface{}) ParseError {
	return ParseError("Syntax error: " + fmt.Sprintf(f, args...))
}
