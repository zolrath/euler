package main

import (
	"flag"
	"fmt"
	"runtime"
	"sort"

	"github.com/zolrath/euler/problems"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
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
	problemList[16] = problems.Euler016
	problemList[17] = problems.Euler017
	problemList[18] = problems.Euler018
	problemList[19] = problems.Euler019

	problemList[20] = problems.Euler020
	problemList[21] = problems.Euler021
	problemList[22] = problems.Euler022
	problemList[23] = problems.Euler023
	problemList[24] = problems.Euler024
	problemList[25] = problems.Euler025
	problemList[26] = problems.Euler026
	problemList[27] = problems.Euler027
	problemList[28] = problems.Euler028
	problemList[29] = problems.Euler029

	problemList[30] = problems.Euler030
	problemList[31] = problems.Euler031
	problemList[32] = problems.Euler032
	problemList[33] = problems.Euler033
	problemList[34] = problems.Euler034
	problemList[35] = problems.Euler035
	problemList[36] = problems.Euler036
	problemList[37] = problems.Euler037
	problemList[38] = problems.Euler038
	problemList[39] = problems.Euler039

	problemList[40] = problems.Euler040
	problemList[41] = problems.Euler041
	problemList[42] = problems.Euler042
	problemList[43] = problems.Euler043
	problemList[44] = problems.Euler044
	problemList[45] = problems.Euler045
	problemList[46] = problems.Euler046
	problemList[47] = problems.Euler047
	problemList[48] = problems.Euler048
	problemList[49] = problems.Euler049

	problemList[50] = problems.Euler050
	problemList[51] = problems.Euler051
	problemList[52] = problems.Euler052
	problemList[53] = problems.Euler053
	problemList[55] = problems.Euler055
	problemList[56] = problems.Euler056

	if *pFlag != 0 {
		if _, ok := problemList[*pFlag]; !ok {
			fmt.Printf("Problem %03d hasn't been completed yet don'tcha know!\n", *pFlag)
			return
		}
		fmt.Printf("Problem %03d: %d\n", *pFlag, problemList[*pFlag]())
	} else {
		problemKeys := make([]int, 0, len(problemList))
		for k, _ := range problemList {
			problemKeys = append(problemKeys, k)
		}
		sort.Ints(problemKeys)
		for _, v := range problemKeys {
			answer := problemList[v]()
			fmt.Printf("Problem %03d: %d\n", v, answer)
		}
	}
}
