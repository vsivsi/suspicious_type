package suspicious

//El is the element type
type El = uint64

// S is the variable length slice of El type
type S []El

// ElV is the vector element type
type ElV = int32

// V is the variable length vector type
type V = []ElV

// I is a struct
type I struct {
	x int
	y int
}

// NewI returns an I
func NewI() I {
	return I{}
}

// F is a function
func (i I) F(s S, v V) (S, V) {
	return s, v
}

// F2 is another function
func (i I) F2(s S, v V) S {
	return s
}
