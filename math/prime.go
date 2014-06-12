package math

import gomath "math"

func IsPrime(n int) bool {
	sr := gomath.Sqrt(float64(n))
	for i := 2; i < int(sr); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
