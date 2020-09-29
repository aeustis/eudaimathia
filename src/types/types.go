package types

// System has methods for adding/recognizing types currently in use.
type System struct {
}

// Ur is a top-level type.
type Ur struct {
	name string
}

// NewUr with the given name.
func NewUr(name string) *Ur {
	return &Ur{
		name,
	}
}

func (u *Ur) String() string {
	return u.name
}

// AddUr to the System.
func (s *System) AddUr(ur *Ur) {

}

// 
