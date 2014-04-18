/*
Problem 42: Coded Triangle Numbers
The nth term of the sequence of triangle numbers is given by, tn = ½n(n+1); so the first ten triangle numbers are:
1, 3, 6, 10, 15, 21, 28, 36, 45, 55, ...
By converting each letter in a word to a number corresponding to its alphabetical position and adding these values we form a word value.
For example, the word value for SKY is 19 + 11 + 25 = 55 = t₁₀. If the word value is a triangle number then we shall call the word a triangle word.

Using words.txt, a 16K text file containing nearly two-thousand common English words, how many are triangle words?
*/

package problems

import (
	"bufio"
	"math"
	"os"
)

const ANSWER_042 = 162

func isTriangle(num int) bool {
	n := int(math.Sqrt(float64(num * 2)))
	return num == n*(n+1)/2
}

func alphaScore(word string) int {
	var score int
	for _, l := range word {
		if l == '"' || l == ',' {
			continue
		}
		score += int(byte(l) - '@')
	}
	return score
}

func Euler042() int {
	count := 0
	f, _ := os.Open("problems/words.txt")
	defer f.Close()
	r := bufio.NewReader(f)
	for {
		line, err := r.ReadString(',')
		if err != nil {
			break
		}
		if isTriangle(alphaScore(line)) {
			count++
		}
	}

	return count
}
