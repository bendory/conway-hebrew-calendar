package conway

import (
	"testing"
	"time"
)

func TestHeight(t *testing.T) {
	for _, tc := range []struct {
		m    time.Month
		d, h int
	}{
		// Year doesn't matter.
		{time.January, 18, 31},
		{time.February, 18, 32},
		{time.March, 18, 21},
		{time.April, 18, 22},
		{time.May, 18, 23},
		{time.June, 18, 24},
		{time.July, 18, 25},
		{time.August, 18, 26},
		{time.September, 18, 27},
		{time.October, 18, 28},
		{time.November, 18, 29},
		{time.December, 18, 30},
	} {
		if got, want := height(tc.d, tc.m), tc.h; got != want {
			t.Errorf("height(%d, %s): got %d; want %d", tc.d, tc.m, got, want)
		}
	}
}
