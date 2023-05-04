package mathematics_test

import (
	"testing"

	"github.com/marksaravi/encryption-fun-go/mathematics"
)

func TestReminder(t *testing.T) {
	const n = 3233
	const e = 17
	const m = 65
	const want = 2790
	got := mathematics.ReminderOfPower(m, e, n)
	t.Logf("calculated reminder of %d^%d to %d is %d", m, e, n, got)
	if got != want {
		t.Errorf("expected the reminder of %d^%d to be %d but got %d", m, e, n, got)
	}
}
