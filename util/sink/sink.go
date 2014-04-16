package sink

import (
	"math"
	"math/big"
	"sort"
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

func IntLength(n int) int {
	return int(math.Log10(float64(n))) + 1
}

func GetDigits(n int) []int {
	nLen := IntLength(n)
	digits := make([]int, nLen)
	for nLen--; n > 0; nLen-- {
		dig := n % 10
		digits[nLen] = dig
		n /= 10
	}
	return digits
}

func LexiStringPermute(in string, perms int) string {
	input := strings.Split(in, "")
	for p := 0; p < perms; p++ {
		var pivot int
		end := len(input) - 1
		// Find pivot, cell in which the cell preceeding it is higher in value.
		for i := end; i > 0; i-- {
			if input[i-1] < input[i] {
				pivot = i - 1
				break
			}
		}
		// Find first number before pivot which is higher than it.
		for i := end; i > pivot; i-- {
			if input[i] > input[pivot] {
				input[i], input[pivot] = input[pivot], input[i]
				pivot++
				break
			}
		}
		for i := end; pivot < i; {
			input[i], input[pivot] = input[pivot], input[i]
			pivot++
			i--
		}
	}
	return strings.Join(input, "")
}

func LexiIntPermute(num int, perms int) int {
	nums := GetDigits(num)
	sort.Ints(nums)
	for p := 0; p < perms; p++ {
		var pivot int
		end := len(nums) - 1
		// Find pivot, cell in which the cell preceeding it is higher in value.
		for i := end; i > 0; i-- {
			if nums[i-1] < nums[i] {
				pivot = i - 1
				break
			}
		}
		// Find first number before pivot which is higher than it.
		for i := end; i > pivot; i-- {
			if nums[i] > nums[pivot] {
				nums[i], nums[pivot] = nums[pivot], nums[i]
				pivot++
				break
			}
		}
		for i := end; pivot < i; {
			nums[i], nums[pivot] = nums[pivot], nums[i]
			pivot++
			i--
		}
	}
	permutation := 0
	for _, v := range nums {
		permutation = (permutation * 10) + v
	}
	return permutation
}

func GetStringPermutations(in string) []string {
	permCount := FactorialInt(len(in))
	permutations := make([]string, permCount)
	for i := 0; i < permCount; i++ {
		permutations[i] = LexiStringPermute(in, i)
	}
	return permutations
}

func GetIntPermutations(in int) []int {
	permCount := FactorialInt(IntLength(in))
	permutations := make([]int, permCount)
	for i := 0; i < permCount; i++ {
		permutations[i] = LexiIntPermute(in, i)
	}
	return permutations
}

func GetIntRotations(in int) []int {
	inLen := IntLength(in)
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
