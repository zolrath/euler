/*
Problem 40: Champernowne's Constant
An irrational decimal fraction is created by concatenating the positive integers:
0.123456789101112131415161718192021...
It can be seen that the 12th digit of the fractional part is 1.

If dn represents the nth digit of the fractional part, find the value of the following expression.
d₁ × d₁₀ × d₁₀₀ × d₁₀₀₀ × d₁₀₀₀₀ × d₁₀₀₀₀₀ × d₁₀₀₀₀₀₀
*/

package problems

import (
	"strconv"
	"strings"

	"github.com/zolrath/euler/util/sink"
)

const ANSWER_040 = 210

func createConstant(limit int) string {
	champ := make([]string, 0)
	strLen := 0
	for i := 1; strLen < limit; i++ {
		strLen += sink.IntLen(i)
		champ = append(champ, strconv.Itoa(i))
	}
	return strings.Join(champ, "")
}

func getDigit(n int, str string) int {
	result, _ := strconv.Atoi(string(str[n-1]))
	return result
}

func Euler040() int {
	champ := createConstant(1000000)
	result := 1
	for i := 1; i <= 1000000; i *= 10 {
		result *= getDigit(i, champ)

	}
	return result
}
