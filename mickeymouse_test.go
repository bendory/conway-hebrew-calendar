package conway

import (
	"fmt"
	"testing"
	"time"
)

func dmy(t time.Time) string {
	return t.Format("02 January 2006")
}

func TestRoshHashannah(t *testing.T) {
	tests := []struct{ m, d, y, hebrewYear int }{
		{9, 14, 2015, 5776},
		{10, 3, 2016, 5777},
		{9, 21, 2017, 5778},
		// Cover a 19-year cycle...
		{9, 10, 2018, 5779}, // Sept 10 2018 == RH 5779...
		{9, 30, 2019, 5780}, // ref: p. 3
		{9, 19, 2020, 5781},
		{9, 7, 2021, 5782},
		{9, 26, 2022, 5783},
		{9, 16, 2023, 5784},
		{10, 3, 2024, 5785},
		{9, 23, 2025, 5786},
		{9, 12, 2026, 5787},
		{10, 2, 2027, 5788},
		{9, 21, 2028, 5789},
		{9, 10, 2029, 5790},
		{9, 28, 2030, 5791},
		{9, 18, 2031, 5792},
		{9, 6, 2032, 5793},
		{9, 24, 2033, 5794},
		{9, 14, 2034, 5795},
		{10, 4, 2035, 5796},
		{9, 22, 2036, 5797},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("%d-%d", test.y, test.hebrewYear), func(t *testing.T) {
			m := gregorianMickeyMouse(test.y)
			wantY, wantM, wantD := test.y, time.Month(test.m), test.d
			gotY, gotM, gotD := m.rh.Date()
			if gotY != wantY {
				t.Errorf("Year(%s, %s): got %d; want %d", m.rh.Weekday(), m.rh, gotY, wantY)
			}
			if gotM != wantM {
				t.Errorf("Month(%s, %s): got %s; want %s", m.rh.Weekday(), m.rh, gotM, wantM)
			}
			if gotD != wantD {
				t.Errorf("Day(%s, %s): got %d; want %d", m.rh.Weekday(), m.rh, gotD, wantD)
			}
			if got, want := m.hebrewYears[1].Y, test.hebrewYear; got != want {
				t.Errorf("got %d; want %d", got, want)
			}
			switch m.rh.Weekday() {
			case time.Sunday, time.Wednesday, time.Friday:
				t.Errorf("לא אד״ו ראש: %s", m.rh.Weekday())
			}
		})
	}
}

func TestGregorianMickeyMouse(t *testing.T) {
	tests := []struct{ he, she, it, year int }{
		// 2016-2018 ref: p. 2
		{71, 53, 42, 2016},
		{59, 70, 30, 2017},
		{48, 59, 19, 2018},
		{68, 49, 39, 2019}, // ref: p. 6
	}
	for _, test := range tests {
		t.Run(fmt.Sprint(test.year), func(t *testing.T) {
			m := gregorianMickeyMouse(test.year)
			if got, want := m.he, test.he; got != want {
				t.Errorf("HE:  got=%d want=%d", got, want)
			}
			if got, want := m.she, test.she; got != want {
				t.Errorf("SHE: got=%d want=%d", got, want)
			}
			if got, want := m.it, test.it; got != want {
				t.Errorf("IT:  got=%d want=%d", got, want)
			}
		})
	}
}

func TestHebrewMickeyMouse(t *testing.T) {
	tests := []struct{ he, she, it, year int }{
		// 5776-5779 ref: p. 2
		{52, 53, 42, 5776}, // I calculated he=52 based on RH 2015
		{71, 70, 30, 5777},
		{59, 59, 19, 5778},
		{48, 49, 39, 5779}, // ref: p. 6; I calculated he=48
	}
	for _, test := range tests {
		t.Run(fmt.Sprint(test.year), func(t *testing.T) {
			m := hebrewMickeyMouse(test.year)
			if got, want := m.he, test.he; got != want {
				t.Errorf("HE:  got=%d want=%d", got, want)
			}
			if got, want := m.she, test.she; got != want {
				t.Errorf("SHE: got=%d want=%d", got, want)
			}
			if got, want := m.it, test.it; got != want {
				t.Errorf("IT:  got=%d want=%d", got, want)
			}
		})
	}
}

// TODO: add test cases
func TestConversions(t *testing.T) {
	tests := []struct {
		hd HebrewDate
		t  time.Time
	}{{
		// ref: p. 1
		HebrewDate{Y: HebrewYear{Y: 5777}, D: 25, M: Kislev},
		time.Date(2016, time.December, 25, 12, 0, 0, 0, time.Local),
	}, {
		// ref: p. 3
		HebrewDate{Y: HebrewYear{Y: 5779}, D: 7, M: Iyar},
		time.Date(2019, time.May, 12, 12, 0, 0, 0, time.Local),
	}, {
		// ref: p. 3
		HebrewDate{Y: HebrewYear{Y: 5779}, D: 2, M: Nissan},
		time.Date(2019, time.April, 7, 12, 0, 0, 0, time.Local),
	}, {
		// ref: p. 3
		HebrewDate{Y: HebrewYear{Y: 5780}, D: 1, M: Tishrei},
		time.Date(2019, time.September, 30, 12, 0, 0, 0, time.Local),
	}, {
		HebrewDate{Y: HebrewYear{Y: 5728}, D: 25, M: Shevat},
		time.Date(1968, time.February, 24, 12, 0, 0, 0, time.Local),
	}, {
		HebrewDate{Y: HebrewYear{Y: 5750}, D: 29, M: Tevet},
		time.Date(1990, time.January, 26, 12, 0, 0, 0, time.Local),
	}, {
		HebrewDate{Y: HebrewYear{Y: 5750}, D: 30, M: Shevat},
		time.Date(1990, time.February, 25, 12, 0, 0, 0, time.Local),
	}, {
		HebrewDate{Y: HebrewYear{Y: 5750}, D: 1, M: Adar},
		time.Date(1990, time.February, 26, 12, 0, 0, 0, time.Local),
	}, {
		HebrewDate{Y: HebrewYear{Y: 5750}, D: 1, M: Elul},
		time.Date(1990, time.August, 22, 12, 0, 0, 0, time.Local),
	}, {
		HebrewDate{Y: HebrewYear{Y: 5750}, D: 29, M: Elul},
		time.Date(1990, time.September, 19, 12, 0, 0, 0, time.Local),
	}, {
		HebrewDate{Y: HebrewYear{Y: 5785}, D: 30, M: Shevat},
		time.Date(2025, time.February, 28, 12, 0, 0, 0, time.Local),
	}, {
		HebrewDate{Y: HebrewYear{Y: 5790}, D: 1, M: Shevat},
		time.Date(2030, time.January, 5, 12, 0, 0, 0, time.Local),
	}, {
		HebrewDate{Y: HebrewYear{Y: 5790}, D: 1, M: Adar_II},
		time.Date(2030, time.March, 6, 12, 0, 0, 0, time.Local),
	}, {
		HebrewDate{Y: HebrewYear{Y: 5784}, D: 29, M: Adar_II},
		time.Date(2024, time.April, 8, 12, 0, 0, 0, time.Local),
	}}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%v==%v", test.hd, dmy(test.t)), func(t *testing.T) {
			t.Run("ToHebrewDate", func(t *testing.T) {
				if got, want := ToHebrewDate(test.t), test.hd; !got.Equals(want) {
					t.Errorf("got %v, want %v", got, want)
				}
			})
			t.Run("FromHebrewDate", func(t *testing.T) {
				if got, want := FromHebrewDate(test.hd), test.t; !got.Equal(want) {
					t.Errorf("got %v, want %v", got, want)
				}
			})
		})
	}
}
