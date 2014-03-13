/*
Problem 4: Largest Palindrome Product
A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.
Find the largest palindrome made from the product of two 3-digit numbers.
Answer: 906609
*/

package problems

func isPalindromeInt(num int) bool {
	return num == reverseInt(num)
}
func reverseInt(num int) (reverse int) {
	for num > 0 {
		lastDigit := num % 10
		reverse = reverse*10 + lastDigit
		num = num / 10
	}
	return reverse
}

func Euler004() int {
	var result, min int

	for i := 100; i < 1000; i++ {
		for j := 100; j < i; j++ {
			result = i * j
			if isPalindromeInt(result) && result > min {
				min = result
			}
		}
	}
	return min
}

/*
   func isPalindromeString(text string) bool {
   	return text == reverseString(text)
   }

   func reverseString(s string) string {
   	runes := []rune(s)
   	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
   		runes[i], runes[j] = runes[j], runes[i]
   	}
   	return string(runes)
   }
*/
