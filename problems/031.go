/*
Problem 31: Coin Sums
In England the currency is made up of pound, £, and pence, p, and there are eight coins in general circulation:
1p, 2p, 5p, 10p, 20p, 50p, £1 (100p) and £2 (200p).
It is possible to make £2 in the following way:
1×£1 + 1×50p + 2×20p + 1×5p + 1×2p + 3×1p
How many different ways can £2 be made using any number of coins?
*/

package problems

// An = 1p
// Bn = 1p+2p
// Cn = 1p+2p+5p
// Dn = 1p+2n+5p+10p
// En = 1p+2n+5p+10p+20p
// Fn = 1p+2n+5p+10p+20p+50p
// Gn = 1p+2n+5p+10p+20p+50p+100p
// Hn = 1p+2n+5p+10p+20p+50p+100p+200p
// MY WORD ENGLAND, GET IT TOGETHER OL CHAP

const ANSWER_031 = 73682

func Euler031() int {
	britCash := []int{1, 2, 5, 10, 20, 50, 100, 200}
	combinations := [201]int{1}
	for _, v := range britCash {
		for j := v; j <= 200; j++ {
			combinations[j] += combinations[j-v]
		}
	}
	return combinations[200]
}
