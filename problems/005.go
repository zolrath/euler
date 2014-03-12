/*
Problem 5: Smallest Multiple
2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.
What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?
Answer: 232792560
*/

package problems

func gcd(a, b int) int {
	if a < b {
		a, b = b, a
	}
	if a%b == 0 {
		return b
	}
	return gcd(b, a%b)
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

func multilcm(nums []int) int {
	m := nums[0]
	for i := 1; i < len(nums)-1; i++ {
		m = lcm(m, nums[i])
	}
	return m
}

func Euler005() int {
	return multilcm([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20})
}
