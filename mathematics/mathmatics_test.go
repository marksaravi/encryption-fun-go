package mathematics_test

import (
	"testing"

	"github.com/marksaravi/encryption-fun-go/mathematics"
)

func TestReminder(t *testing.T) {
	var testCases = []struct {
		m, e, n, want int
	}{
		{
			n:    3233,
			e:    17,
			m:    65,
			want: 2790,
		},
		{
			n:    3233,
			e:    413,
			m:    2790,
			want: 65,
		},
		{
			n:    17,
			e:    3,
			m:    14,
			want: 7,
		},
		{
			n:    3,
			e:    2,
			m:    4,
			want: 1,
		},
	}
	for i, tc := range testCases {
		got := mathematics.ReminderOfPower(tc.m, tc.e, tc.n)
		t.Logf("test#%d: calculated reminder of %d^%d to %d is %d", i+1, tc.m, tc.e, tc.n, got)
		if got != tc.want {
			t.Errorf("test#%d: expected the reminder of %d^%d to be %d but got %d", i+1, tc.m, tc.e, tc.n, got)
		}
	}
}

func TestLargestCommonFactor(t *testing.T) {
	var testCases = []struct {
		a, b, want int
	}{
		{
			a:    49,
			b:    35,
			want: 7,
		},
		{
			a: 16, b: 64, want: 16,
		},
		{
			a:    33,
			b:    44,
			want: 11,
		},
	}

	for _, tc := range testCases {
		got := mathematics.GreatestCommonDivisor(tc.a, tc.b)
		if got != tc.want {
			t.Errorf("expected  %d, got %d", tc.want, got)
		}
	}
}

func TestNumOfCoPrimes(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36}
	yN := []int{1, 1, 2, 2, 4, 2, 6, 2, 6, 4, 10, 2, 12, 6, 4, 4, 16, 6, 18, 4, 6, 10, 22, 2, 20, 12, 18, 6, 28, 4, 30, 8, 10, 16, 12, 6}       //λ(n)
	phiN := []int{1, 1, 2, 2, 4, 2, 6, 4, 6, 4, 10, 4, 12, 6, 8, 8, 16, 6, 18, 8, 12, 10, 22, 8, 20, 12, 18, 12, 28, 8, 30, 16, 20, 16, 24, 12} //φ(n)
	for testIndex, n := range numbers {
		t.Logf("test#%d: finding number of copromes for %d\n", testIndex, n)
		phiNgot := mathematics.FindCoprimes(n)
		yNgot := mathematics.FindCarmichael(n)
		if len(phiNgot) != phiN[testIndex] {
			t.Errorf("Test failed: expected number of co-primes: %d, got: %d\n", phiN[testIndex], phiNgot)
		}
		if yNgot != yN[testIndex] {
			t.Errorf("Test failed: expected carmichael: %d, got: %d\n", yN[testIndex], yNgot)
		}
	}
}

func TestModularMultiplicativeInverse(t *testing.T) {
	var testCases = []struct{ a, b, s int }{
		{
			a: 17,
			b: 780,
		},
		{
			a: 46,
			b: 240,
		},
	}

	for _, tc := range testCases {
		c0, c1, c2, c3, gcd := mathematics.ModularMultiplicativeInverse(tc.a, tc.b)
		r0 := c0*tc.a + c1*tc.b
		r1 := c2*tc.a + c3*tc.b
		t.Logf("C0: %d, C1: %d, C2: %d, C3: %d\n", c0, c1, c2, c3)
		if r0 != gcd || r1 != gcd {
			t.Errorf("test failed: modular multiplicative inverse incorrect : %d, and: %d\n", r0, r1)
		}
	}
}
