package conway

import (
	"testing"
	"time"
)

// TODO: get a table of years and add tests.
func TestRoshHashannah(t *testing.T) {
	y := 2019
	m := newMickeyMouse(y)
	wantY, wantM, wantD := y, time.September, 30
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
}
