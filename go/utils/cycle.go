package utils

import "slices"

func FindPeriod[T comparable](data []T) (period int) {
	for {
		period = FindNextPeriod(data, period)
		if period == 0 {
			return 0
		}
		if CheckPeriod(data, period, 0) {
			return period
		}
	}
}

func FindNextPeriod[T comparable](data []T, start int) (period int) {
	i := 1
	for i < len(data) {
		if data[i] == data[0] {
			p1 := 0
			p2 := i
			for data[p1] == data[p2] {
				p1++
				p2++
				if p1 == i && i > start {
					return i
				}
			}
		}
		i++
	}
	return 0
}

func CheckPeriod[T comparable](data []T, period int, dump int) bool {
	if period > len(data)/2 {
		return false
	}

	var enoughChecked bool

	k := period
	for k+period < len(data)-dump {
		if !slices.Equal(data[:period], data[k:k+period]) {
			return false
		}
		enoughChecked = true
		k += period
	}
	remain := len(data) - k
	if remain == 0 {
		return true
	}

	var count int
	for i := k; i < len(data); i++ {
		if (data[i] != data[i%period]) && ((len(data) - i) > dump) {
			return false
		}
		if data[i] == data[i%period] {
			count++
		}
	}
	if count < period && !enoughChecked {
		return false
	}

	return true
}
