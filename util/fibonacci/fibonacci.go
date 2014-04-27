// Package fibonacci provides various methods of generating the fibonacci sequence.
package fibonacci

import (
	"math"
	"math/big"
)

// To creates a slice containing n fibonacci numbers.
func To(limit int) []int64 {
	fibSeq := []int64{0, 1}
	for i := 2; i < limit; i++ {
		newNum := fibSeq[i-1] + fibSeq[i-2]
		fibSeq = append(fibSeq, newNum)
	}
	return fibSeq
}

func Nth(n float64) int {
	return int(0.25 + (math.Pow(1+math.Sqrt(5), n)-math.Pow(1-math.Sqrt(5), n))/(math.Pow(2, n)*math.Sqrt(5)))
}

// IntGenerator returns a function which returns the next fibonacci number when called.
func IntGenerator() func() int {
	a, b, ans := 0, 1, 0
	return func() int {
		a, b, ans = b, a+b, a
		return ans
	}
}

// BigIntGenerator returns a function which returns the next fibonacci number when called.
func BigIntGenerator() func() *big.Int {
	a, b, ans := big.NewInt(0), big.NewInt(1), big.NewInt(0)
	return func() *big.Int {
		ans.Set(a)
		a.Set(b)
		b.Add(ans, b)
		return ans
	}
}
