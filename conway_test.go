package conway

import (
	"testing"
	"time"
)

// TODO: get a table of years and add tests.
func TestRoshHashannah(t *testing.T) {
	y := 2019
	rh := roshHashannah(y)
	want := GregorianDate{y: y, d: 30, m: time.September}
	if rh.d != want.d || rh.m != want.m || rh.y != want.y {
		t.Errorf("got %v; want %v", rh, want)
	}
}
