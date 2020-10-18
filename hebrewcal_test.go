package conway

import (
	"fmt"
	"testing"
)

func TestHebrew(t *testing.T) {
	t.Run("format", func(t *testing.T) {
		hebrew := HebrewDate{Y: 5278, M: Shevat, D: 25}
		if got, want := fmt.Sprintf("%s", hebrew), "25 Shevat 5278"; got != want {
			t.Errorf("got %q; want %q", got, want)
		}
	})
}
