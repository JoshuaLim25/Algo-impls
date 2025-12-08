package main

import (
	"fmt"
	"os"
)

type Interval struct {
	start int
	end   int
}

func AllMeetingsPossible(times []Interval) bool {
	curTime := 0
	possible := true
	for _, interval := range times {
		if interval.start < curTime {
			return !possible
		}
		timeSpan := interval.end - interval.start
		curTime += timeSpan
	}
	return possible
}

func main() {
	intervals1 := []Interval{
		{start: 0, end: 30},
		{start: 5, end: 10},
		{start: 15, end: 20},
	}
	intervals2 := []Interval{
		{start: 5, end: 8},
		{start: 8, end: 20},
	}
	maybe1 := AllMeetingsPossible(intervals1)
	fmt.Fprintf(os.Stderr, "DEBUG[1]: meeting_rooms.go:33: maybe1=%+v\n", maybe1)
	maybe2 := AllMeetingsPossible(intervals2)
	fmt.Fprintf(os.Stderr, "DEBUG[2]: meeting_rooms.go:35: maybe2=%+v\n", maybe2)
}
