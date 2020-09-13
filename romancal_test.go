package conway

import (
	"fmt"
	"testing"
	"time"
)

func TestGregorian(t *testing.T) {
	t.Run("round-trip", func(t *testing.T) {
		want := time.Date(1968, time.February, 24, 12, 0, 0, 0, time.Local)
		gd := NewGregorianDate(want)
		got := gd.Time()

		if !want.Equal(got) {
			t.Errorf("got %v; want %v", got, want)
		}
	})
	t.Run("round-trip-2", func(t *testing.T) {
		gregorian := GregorianDate{y: 1968, d: 24, m: time.February}
		goDate := gregorian.Time()
		got := NewGregorianDate(goDate)

		if gregorian.y != got.y || gregorian.m != got.m || gregorian.d != got.d {
			t.Errorf("got %v; want %v", got, gregorian)
		}
	})
	t.Run("format", func(t *testing.T) {
		gregorian := GregorianDate{y: 1968, d: 24, m: time.February}
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
			{GregorianDate{y: 2019, d: 18, m: time.March}, 21},
			{GregorianDate{y: 0, d: 18, m: time.March}, 21}, // year doesn't matter
			{GregorianDate{y: 2019, d: 18, m: time.January}, 31},
			{GregorianDate{y: 2019, d: 18, m: time.February}, 32},
			{GregorianDate{y: 2019, d: 18, m: time.April}, 22},
			{GregorianDate{y: 2019, d: 18, m: time.May}, 23},
			{GregorianDate{y: 2019, d: 18, m: time.June}, 24},
			{GregorianDate{y: 2019, d: 18, m: time.July}, 25},
			{GregorianDate{y: 2019, d: 18, m: time.August}, 26},
			{GregorianDate{y: 2019, d: 18, m: time.September}, 27},
			{GregorianDate{y: 2019, d: 18, m: time.October}, 28},
			{GregorianDate{y: 2019, d: 18, m: time.November}, 29},
			{GregorianDate{y: 2019, d: 18, m: time.December}, 30},
		} {
			if got, want := tc.g.height(), tc.h; got != want {
				t.Errorf("%s: got %d; want %d", tc.g, got, want)
			}
		}
	})
	t.Run("squash", func(t *testing.T) {
		for _, tc := range []struct {
			g, want GregorianDate
		}{
			{GregorianDate{y: 2019, d: 15, m: time.January}, GregorianDate{y: 2019, d: 15, m: time.January}},
			{GregorianDate{y: 2019, d: -5, m: time.January}, GregorianDate{y: 2018, d: 26, m: time.December}},
			{GregorianDate{y: 2019, d: 34, m: time.December}, GregorianDate{y: 2020, d: 3, m: time.January}},
			{GregorianDate{y: 2020, d: 34, m: time.February}, GregorianDate{y: 2020, d: 5, m: time.March}},
			{GregorianDate{y: 2019, d: 34, m: time.February}, GregorianDate{y: 2019, d: 6, m: time.March}},
			{GregorianDate{y: 1900, d: 34, m: time.February}, GregorianDate{y: 1900, d: 6, m: time.March}},
		} {
			was := fmt.Sprintf("%#v", tc.g)
			tc.g.squash()
			if tc.g.y != tc.want.y || tc.g.d != tc.want.d || tc.g.m != tc.want.m {
				t.Errorf("%s: got %#v; want %#v", was, tc.g, tc.want)
			}
		}
	})
}
