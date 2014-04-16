/*
Problem 25: 1000-Digit Fibonacci Number
The Fibonacci sequence is defined by the recurrence relation:
Fn = Fn−1 + Fn−2, where F1 = 1 and F2 = 1.
Hence the first 12 terms will be:
F1 = 1
F2 = 1
F3 = 2
F4 = 3
F5 = 5
F6 = 8
F7 = 13
F8 = 21
F9 = 34
F10 = 55
F11 = 89
F12 = 144
The 12th term, F12, is the first term to contain three digits.
What is the first term in the Fibonacci sequence to contain 1000 digits?
Answer: 4782
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
