/*
Problem 25: 1000-Digit Fibonacci Number
The Fibonacci sequence is defined by the recurrence relation:
Fᵢ = Fᵢ₋₁ + Fᵢ₋₂ where F₁ = 1 and F₂ = 1.

Hence the first 12 terms will be:
F₁  = 1
F₂  = 1
F₃  = 2
F₄  = 3
F₅  = 5
F₆  = 8
F₇  = 13
F₈  = 21
F₉  = 34
F₁₀ = 55
F₁₁ = 89
F₁₂ = 144
The 12th term, F₁₂ is the first term to contain three digits.

What is the first term in the Fibonacci sequence to contain 1000 digits?
*/

package problems

import (
	"math/big"

	"github.com/zolrath/euler/util/fibonacci"
)

const ANSWER_025 = 4782

func Euler025() int {
	fib := fibonacci.BigIntGenerator()
	var num *big.Int
	for i := 0; ; i++ {
		num = fib()
		if len(num.String()) == 1000 {
			return i
		}
	}
	return -1
}

// func bigIntFibonacci() func() *big.Int {
// 	a, b, ans, tmp := big.NewInt(0), big.NewInt(1), big.NewInt(0), big.NewInt(0)
// 	return func() *big.Int {
// 		tmp.Set(b)
// 		ans.Set(a)
// 		b.Add(a, b)
// 		a.Set(tmp)
// 		return ans
// 	}
// }
