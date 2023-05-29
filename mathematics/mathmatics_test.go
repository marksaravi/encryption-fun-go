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
