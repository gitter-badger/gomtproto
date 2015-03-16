package mtproto

import (
	"math"
)

func IsPrime(x int64) bool {
	if x%2 == 0 {
		return false
	}

	s := int64(math.Sqrt(float64(x)))
	if s%2 == 0 {
		s += 1
	}

	for i := int64(3); i <= s; i += 2 {
		if x%i == 0 {
			return false
		}
	}

	return true
}

func Factorize(x int64) int64 {
	s := int64(math.Sqrt(float64(x)))
	if s%2 == 0 {
		s += 1
	}

	fx := float64(x)
	for i := s; i > 1; i -= 2 {
		r := fx / float64(i)
		if math.Mod(r, 1.0) == 0 && IsPrime(i) && IsPrime(int64(r)) {
			return i
		}
	}

	return 1
}
