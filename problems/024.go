/*
Problem 24: Lexicographic Permutations
A permutation is an ordered arrangement of objects. For example, 3124 is one
possible permutation of the digits 1, 2, 3 and 4. If all of the permutations are
listed numerically or alphabetically, we call it lexicographic order.
The lexicographic permutations of 0, 1 and 2 are:
012   021   102   120   201   210
What is the millionth lexicographic permutation of the digits 0, 1, 2, 3, 4, 5, 6, 7, 8 and 9?
Answer: 2783915460
*/

package problems

func lexiPermute(nums []byte, perms int) int {
	for p := 0; p < perms-1; p++ {
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
		permutation = (permutation * 10) + int(v)
	}
	return permutation
}

func Euler024() int {
	milperm := lexiPermute([]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, 1000000)
	return milperm
}
