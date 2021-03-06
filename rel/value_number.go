package rel

import (
	"reflect"
	"strconv"

	"github.com/arr-ai/hash"
)

// Number is a number.
type Number float64

// NewNumber returns a Number for the given number.
func NewNumber(n float64) Number {
	return Number(n)
}

// Float64 returns the value of the number.
func (n Number) Float64() float64 {
	return float64(n)
}

// Hash computes a hash for a Number.
func (n Number) Hash(seed uintptr) uintptr {
	return hash.Float64(float64(n), seed)
}

// Equal tests two Values for equality. Any other type returns false.
func (n Number) Equal(v interface{}) bool {
	if b, ok := v.(Number); ok {
		return n == b
	}
	return false
}

// String returns a string representation of a Number.
func (n Number) String() string {
	return strconv.FormatFloat(float64(n), 'G', -1, 64)
}

// Eval returns the number.
func (n Number) Eval(local, global Scope) (Value, error) {
	return n, nil
}

var numberKind = registerKind(100, reflect.TypeOf(Number(0)))

// Kind returns a number that is unique for each major kind of Value.
func (n Number) Kind() int {
	return numberKind
}

// Bool returns true iff the tuple has attributes.
func (n Number) Bool() bool {
	return n != 0
}

// Less returns true iff v is not a number or n < v.
func (n Number) Less(v Value) bool {
	if n.Kind() != v.Kind() {
		return n.Kind() < v.Kind()
	}
	return n < v.(Number)
}

// Negate returns -n.
func (n Number) Negate() Value {
	if !n.Bool() {
		return n
	}
	return NewNumber(-float64(n))
}

// Export exports a Number.
func (n Number) Export() interface{} {
	return n.Float64()
}
