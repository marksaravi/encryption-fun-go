package mathematics

func ReminderOfMultiplications(m []int64, n int64) int64 {
	var rem int64 = 1

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

func GreatestCommonDivisor(a, b int64) int64 {
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

func ReminderOfPower(m, e, n int64) int64 {
	var r int64 = 1
	for i := int64(0); i < e; i++ {
		r *= m
		r %= n
	}
	return r
}

func FindCoprimes(a int64) []int64 {
	coprimes := make([]int64, 0, 1000)
	coprimes = append(coprimes, 1)
	for n := int64(2); n < a; n++ {
		if GreatestCommonDivisor(a, n) == 1 {
			coprimes = append(coprimes, n)
		}
	}
	return coprimes
}

func FindCarmichael(n int64) int64 {
	coPrimes := FindCoprimes(n)
	c := int64(1)
	for m := int64(1); m < n; m++ {
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

func CarmichaelOfPQ(p, q int64) int64 {
	n1 := p - 1
	n2 := q - 1
	return n1 * n2 / GreatestCommonDivisor(n1, n2)
}

func ModularMultiplicativeInverse(a, b int64) (int64, int64, int64, int64, int64) {
	gcd := GreatestCommonDivisor(a, b)
	c0, c1, c2, c3 := ModularMultiplicativeInverseGCD1(a/gcd, b/gcd)
	return c0, c1, c2, c3, gcd
}

func ModularMultiplicativeInverseGCD1(a, b int64) (int64, int64, int64, int64) {
	r0 := int64(a)
	s0 := int64(1)
	t0 := int64(0)

	r1 := int64(b)
	s1 := int64(0)
	t1 := int64(1)

	for r1 > 0 {
		q := r0 / r1
		r := r0 - q*r1
		s := s0 - q*s1
		t := t0 - q*t1
		r0 = r1
		s0 = s1
		t0 = t1
		r1 = r
		t1 = t
		s1 = s
	}
	c0 := s0
	c1 := t0
	c2 := s0 + b
	c3 := t0 - a
	if s0 < 0 {
		c0, c2 = c2, c0
		c1, c3 = c3, c1
	}
	return c0, c1, c2, c3
}
