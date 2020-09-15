package conway

import (
	"testing"
	"time"
)

// TODO: get a table of years and add tests.
func TestRoshHashannah(t *testing.T) {
	y := 2019
	m := newMickeyMouse(y)
	want := GregorianDate{y: y, d: 30, m: time.September}
	got := m.rh
	if got.d != want.d || got.m != want.m || got.y != want.y {
		t.Errorf("got %v; want %v", got, want)
	}
}
