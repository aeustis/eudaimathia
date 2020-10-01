package types

// System has methods for adding/recognizing types currently in use.
type System struct {
	urs   []*Ur
	urMap map[string]uint32
	funMap map[string]*fun
}

// NewSystem initializes the type-system.
func NewSystem() *System {
	return &System{
		urMap: make(map[string]uint32),
		funMap: make(map[string]*fun),
	}
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

type fun struct {
	from, to *fun
	ur *Ur
}

func (f *fun) Ur() *Ur {
	return f.ur
}

func (f *fun) From() T {
	return f.from.urify()
}

func (f *fun) To() T {
	return f.to.urify()
}

func (f *fun) urify() T {
	if f == nil {
		return nil
	}
	if f.ur != nil {
		return f.ur
	}
	return f
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

// NewUr with the given name.  Caller must ensure that name is alphanumeric.
func NewUr(name string) *Ur {
	return &Ur{
		name,
	}
}

// AddUr to the System.
func (s *System) AddUr(ur *Ur) {
	name := ur.String()
	if _, exists := s.urMap[name]; exists {
		panic("Type already defined: " + name)
	}
	s.urMap[name] = uint32(len(s.urs))
	s.urs = append(s.urs, ur)
}
