package main

import (
	"fmt"
	"sort"
)

/**
 * Definition for an interval.
 * type Interval struct {
 *	   Start int
 *	   End   int
 * }
 */

type Interval struct {
	Start int
	End   int
}

type IntervalSlice []Interval

func (s IntervalSlice) Len() int {
	return len(s)
}

func (s IntervalSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s IntervalSlice) Less(i, j int) bool {
	return s[i].End < s[j].End
}

func eraseOverlapIntervals(intervals []Interval) int {
	intervalsLen := len(intervals)
	if intervalsLen <= 1 {
		return 0
	}
	sort.Sort(IntervalSlice(intervals))
	end := intervals[0].End
	count := 1
	for _, interval := range intervals {
		if interval.Start < end {
			continue
		}
		end = interval.End
		count++
	}
	return len(intervals) - count
}

func main() {
	intervals := []Interval{{1, 2}, {1, 2}, {1, 2}}
	fmt.Println(eraseOverlapIntervals(intervals))
}
