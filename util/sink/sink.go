package sink

import (
	"math"
	"math/big"
	"strings"

	"github.com/zolrath/euler/util/primes"
)

func GCD(a, b int) int {
	if a < b {
		a, b = b, a
	}
	if a%b == 0 {
		return b
	}
	return GCD(b, a%b)
}

func LCM(a, b int) int {
	return a * b / GCD(a, b)
}

func Factorial(n float64) float64 {
	result := float64(1)
	for i := float64(2); i <= n; i++ {
		result *= i
	}
	return result
}

func FactorialInt(n int) int {
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result
}

func FactorialBigInt(num int64) *big.Int {
	n, one, result := big.NewInt(num), big.NewInt(1), big.NewInt(1)
	for i := big.NewInt(2); i.Cmp(n) != 1; i = i.Add(i, one) {
		result = result.Mul(result, i)
	}
	return result
}

func SumBigDigits(n *big.Int) int {
	mod, sum, zero := big.NewInt(0), big.NewInt(0), big.NewInt(0)
	ten := big.NewInt(10)
	for n.Cmp(zero) != 0 {
		n.DivMod(n, ten, mod)
		sum.Add(sum, mod)
	}
	return int(sum.Int64())
}

func Int64Len(n int64) int64 {
	return int64(math.Log10(float64(n))) + 1
}

func IntLen(n int) int {
	length := 0
	for n > 0 {
		length++
		n /= 10
	}
	return length
}

func SplitDigits(n int) []int {
	nLen := IntLen(n)
	digits := make([]int, nLen)
	for nLen--; n > 0; nLen-- {
		dig := n % 10
		digits[nLen] = dig
		n /= 10
	}
	return digits
}

func JoinDigits(nums []int) int {
	num := 0
	for _, d := range nums {
		num = (num * 10) + d
	}
	return num
}

func GetIntRotations(in int) []int {
	inLen := IntLen(in)
	rotations := make([]int, inLen)
	for i := 0; i < inLen; i++ {
		rotations[i] = RotateInt(in, i)
	}
	return rotations
}

func GetStringRotations(in string) []string {
	rotations := make([]string, len(in))
	for i := 0; i < len(in); i++ {
		rotations[i] = RotateString(in, i)
	}
	return rotations
}

func RotateInt(in int, rotations int) int {
	inLen := int(math.Log10(float64(in)))
	powage := int(math.Pow10(inLen))
	for i := 0; i < rotations; i++ {
		head := in / powage
		tail := in % powage
		in = tail*10 + head
	}
	return in
}

func RotateString(in string, rotations int) string {
	input := strings.Split(in, "")
	for i := 0; i < rotations; i++ {
		input = append(input[1:], input[:1]...)

	}
	return strings.Join(input, "")
}

func ReverseInt(num int) (reverse int) {
	for num > 0 {
		lastDigit := num % 10
		reverse = reverse*10 + lastDigit
		num = num / 10
	}
	return reverse
}

func ReverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func IsPalindromeInt(num int) bool {
	return num == ReverseInt(num)
}

func IsPalindromeString(text string) bool {
	return text == ReverseString(text)
}

func ConcatInt(a, b int) int {
	l := IntLen(b)
	a *= int(math.Pow10(l))
	return a + b
}

// Prime Factorization to the rescue!
// σ(n) is the sum of divisors of a natural number.
// σ(p^a) = (p^a+1 − 1)/(p − 1).
// Sum of all proper divisors (not including the number itself)
func ProperDivisorSum(n int) int {
	factors := primes.PrimeFactors(n)
	sum := 1
	for p, e := range factors {
		sum *= (int(math.Pow(float64(p), float64(e+1))) - 1) / (p - 1)
	}
	return sum - n
}

// IsPandigital requires a number have all numbers up to len(n)
func IsPandigital(n int) bool {
	l := IntLen(n)
	nums := make([]int, l+1)
	for n > 0 {
		dig := n % 10
		if dig == 0 {
			l--
		}
		if dig > l || nums[dig] == 1 {
			return false
		}
		nums[dig] = 1
		n /= 10
	}
	return true
}

func numberOfTrailingZeros(i int64) int64 {
	zerosOnRightModLookup := []int64{32, 0, 1, 26, 2, 23, 27, 0, 3, 16, 24, 30, 28, 11, 0, 13, 4, 7, 17,
		0, 25, 22, 31, 15, 29, 10, 12, 6, 0, 21, 14, 9, 5, 20, 8, 19, 18}
	return zerosOnRightModLookup[(i&-i)%37]
}

func IsPerfectSquare(n int) bool {
	rt := int(math.Sqrt(float64(n)))
	return rt*rt == n
}
