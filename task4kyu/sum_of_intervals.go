package task4kyu

import (
	"sort"
)

func SumOfIntervals(intervals [][2]int) int {
	if len(intervals) == 0 {
		return 0
	}
	// Sort the intervals based on their start time
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	// Merge overlapping intervals and calculate their length
	mergedIntervals := make([][2]int, 0)
	mergedIntervals = append(mergedIntervals, intervals[0])
	for _, interval := range intervals {
		lastMerged := mergedIntervals[len(mergedIntervals)-1]
		if interval[0] <= lastMerged[1] && interval[1] > lastMerged[1] {
			mergedIntervals[len(mergedIntervals)-1][1] = interval[1]
		} else if interval[0] > lastMerged[1] {
			mergedIntervals = append(mergedIntervals, interval)
		}
	}
	// Calculate the sum of the lengths of the merged intervals
	sum := 0
	for _, interval := range mergedIntervals {
		sum += interval[1] - interval[0]
	}
	return sum
}
