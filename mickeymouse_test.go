package conway

import (
	"fmt"
	"testing"
	"time"
)

// TODO: get a table of years and add tests.
func TestRoshHashannah(t *testing.T) {
	tests := []struct{ m, d, y, hebrewYear int }{
		// Cover a 19-year cycle...
		{9, 10, 2018, 5779},
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
				t.Errorf("Year: got %d; want %d", gotY, wantY)
			}
			if gotM != wantM {
				t.Errorf("Month: got %s; want %s", gotM, wantM)
			}
			if gotD != wantD {
				t.Errorf("Day: got %d; want %d", gotD, wantD)
			}
			if got, want := m.hebrewYears[1], test.hebrewYear; got != want {
				t.Errorf("got %d; want %d", got, want)
			}
		})
	}
}

// TODO: add test cases
// 25 Kislev 5777 == 25 December 2016 p.1
// 7 Iyar 5779 == 12 May 2019 p.3
// 7 April 2019 == 2 Nissan 5779 p. 3
// 30 September 2019 == 1 Tishrei 5780 p. 3
