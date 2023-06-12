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

func FindCoprimes(a int) []int {
	coprimes := make([]int, 0, 1000)
	coprimes = append(coprimes, 1)
	for n := 2; n < a; n++ {
		if GreatestCommonDivisor(a, n) == 1 {
			coprimes = append(coprimes, n)
		}
	}
	return coprimes
}

func FindCarmichael(n int) int {
	coPrimes := FindCoprimes(n)
	c := 1
	for m := 1; m < n; m++ {
		found := true
		for _, a := range coPrimes {
			if ReminderOfPower(a, m, n) != 1 {
				found = false
				break
			}
		}
		if found {
			c = m
			break
		}
	}
	return c
}

func CarmichaelOfPQ(p, q int) int {
	n1 := p - 1
	n2 := q - 1
	return n1 * n2 / GreatestCommonDivisor(n1, n2)
}

func D(a, b int) int {
	for n := 1; n < 1000; n++ {
		k := n*b + 1
		if k%a == 0 {
			return k / a
		}
	}
	return 0
}
