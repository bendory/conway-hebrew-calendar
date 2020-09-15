package conway

import (
	"fmt"
	"testing"
	"time"
)

func TestGregorian(t *testing.T) {
	t.Run("format", func(t *testing.T) {
		gregorian := NewGregorianDate(1968, time.February, 24)
		got, want := fmt.Sprintf("%s", gregorian), "24 February 1968"
		if want != got {
			t.Errorf("got %q; want %q", got, want)
		}
	})
	t.Run("height", func(t *testing.T) {
		for _, tc := range []struct {
			g GregorianDate
			h int
		}{
			{NewGregorianDate(2019, time.March, 18), 21},
			{NewGregorianDate(0, time.March, 18), 21}, // year doesn't matter
			{NewGregorianDate(2019, time.April, 18), 22},
			{NewGregorianDate(2019, time.May, 18), 23},
			{NewGregorianDate(2019, time.June, 18), 24},
			{NewGregorianDate(2019, time.July, 18), 25},
			{NewGregorianDate(2019, time.August, 18), 26},
			{NewGregorianDate(2019, time.September, 18), 27},
			{NewGregorianDate(2019, time.October, 18), 28},
			{NewGregorianDate(2019, time.November, 18), 29},
			{NewGregorianDate(2019, time.December, 18), 30},
			{NewGregorianDate(2019, time.January, 18), 31},
			{NewGregorianDate(2019, time.February, 18), 32},
		} {
			if got, want := tc.g.height(), tc.h; got != want {
				t.Errorf("%s: got %d; want %d", tc.g, got, want)
			}
		}
	})
}
