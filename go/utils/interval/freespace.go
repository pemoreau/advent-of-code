package interval

import (
	"fmt"
)

// ordered list of disjoint intervals
type FreeSpace struct {
	intervals []Interval
}

func (fs *FreeSpace) String() string {
	res := ""
	//res += fmt.Sprintf("[%d] ", len(fs.intervals))
	for _, i := range fs.intervals {
		res += fmt.Sprintf("%v ", i)
	}
	return res
}

func (fs *FreeSpace) Add(ab Interval) {
	// https://coderbyte.com/algorithm/insert-interval-into-list-of-sorted-disjoint-intervals
	var newSet []Interval
	var endSet []Interval
	i := 0
	// add intervals that come before the new interval
	for i < len(fs.intervals) && fs.intervals[i].Max < ab.Min {
		newSet = append(newSet, fs.intervals[i])
		i++
	}

	// add our new interval to this final list
	newSet = append(newSet, ab)

	// check each interval that comes after the new interval to determine if we can merge
	// if no merges are required then populate a list of the remaining intervals
	for i < len(fs.intervals) {
		var last = newSet[len(newSet)-1]
		if fs.intervals[i].Min < last.Max {
			newInterval := Interval{min(last.Min, fs.intervals[i].Min), max(last.Max, fs.intervals[i].Max)}
			newSet[len(newSet)-1] = newInterval
		} else {
			endSet = append(endSet, fs.intervals[i])
		}
		i++
	}
	fs.intervals = append(newSet, endSet...)
}

func (fs *FreeSpace) Merge() {
	if len(fs.intervals) == 0 {
		return
	}
	var newSet []Interval
	var last = fs.intervals[0]
	i := 1

	for i < len(fs.intervals) {
		if last.Max+1 >= fs.intervals[i].Min {
			last.Max = fs.intervals[i].Max
		} else {
			newSet = append(newSet, last)
			last = fs.intervals[i]
		}
		i++
	}
	newSet = append(newSet, last)
	fs.intervals = newSet
}

func (fs *FreeSpace) Intersect(ab Interval) {
	var newSet []Interval
	for _, i := range fs.intervals {
		if i.Max < ab.Min {
			continue
		}
		if i.Min > ab.Max {
			break
		}
		newSet = append(newSet, Interval{max(i.Min, ab.Min), min(i.Max, ab.Max)})
	}
	fs.intervals = newSet
}

func (fs *FreeSpace) Cardinality() int {
	res := 0
	for _, i := range fs.intervals {
		res += i.Max - i.Min + 1
	}
	return res
}

func (fs *FreeSpace) Len() int {
	return len(fs.intervals)
}

func (fs *FreeSpace) IsEmpty() bool {
	return len(fs.intervals) == 0
}

func (fs *FreeSpace) Get(index int) Interval {
	return fs.intervals[index]
}
