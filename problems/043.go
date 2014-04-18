/*
Problem 43: Sub-String Divisibility
The number, 1406357289, is a 0 to 9 pandigital number because it is made up of each of the
digits 0 to 9 in some order, but it also has a rather interesting sub-string divisibility property.

Let d1 be the 1st digit, d2 be the 2nd digit, and so on. In this way, we note the following:

d2d3d4=406 is divisible by 2
d3d4d5=063 is divisible by 3
d4d5d6=635 is divisible by 5
d5d6d7=357 is divisible by 7
d6d7d8=572 is divisible by 11
d7d8d9=728 is divisible by 13
d8d9d10=289 is divisible by 17
Find the sum of all 0 to 9 pandigital numbers with this property.
*/

package problems

import "github.com/zolrath/euler/util/sink"

const ANSWER_043 = 16695334890

func concatInts(a, b, c int) int {
	return (a * 100) + (b * 10) + c
}

//TODO: Use a less naive method. Faster to calc multiples of 17, 13, .. ?
func nextDivisiblePermute(num int, divisors []int) (int, bool) {
	nums := sink.SplitDigits(num)
	isDivisible := false
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

	for i := 1; i <= len(nums)-3; i++ {
		subdigit := concatInts(nums[i], nums[i+1], nums[i+2])
		if subdigit%divisors[i-1] != 0 {
			isDivisible = false
			break
		}
		isDivisible = true
	}

	result := 0
	for _, digit := range nums {
		result = (result * 10) + digit
	}
	if result < num {
		return -1, false
	}
	return result, isDivisible
}

func Euler043() int {
	sum := 0
	isDiv := false
	divisors := []int{2, 3, 5, 7, 11, 13, 17}
	for i := 1234567890; i != -1; {
		if i, isDiv = nextDivisiblePermute(i, divisors); isDiv {
			sum += i
		}
	}
	return sum

}
