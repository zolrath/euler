/*
Problem 17: Number Letter Counts
If the numbers 1 to 5 are written out in words: one, two, three, four, five, then there are 3 + 3 + 5 + 4 + 4 = 19 letters used in total.
If all the numbers from 1 to 1000 (one thousand) inclusive were written out in words, how many letters would be used?
NOTE: Do not count spaces or hyphens. For example, 342 (three hundred and forty-two) contains 23 letters and
      115 (one hundred and fifteen) contains 20 letters. The use of "and" when writing out numbers is in compliance with British usage.
Answer: 21124
*/

package problems

const ANSWER_017 = 21124

var onesCount, tensCount []int

func numLetterCount(n int) int {
	var letters int

	hundreds := n / 100
	tens := (n - (hundreds * 100)) / 10
	ones := n % 10

	if n > 100 && tens > 0 || n > 100 && ones > 0 {
		letters += 3
	}
	if hundreds > 0 {
		letters += onesCount[hundreds]
		letters += 7
	}
	if tens > 1 {
		letters += tensCount[tens]
	}
	if tens == 1 {
		ones += 10
	}
	letters += onesCount[ones]
	return letters
}

func Euler017() int {
	onesCount = []int{0, 3, 3, 5, 4, 4, 3, 5, 5, 4, 3, 6, 6, 8, 8, 7, 7, 9, 8, 8}
	tensCount = []int{0, 3, 6, 6, 5, 5, 5, 7, 6, 6}
	// onesList = []string{"", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten", "eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen", "seventeen", "eighteen", "nineteen"}
	// tensList = []string{"", "ten", "twenty", "thirty", "forty", "fifty", "sixty", "seventy", "eighty", "ninety"}
	// hundList = []string{"", "hundred", "thousand", "million", "billion", "trillion"}

	letterCount := 0
	for i := 1; i < 1000; i++ {
		letterCount += numLetterCount(i)
	}
	// add 11 characters for 'one thousand'
	return letterCount + 11
}
