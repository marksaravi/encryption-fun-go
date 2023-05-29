package mathematics

func evenArrayCloning(a []int) []int {
	l := len(a)
	if l%2 != 0 {
		l++
	}
	b := make([]int, l)
	b[l-1] = 1
	copy(b, a)
	return b
}

// func ReminderOfPower(m, e, n int) int {
// 	ma := make([]int, e)
// 	for i := 0; i < e; i++ {
// 		ma[i] = m
// 	}
// 	return ReminderOfMultiplications(ma, n)
// }

func ReminderOfMultiplications(m []int, n int) int {
	var rem int = 1

	for i := 0; i < len(m); i += 2 {
		k := m[i] * rem
		if i < len(m)-1 {
			k *= m[i+1]
		}
		if k < n {
			rem = k
		} else {
			rem = k % n
		}
	}
	return rem
}

func GreatestCommonDivisor(a, b int) int {
	n1 := a
	n2 := b

	for n1 > 1 && n2 > 1 && n1 != n2 {
		if n1 > n2 {
			n1, n2 = n2, n1
		}
		n2 %= n1
		n1 -= n2
	}
	if n1 == 0 || n1 == n2 {
		return n2
	}
	if n2 == 0 {
		return n1
	}
	return 1
}

func ReminderOfPower(m, e, n int) int {
	var r int = 1
	for i := 0; i < e; i++ {
		r *= m
		r %= n
	}
	return r
}
