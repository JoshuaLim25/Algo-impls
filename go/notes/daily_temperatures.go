package main 

import (
	"fmt"
	"os"
)

func DaysUntilWarmer(dailyTemps []int) []int {
	answer := []int {} // a[i] = # days to wait after day i to get a warmer temp
	// a mapping from idx `i` |-> # days a[i]
	// NOTE:  If there is no future day for which this is possible, keep answer[i] == 0 instea
	for i := range dailyTemps {
		daysWaited := 1
		changed := false
		for j := i + 1; j < len(dailyTemps); j++ {
			if dailyTemps[j] > dailyTemps[i] {
				answer = append(answer, daysWaited)
				changed = true
				break
			}
			daysWaited++
		}
		if !changed {
			// answer[i] = 0
			answer = append(answer, 0)
		}
	}
	return answer
}


func main() {
	temps := []int{ 73,74,75,71,69,72,76,73 }
	fmt.Fprintf(os.Stderr, "DEBUG[2]: daily-temperatures.go:18: temps=%+v\n", temps)
	sanityCheck := []int {1,1,4,2,1,1,0,0}
	daysToWait := DaysUntilWarmer(temps)
	fmt.Fprintf(os.Stderr, "DEBUG[3]: daily-temperatures.go:20: daysToWait(temps) should be %+v, is %+v\n", sanityCheck, daysToWait)

	temps2  := []int{ 30,40,50,60 }
	sanityCheck2 := []int {1,1,1,0}
	daysToWait2 := DaysUntilWarmer(temps2)
	fmt.Fprintf(os.Stderr, "DEBUG[3]: daily-temperatures.go:20: daysToWait(temps2) should be %+v, is %+v\n", sanityCheck2, daysToWait2)
}
