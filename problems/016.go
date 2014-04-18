/*
Problem 16: Power Digit Sum
2^15 = 32768 and the sum of its digits is 3 + 2 + 7 + 6 + 8 = 26.

What is the sum of the digits of the number 2^1000?
*/

package problems

import (
	"math/big"

	"github.com/zolrath/euler/util/sink"
)

const ANSWER_016 = 1366

func Euler016() int {
	num := big.NewInt(0).Exp(big.NewInt(2), big.NewInt(1000), nil)
	return sink.SumBigDigits(num)
}

// func sumBigDigits(n *big.Int) int {
// 	mod, sum, zero := big.NewInt(0), big.NewInt(0), big.NewInt(0)
// 	ten := big.NewInt(10)
// 	for n.Cmp(zero) != 0 {
// 		n.DivMod(n, ten, mod)
// 		sum.Add(sum, mod)
// 	}
// 	return int(sum.Int64())
// }
