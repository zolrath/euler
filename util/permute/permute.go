package permute

import (
	"sort"
	"strings"

	"github.com/zolrath/euler/util/sink"
)

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
	nums := sink.SplitDigits(num)
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
	return sink.JoinDigits(nums)
}

func Strings(in string) []string {
	permCount := sink.FactorialInt(len(in))
	permutations := make([]string, permCount)
	for i := 0; i < permCount; i++ {
		permutations[i] = LexiStringPermute(in, i)
	}
	return permutations
}

func Ints(in int) []int {
	permCount := sink.FactorialInt(sink.IntLen(in))
	permutations := make([]int, permCount)
	for i := 0; i < permCount; i++ {
		permutations[i] = LexiIntPermute(in, i)
	}
	return permutations
}

func UniqInts(in int) []int {
	permCount := sink.FactorialInt(sink.IntLen(in))
	permutationSet := make(map[int]bool)
	for i := 0; i < permCount; i++ {
		permutationSet[LexiIntPermute(in, i)] = true
	}
	permutations := make([]int, 0, len(permutationSet))
	for k, _ := range permutationSet {
		permutations = append(permutations, k)
	}
	sort.Ints(permutations)
	return permutations
}
