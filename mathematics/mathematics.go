package mathematics

import (
	"testing"
)

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

func ModularMultiplicativeInverse(a, b int, tst *testing.T) (int, int, int, int, int) {
	gcd := GreatestCommonDivisor(a, b)
	r0 := a / gcd
	s0 := 1
	t0 := 0

	r1 := b / gcd
	s1 := 0
	t1 := 1

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
		tst.Logf("%d, %d, %d, %d, %d, %d, %d\n", q, r0, s0, t0, r1, s1, t1)
	}
	c0 := s0
	c1 := t0
	c2 := s0 + b
	c3 := t0 - a
	if s0 < 0 {
		c0, c2 = c2, c0
		c1, c3 = c3, c1
	}
	return c0, c1, c2, c3, gcd
}

func ModularMultiplicativeInverseGCD1(a, b int, tst *testing.T) (int, int, int, int, int) {
	gcd := GreatestCommonDivisor(a, b)
	r0 := a / gcd
	s0 := 1
	t0 := 0

	r1 := b / gcd
	s1 := 0
	t1 := 1

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
		tst.Logf("%d, %d, %d, %d, %d, %d, %d\n", q, r0, s0, t0, r1, s1, t1)
	}
	c0 := s0
	c1 := t0
	c2 := s0 + b
	c3 := t0 - a
	if s0 < 0 {
		c0, c2 = c2, c0
		c1, c3 = c3, c1
	}
	return c0, c1, c2, c3, gcd
}
