package types

import "github.com/eudaimathia/src/token"

// System has methods for adding/recognizing types currently in use.
type System struct {
}

// NewSystem initializes the type-system.
func NewSystem() *System {
	return &System{}
}

// Ur is a bottom-level type.
type Ur struct {
	name string
}

func (u *Ur) String() string {
	return u.name
}

// Ur returns itself.
func (u *Ur) Ur() *Ur {
	return u
}

// From returns nil.
func (u *Ur) From() T {
	return nil
}

// To returns nil.
func (u *Ur) To() T {
	return nil
}

// T is the interface used to represent a type.
// The T values returned from System will have correct behavior with respect to ==.
type T interface {
	// Ur type, or nil.
	Ur() *Ur
	// From returns the domain type of this Fn, or nil if Ur.
	From() T
	// To returns the codomain type of this Fn, or nil if Ur.
	To() T
}

// NewUr with the given name.
func NewUr(name string) *Ur {
	return &Ur{
		name,
	}
}

// AddUr to the System.
func (s *System) AddUr(ur *Ur) {

}

// Parse a type from the given token stream.
func (s *System) Parse(ts *token.Stream) T {
	return nil
}
