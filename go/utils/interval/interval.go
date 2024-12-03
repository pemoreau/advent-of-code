package interval

import (
	"fmt"
	"github.com/pemoreau/advent-of-code/go/utils"
)

type Interval struct{ Min, Max int }

func (i Interval) String() string {
	return fmt.Sprintf("[%d, %d]", i.Min, i.Max)
}

// https{//stackoverflow.com/questions/31057473/calculating-the-modulo-of-two-Intervals
func (i Interval) Len() int {
	return i.Max - i.Min + 1
}

func (i Interval) Negate() Interval {
	return Interval{-i.Max, -i.Min}
}

func (i Interval) Add(j Interval) Interval {
	return Interval{i.Min + j.Min, i.Max + j.Max}
}

func (i Interval) Sub(j Interval) Interval {
	a := i.Min - j.Min
	b := i.Max - j.Max
	if a < b {
		return Interval{a, b}
	} else {
		return Interval{b, a}
	}
}

func (i Interval) Contains(x int) bool {
	return i.Min <= x && x <= i.Max
}

func minmax(a, b, c, d int) (int, int) {
	var min1, max1, min2, max2 int
	if a <= b {
		min1 = a
		max1 = b
	} else {
		min1 = b
		max1 = a
	}
	if c <= d {
		min2 = c
		max2 = d
	} else {
		min2 = d
		max2 = c
	}
	return min(min1, min2), max(max1, max2)
}

func (i Interval) Mul(j Interval) Interval {
	if i.Min >= 0 && j.Min >= 0 {
		return Interval{i.Min * j.Min, i.Max * j.Max}
	} else if i.Max <= 0 && j.Max <= 0 {
		return Interval{i.Max * j.Max, i.Min * j.Min}
	}

	// From https://doi.org/10.1145/502102.502106
	// Figure 6
	// but it is not more efficient than the following 2 lines

	a, b := minmax(i.Min*j.Min, i.Min*j.Max, i.Max*j.Min, i.Max*j.Max)
	return Interval{a, b}
}

func (i Interval) Div(j Interval) Interval {
	if i.Min >= 0 && j.Min >= 0 {
		return Interval{i.Min / j.Max, i.Max / j.Min}
	} else if i.Max <= 0 && j.Max <= 0 {
		return Interval{i.Max / j.Min, i.Min / j.Max}
	}

	// From https://doi.org/10.1145/502102.502106
	// Figure 7
	// but it is not more efficient than the following 2 lines

	a, b := minmax(i.Min/j.Min, i.Min/j.Max, i.Max/j.Min, i.Max/j.Max)
	return Interval{a, b}
}

func Empty() Interval { return Interval{0, -1} }

func (i Interval) Inter(j Interval) Interval {
	if i.Max < j.Min || j.Max < i.Min {
		return Empty()
	}
	a := max(i.Min, j.Min)
	b := min(i.Max, j.Max)
	return Interval{a, b}
}

func (i Interval) Disjoint(j Interval) bool {
	return i.Max < j.Min || j.Max < i.Min
}

func (i Interval) union(j Interval) Interval {
	return Interval{min(i.Min, j.Min), max(i.Max, j.Max)}
}

func (i Interval) Mod1(m int) Interval {
	a := i.Min
	b := i.Max
	switch {
	case a > b || m == 0:
		// (1): empty Interval
		return Interval{0, -1}
	case b < 0:
		// (2): compute modulo with positive Interval and negate
		return Interval{-b, -a}.Mod1(m).Negate()
	case a < 0:
		// (3): split into negative and non-negative Interval, compute and join
		return Interval{a, -1}.Mod1(m).union(Interval{0, b}.Mod1(m))
	case b-a < utils.Abs(m) && a%m <= b%m:
		// (4): there is no k > 0 such that a < k*m <= b
		return Interval{a % m, b % m}
	default:
		// (5): we can't do better than that
		return Interval{0, utils.Abs(m) - 1}
	}
}

func (i Interval) Mod2(i2 Interval) Interval {
	a := i.Min
	b := i.Max
	m := i2.Min
	n := i2.Max
	switch {
	case a > b || m > n:
		// (1): empty Interval
		return Interval{0, -1}
	case b < 0:
		// (2): compute modulo with positive Interval and negate
		return Interval{-b, -a}.Mod2(Interval{m, n}).Negate()
	case a < 0:
		// (3): split into negative and non-negative Interval, compute, and join
		return Interval{a, -1}.Mod2(Interval{m, n}).union(Interval{0, b}.Mod2(Interval{m, n}))
	case m == n:
		// (4): use the simpler function from before
		return Interval{a, b}.Mod1(m)
	case n <= 0:
		// (5): use only non-negative m and n
		return Interval{a, b}.Mod2(Interval{-n, -m})
	case m <= 0:
		// (6): similar to (5), make modulus non-negative
		return Interval{a, b}.Mod2(Interval{1, max(-m, n)})
	case b-a >= n:
		// (7): compare to (4) in mod1, check b-a < |modulus|
		return Interval{0, n - 1}
	case b-a >= m:
		// (8): similar to (7), split Interval, compute, and join
		return Interval{0, b - a - 1}.union(Interval{a, b}.Mod2(Interval{b - a + 1, n}))
	case m > b:
		// (9): modulo has no effect
		return Interval{a, b}
	case n > b:
		// (10): there is some overlapping of [a,b] and [n,m]
		return Interval{0, b}
	default:
		// (11): either compute all possibilities and join, or be imprecise
		return Interval{0, n - 1} // imprecise
	}
}
