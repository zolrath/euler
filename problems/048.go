/*
Problem 48: Self Powers
The series, 1¹ + 2² + 3³ + ... + 10¹⁰ = 10405071317

Find the last ten digits of the series, 1¹ + 2² + 3³ + ... + 1000¹⁰⁰⁰
*/

package problems

const ANSWER_048 = 9110846700

func Euler048() int {
	totalDigits := int64(0)
	for i := int64(1); i <= 1000; i++ {
		posDigits := int64(1)
		for j := int64(1); j <= i; j++ {
			posDigits = (posDigits * i) % 100000000000
		}
		totalDigits = (totalDigits + posDigits) % 10000000000
	}
	return int(totalDigits)
}
