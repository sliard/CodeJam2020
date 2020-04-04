package main

import (
	"fmt"
	"sort"
)

type activity struct {
	index  int
	start  int
	end    int
	parent rune
}


type sortByStart []*activity
type sortByEnd []*activity

func (activities sortByStart) Len() int {
	return len(activities)
}

func (activities sortByStart) Less(i, j int) bool {
	return activities[i].start < activities[j].start
}

func (activities sortByStart) Swap(i, j int) {
	activities[i], activities[j] = activities[j], activities[i]
}

func (activities sortByEnd) Len() int {
	return len(activities)
}

func (activities sortByEnd) Less(i, j int) bool {
	return activities[i].end < activities[j].end
}

func (activities sortByEnd) Swap(i, j int) {
	activities[i], activities[j] = activities[j], activities[i]
}


func main() {

	var nbCase int64
	if _, err := fmt.Scanln(&nbCase); err != nil {
		panic(err)
	}
	for caseId := int64(1); caseId <= nbCase; caseId++ {
		var nbRange int
		if _, err := fmt.Scanln(&nbRange); err != nil {
			panic(err)
		}

		var allStartTime sortByStart
		var allEndTime sortByEnd
		var allActivities []*activity

		for rangeIndex := 0; rangeIndex < nbRange; rangeIndex++ {
			var startTime int
			var endTime int
			if _, err := fmt.Scanln(&startTime, &endTime); err != nil {
				panic(err)
			}
			currentActivity := activity{index: rangeIndex, start:startTime, end:endTime, parent:' '}

			allStartTime = append(allStartTime, &currentActivity)
			allEndTime = append(allEndTime, &currentActivity)
			allActivities = append(allActivities, &currentActivity)
		}

		QuicksortArray(allStartTime)
		QuicksortArray(allEndTime)

		indexStart := 0
		indexEnd := 0

		var result []rune

		allParentAvailable := []rune {'C', 'J'}

		noSolution := false
		for indexStart < len(allStartTime) {
			if allStartTime[indexStart].start >= allEndTime[indexEnd].end {
				allParentAvailable = append(allParentAvailable, allEndTime[indexEnd].parent)
				indexEnd++
			} else if allStartTime[indexStart].start < allEndTime[indexEnd].end {
				if len(allParentAvailable) == 0 {
					noSolution = true
					break
				}
				allStartTime[indexStart].parent = allParentAvailable[0]
				allParentAvailable = allParentAvailable[1:]
				indexStart++
			}
		}

		for acIndex:=0; acIndex<len(allActivities); acIndex++ {
			result = append(result,allActivities[acIndex].parent)
		}

		if noSolution {
			fmt.Printf("Case #%d: IMPOSSIBLE\n",caseId)
		} else {
			fmt.Printf("Case #%d: %s\n",caseId,string(result))
		}

	}
}

func QuicksortArray(s sort.Interface) {
	quicksort(s, 0, s.Len()-1)
}

func quicksort(s sort.Interface, from int, to int) {

	if from >= to {
		return
	}

	p := partition(s, from, to)
	if p > 0 {
		quicksort(s, from, p-1)
	}
	if p+1 < s.Len() {
		quicksort(s, p+1, to)
	}
}

// partition finds a good pivot, partitions the elements in bigger and smaller
// than the pivot, and returns the pivot's new index
func partition(s sort.Interface, from int, to int) int {
	// move pivot to the end
	s.Swap(from, to)

	// last sorted element
	walk := from
	for i := from; i < to; i++ {
		if s.Less(i, to) {
			s.Swap(i, walk)
			walk += 1
		}
	}
	s.Swap(walk, to)
	return walk
}