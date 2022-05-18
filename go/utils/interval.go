package utils

type Interval struct{ Min, Max int }

// https{//stackoverflow.com/questions/31057473/calculating-the-modulo-of-two-Intervals
func (a Interval) Len() int {
	return a.Max - a.Min + 1
}

func (a Interval) Negate() Interval {
	return Interval{-a.Max, -a.Min}
}

func (a Interval) Add(b Interval) Interval {
	return Interval{a.Min + b.Min, a.Max + b.Max}
}

func (a Interval) Sub(b Interval) Interval {
	min := a.Min - b.Min
	max := a.Max - b.Max
	if min < max {
		return Interval{min, max}
	} else {
		return Interval{max, min}
	}
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
	return Min(min1, min2), Max(max1, max2)
}

func (a Interval) Mul(b Interval) Interval {
	if a.Min >= 0 && b.Min >= 0 {
		return Interval{a.Min * b.Min, a.Max * b.Max}
	} else if a.Max <= 0 && b.Max <= 0 {
		return Interval{a.Max * b.Max, a.Min * b.Min}
	}

	// From https://doi.org/10.1145/502102.502106
	// Figure 6
	// but it is not more efficient than the following 2 lines

	min, max := minmax(a.Min*b.Min, a.Min*b.Max, a.Max*b.Min, a.Max*b.Max)
	return Interval{min, max}
}

func (a Interval) Div(b Interval) Interval {
	if a.Min >= 0 && b.Min >= 0 {
		return Interval{a.Min / b.Max, a.Max / b.Min}
	} else if a.Max <= 0 && b.Max <= 0 {
		return Interval{a.Max / b.Min, a.Min / b.Max}
	}

	// From https://doi.org/10.1145/502102.502106
	// Figure 7
	// but it is not more efficient than the following 2 lines

	min, max := minmax(a.Min/b.Min, a.Min/b.Max, a.Max/b.Min, a.Max/b.Max)
	return Interval{min, max}
}

func (a Interval) Inter(b Interval) Interval {
	if a.Max < b.Min || b.Max < a.Min {
		return Interval{0, -1}
	}
	min := Max(a.Min, b.Min)
	max := Min(a.Max, b.Max)
	return Interval{min, max}
}

func (a Interval) union(b Interval) Interval {
	min := Min(a.Min, b.Min)
	max := Max(a.Max, b.Max)
	return Interval{min, max}
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
	case b-a < Abs(m) && a%m <= b%m:
		// (4): there is no k > 0 such that a < k*m <= b
		return Interval{a % m, b % m}
	default:
		// (5): we can't do better than that
		return Interval{0, Abs(m) - 1}
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
		return Interval{a, b}.Mod2(Interval{1, Max(-m, n)})
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
