package conway

import (
	"testing"
	"time"
)

// TODO: get a table of years and add tests.
func TestRoshHashannah(t *testing.T) {
	y := 2019
	c := newConway(y)
	want := GregorianDate{y: y, d: 30, m: time.September}
	got := c.rh
	if got.d != want.d || got.m != want.m || got.y != want.y {
		t.Errorf("got %v; want %v", got, want)
	}
}
