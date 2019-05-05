package recursive

import (
	"errors"
	"math"
)

// Resolve1st resolve the nth item of F(n) = a * f(n-1) + c, where f1 is the initialization
func Resolve1st(f1, a, c, n int64) (int64, error) {
	if n <= 0 {
		return 0, errors.New("Invalid Argument(nth)")
	} else if n == 1 {
		return f1, nil
	}

	if a == 1 {
		return f1 + (n)*c, nil
	}

	k := n - 1
	a2kth := int64(math.Pow(float64(a), float64(k)))

	return f1*a2kth + c*((a2kth-1)/(a-1)), nil
}
