package suspicious

// Version is just that
const Version = "v1.0.0"

// El is
type El = uint64

// S is
type S []El

// D is
func (b1 S) D(b2 S) int {
	return 0
}

// I is
type I struct {
	x int
	y int
}

// SE is
type SE interface {
	N() ([]string, uint8)
}

// SI is
type SI interface {
	N() string
}

// SIType is
type SIType struct {
	s string
}

// N is
func (sit SIType) N() string {
	return sit.s
}

// NewI generates I
func NewI(x, y int) I {
	return I{x, y}
}

// VEl is
type VEl = int32

// V is
type V = []VEl

// CS returns
func (i *I) CS(f SI, b S, s V) S {
	return b
}

// SIt is
type SIt struct {
	x int
	y int
	i *I
}

// NewSIt returns a structure
func (i *I) NewSIt(x, y int) *SIt {

	si := SIt{x, y, i}

	return &si
}

// N is
func (si *SIt) N(b S) S {
	return b
}
