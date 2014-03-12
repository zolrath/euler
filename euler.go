package main

import (
	"flag"
	"fmt"
	"sort"

	"github.com/zolrath/euler/problems"
)

func main() {
	var pFlag = flag.Int("p", 0, "Run a specific problem.")
	flag.Parse()

	problemList := map[int]func() int{}
	problemList[1] = problems.Euler001
	problemList[2] = problems.Euler002
	problemList[3] = problems.Euler003
	problemList[4] = problems.Euler004
	problemList[5] = problems.Euler005
	problemList[6] = problems.Euler006
	problemList[7] = problems.Euler007
	problemList[8] = problems.Euler008
	problemList[9] = problems.Euler009
	problemList[10] = problems.Euler010
	problemList[11] = problems.Euler011
	problemList[12] = problems.Euler012
	problemList[13] = problems.Euler013
	problemList[14] = problems.Euler014
	problemList[15] = problems.Euler015

	if *pFlag != 0 {
		fmt.Printf("Problem %03d: %d\n", *pFlag, problemList[*pFlag]())
	} else {
		problemKeys := []int{}
		for k, _ := range problemList {
			problemKeys = append(problemKeys, k)
		}
		sort.Ints(problemKeys)
		for _, v := range problemKeys {
			fmt.Printf("Problem %03d: %d\n", v, problemList[v]())
		}
	}
}
