/*
Problem 33: Digit Canceling Fractions
The fraction 49/98 is a curious fraction, as an inexperienced mathematician in attempting to
simplify it may incorrectly believe that 49/98 = 4/8, which is correct, is obtained by cancelling the 9s.
We shall consider fractions like, 30/50 = 3/5, to be trivial examples.
There are exactly four non-trivial examples of this type of fraction, less than one in value,
and containing two digits in the numerator and denominator.

If the product of these four fractions is given in its lowest common terms, find the value of the denominator.
*/

package problems

import (
	"strconv"
	"strings"

	"github.com/zolrath/euler/util/sink"
)

const ANSWER_033 = 100

// TODO: Math this up.
// People in the thread have done this in a much more mathy/elegant manner.

func cancelDigits(n, d int) (int, int) {
	// Get individual components of the two numbers.
	n1, n2 := n/10, n%10
	d1, d2 := d/10, d%10

	// If numbers %10 == 0 this would be a trivial example.
	if n2 == 0 && d2 == 0 {
		return 0, 0
	}

	// Find duplicated number in both numerator and denominator.
	cancel := 0
	if n1 == d1 || n1 == d2 {
		cancel = n1
	} else if n2 == d1 || n2 == d2 {
		cancel = n2
	} else {
		return 0, 0
	}

	// Convert numbers to strings
	sn := strconv.Itoa(n)
	sd := strconv.Itoa(d)
	sc := strconv.Itoa(cancel)

	// Remove the number shared by the numerator and denominator.
	srn := strings.Replace(sn, sc, "", 1)
	srd := strings.Replace(sd, sc, "", 1)

	// Convert strings back to numbers
	rn, _ := strconv.Atoi(srn)
	rd, _ := strconv.Atoi(srd)

	// Check to see if reduced form == original fraction.
	if float64(rn)/float64(rd) == float64(n)/float64(d) {
		return rn, rd
	}
	return 0, 0
}

func Euler033() int {
	num, den := 1, 1

	for d := 10; d < 100; d++ {
		// Ensure numerator is always less than denominator
		for n := 10; n < d; n++ {
			a, b := cancelDigits(n, d)
			if a != 0 {
				// Calculate product of fractions.
				num *= a
				den *= b
			}
		}
	}

	// Find GCD for reduction.
	g := sink.GCD(num, den)
	return den / g
}
