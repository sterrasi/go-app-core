package core

import (
	"math"
	"strings"

	"golang.org/x/exp/constraints"
)

type Number interface {
	constraints.Integer | constraints.Float
}

// Normalize will convert the given string to lower case and trim any leading
// and trailing whitespace.
func Normalize(value string) string {
	if value == "" {
		return ""
	}
	nml := strings.ToLower(value)
	return strings.TrimSpace(nml)
}

// GetMapValues returns a slice of the specified maps values
func GetMapValues[K comparable, V comparable](m map[K]V) []V {
	r := make([]V, 0, len(m))
	for _, val := range m {
		r = append(r, val)
	}
	return r
}

// IsWithinRange returns true if the two values fall within the given rangeVal
func IsWithinRange(v1 any, v2 any, rangeVal float64) bool {
	f1 := ToFloat64(v1)
	f2 := ToFloat64(v2)

	if f1 == nil || f2 == nil {
		return false
	}
	return math.Abs(*f1-*f2) < rangeVal
}

type Tuple[T, U any] struct {
	First  T
	Second U
}

type Triple[T, U, V any] struct {
	First  T
	Second U
	Third  V
}

// ToFloat64 casts the given value to a float64 and returns a pointer to it. If the value is not numeric then
// nil is returned
func ToFloat64(value any) *float64 {
	var result float64

	switch v := value.(type) {
	case float32:
		result = float64(v)
	case float64:
		result = v
	case int8:
		result = float64(v)
	case int16:
		result = float64(v)
	case int32:
		result = float64(v)
	case int64:
		result = float64(v)
	case uint8:
		result = float64(v)
	case uint16:
		result = float64(v)
	case uint32:
		result = float64(v)
	case uint64:
		result = float64(v)
	case uintptr:
		result = float64(v)

	default:
		return nil
	}
	return &result
}
