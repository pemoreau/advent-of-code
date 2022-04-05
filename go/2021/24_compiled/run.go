package main

func Run(inp []int) (w, x, y, z int) {
	w = 0
	x = 0
	y = 0
	z = 0

	w = inp[0]
	x = 0
	x = x + z
	x = x % 26

	x = x + 14
	if x == w {
		x = 1
	} else {
		x = 0
	}
	if x == 0 {
		x = 1
	} else {
		x = 0
	}
	y = 0
	y = y + 25
	y = y * x
	y = y + 1
	z = z * y
	y = 0
	y = y + w
	y = y + 12
	y = y * x
	z = z + y
	w = inp[1]
	x = 0
	x = x + z
	x = x % 26

	x = x + 15
	if x == w {
		x = 1
	} else {
		x = 0
	}
	if x == 0 {
		x = 1
	} else {
		x = 0
	}
	y = 0
	y = y + 25
	y = y * x
	y = y + 1
	z = z * y
	y = 0
	y = y + w
	y = y + 7
	y = y * x
	z = z + y
	w = inp[2]
	x = 0
	x = x + z
	x = x % 26

	x = x + 12
	if x == w {
		x = 1
	} else {
		x = 0
	}
	if x == 0 {
		x = 1
	} else {
		x = 0
	}
	y = 0
	y = y + 25
	y = y * x
	y = y + 1
	z = z * y
	y = 0
	y = y + w
	y = y + 1
	y = y * x
	z = z + y
	w = inp[3]
	x = 0
	x = x + z
	x = x % 26

	x = x + 11
	if x == w {
		x = 1
	} else {
		x = 0
	}
	if x == 0 {
		x = 1
	} else {
		x = 0
	}
	y = 0
	y = y + 25
	y = y * x
	y = y + 1
	z = z * y
	y = 0
	y = y + w
	y = y + 2
	y = y * x
	z = z + y
	w = inp[4]
	x = 0
	x = x + z
	x = x % 26
	z = z / 26
	x = x + -5
	if x == w {
		x = 1
	} else {
		x = 0
	}
	if x == 0 {
		x = 1
	} else {
		x = 0
	}
	y = 0
	y = y + 25
	y = y * x
	y = y + 1
	z = z * y
	y = 0
	y = y + w
	y = y + 4
	y = y * x
	z = z + y
	w = inp[5]
	x = 0
	x = x + z
	x = x % 26

	x = x + 14
	if x == w {
		x = 1
	} else {
		x = 0
	}
	if x == 0 {
		x = 1
	} else {
		x = 0
	}
	y = 0
	y = y + 25
	y = y * x
	y = y + 1
	z = z * y
	y = 0
	y = y + w
	y = y + 15
	y = y * x
	z = z + y
	w = inp[6]
	x = 0
	x = x + z
	x = x % 26

	x = x + 15
	if x == w {
		x = 1
	} else {
		x = 0
	}
	if x == 0 {
		x = 1
	} else {
		x = 0
	}
	y = 0
	y = y + 25
	y = y * x
	y = y + 1
	z = z * y
	y = 0
	y = y + w
	y = y + 11
	y = y * x
	z = z + y
	w = inp[7]
	x = 0
	x = x + z
	x = x % 26
	z = z / 26
	x = x + -13
	if x == w {
		x = 1
	} else {
		x = 0
	}
	if x == 0 {
		x = 1
	} else {
		x = 0
	}
	y = 0
	y = y + 25
	y = y * x
	y = y + 1
	z = z * y
	y = 0
	y = y + w
	y = y + 5
	y = y * x
	z = z + y
	w = inp[8]
	x = 0
	x = x + z
	x = x % 26
	z = z / 26
	x = x + -16
	if x == w {
		x = 1
	} else {
		x = 0
	}
	if x == 0 {
		x = 1
	} else {
		x = 0
	}
	y = 0
	y = y + 25
	y = y * x
	y = y + 1
	z = z * y
	y = 0
	y = y + w
	y = y + 3
	y = y * x
	z = z + y
	w = inp[9]
	x = 0
	x = x + z
	x = x % 26
	z = z / 26
	x = x + -8
	if x == w {
		x = 1
	} else {
		x = 0
	}
	if x == 0 {
		x = 1
	} else {
		x = 0
	}
	y = 0
	y = y + 25
	y = y * x
	y = y + 1
	z = z * y
	y = 0
	y = y + w
	y = y + 9
	y = y * x
	z = z + y
	w = inp[10]
	x = 0
	x = x + z
	x = x % 26

	x = x + 15
	if x == w {
		x = 1
	} else {
		x = 0
	}
	if x == 0 {
		x = 1
	} else {
		x = 0
	}
	y = 0
	y = y + 25
	y = y * x
	y = y + 1
	z = z * y
	y = 0
	y = y + w
	y = y + 2
	y = y * x
	z = z + y
	w = inp[11]
	x = 0
	x = x + z
	x = x % 26
	z = z / 26
	x = x + -8
	if x == w {
		x = 1
	} else {
		x = 0
	}
	if x == 0 {
		x = 1
	} else {
		x = 0
	}
	y = 0
	y = y + 25
	y = y * x
	y = y + 1
	z = z * y
	y = 0
	y = y + w
	y = y + 3
	y = y * x
	z = z + y
	w = inp[12]
	x = 0
	x = x + z
	x = x % 26
	z = z / 26
	x = x + 0
	if x == w {
		x = 1
	} else {
		x = 0
	}
	if x == 0 {
		x = 1
	} else {
		x = 0
	}
	y = 0
	y = y + 25
	y = y * x
	y = y + 1
	z = z * y
	y = 0
	y = y + w
	y = y + 3
	y = y * x
	z = z + y
	w = inp[13]
	x = 0
	x = x + z
	x = x % 26
	z = z / 26
	x = x + -4
	if x == w {
		x = 1
	} else {
		x = 0
	}
	if x == 0 {
		x = 1
	} else {
		x = 0
	}
	y = 0
	y = y + 25
	y = y * x
	y = y + 1
	z = z * y
	y = 0
	y = y + w
	y = y + 11
	y = y * x
	z = z + y
	return
}
