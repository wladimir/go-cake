package main

import (
	"github.com/wladimir/go-cake/server/problems/common"
	"sort"
)

type meeting struct {
	start, end int
}

func main() {
	meetings := []meeting{{4, 5}, {2, 3}, {1, 6}}

	sort.Slice(meetings, func(i, j int) bool {
		return meetings[i].start < meetings[j].start
	})

	result := []meeting{meetings[0]}

	for i := range meetings {
		if meetings[i].start <= result[len(result)-1].end {
			result[len(result)-1].end, _ = common.Max(result[len(result)-1].end, meetings[i].end)
		} else {
			result = append(result, meetings[i])
		}
	}
}
