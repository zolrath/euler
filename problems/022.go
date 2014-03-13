/*
Problem 22: Names Scores
Using 022.txt, a 46K text file containing over five-thousand first names,
begin by sorting it into alphabetical order. Then working out the alphabetical
value for each name, multiply this value by its alphabetical position in the
list to obtain a name score.
For example, when the list is sorted into alphabetical order,
COLIN, which is worth 3 + 15 + 12 + 9 + 14 = 53, is the 938th name in the list.
So, COLIN would obtain a score of 938 × 53 = 49714.
What is the total of all the name scores in the file?
Answer: 871198282
*/

package problems

import (
	"io/ioutil"
	"sort"
	"strings"
)

var names []string

func newNameList() []string {
	f, err := ioutil.ReadFile("problems/022.txt")
	if err != nil {
		panic(err)
	}
	namelist := strings.Split(string(f), "\",\"")
	namelist[0] = strings.TrimLeft(namelist[0], "\"")
	namelist[len(namelist)-1] = strings.TrimRight(namelist[len(namelist)-1], "\"")
	sort.Strings(namelist)
	return namelist
}

func scoreName(name string) int {
	var score int
	for _, l := range name {
		score += int(byte(l) - '@')
	}
	return score
}

func Euler022() int {
	var score int
	names = newNameList()
	for i, v := range names {
		score += scoreName(v) * (i + 1)
	}
	return score
}
