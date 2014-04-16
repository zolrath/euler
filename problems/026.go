/*
Problem 26: Reciprocal Cycles
A unit fraction contains 1 in the numerator. The decimal representation of the unit fractions with denominators 2 to 10 are given:

1/2	= 	0.5
1/3	= 	0.(3)
1/4	= 	0.25
1/5	= 	0.2
1/6	= 	0.1(6)
1/7	= 	0.(142857)
1/8	= 	0.125
1/9	= 	0.(1)
1/10	= 	0.1
Where 0.1(6) means 0.166666..., and has a 1-digit recurring cycle. It can be seen that 1/7 has a 6-digit recurring cycle.

Find the value of d < 1000 for which 1/d contains the longest recurring cycle in its decimal fraction part.
*/

package problems

const ANSWER_026 = 983

func findCyclicCount(n int) int {
	rmap := map[int]int{}
	count := 0
	for rem := 1; ; {
		rem = rem % n
		if rem == 0 {
			break
		}
		if rmap[rem] == 1 {
			break
		}
		count++
		if rmap[rem] == 0 {
			rmap[rem] = 1
		}
		rem = 10 * rem
	}
	return count
}

func Euler026() int {
	longest := 0
	longestnum := 0
	for i := 1; i < 1000; i++ {
		count := findCyclicCount(i)
		if count > longest {
			longest = count
			longestnum = i
		}
	}
	return longestnum
	return 1
}
