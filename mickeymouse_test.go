package conway

import (
	"fmt"
	"testing"
	"time"
)

// TODO: not all tests pass.
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
		t.Run(fmt.Sprint(test.hebrewYear), func(t *testing.T) {
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
			if got, want := m.hebrewYears[1].y, test.hebrewYear; got != want {
				t.Errorf("got %d; want %d", got, want)
			}
			if gotWD := m.rh.Weekday(); gotWD == time.Sunday || gotWD == time.Wednesday || gotWD == time.Friday {
				t.Errorf("לא אד״ו ראש: %s", gotWD)
			}
		})
	}
}

func TestMickeyMouse(t *testing.T) {
	tests := []struct{ he, she, it, year int }{
		{71, 53, 42, 2016},
		{59, 70, 30, 2017},
		{48, 59, 19, 2018},
		{68, 49, 39, 2019},
	}
	for _, test := range tests {
		t.Run(fmt.Sprint(test.year), func(t *testing.T) {
			m := gregorianMickeyMouse(test.year)
			if gotHe, wantHe := m.he, test.he; gotHe != wantHe {
				t.Errorf("HE:  got=%d want=%d", gotHe, wantHe)
			}
			if gotShe, wantShe := m.she, test.she; gotShe != wantShe {
				t.Errorf("SHE: got=%d want=%d", gotShe, wantShe)
			}
			if gotIt, wantIt := m.it, test.it; gotIt != wantIt {
				t.Errorf("IT:  got=%d want=%d", gotIt, wantIt)
			}
		})
	}
}

// TODO: add test cases
// 25 Kislev 5777 == 25 December 2016 p.1
// 7 Iyar 5779 == 12 May 2019 p.3
// 7 April 2019 == 2 Nissan 5779 p. 3
// 30 September 2019 == 1 Tishrei 5780 p. 3
