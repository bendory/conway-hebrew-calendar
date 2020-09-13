package conway

import (
	"time"
)

type GregorianDate struct {
	y, d int
	m    time.Month
}

func NewGregorianDate(t time.Time) GregorianDate {
	return GregorianDate{y: t.Year(), m: t.Month(), d: t.Day()}
}

// String implements stringer.String.
func (g GregorianDate) String() string {
	return g.Format("2 January 2006")
}

// Format prints the date based on time.Time.Format layouts.
func (g GregorianDate) Format(layout string) string {
	return g.Time().Format(layout)
}

// Time returns a time.Time object corresponding to this GregorianDate.
func (g GregorianDate) Time() time.Time {
	return time.Date(g.y, time.Month(g.m), g.d, 12, 0, 0, 0, time.Local)
}

// height gives the "height" of the date, per Conway.
func (g GregorianDate) height() int {
	h := g.d + int(g.m)
	if g.m <= time.February {
		h += 12
	}
	return h
}

// squash fixes invalid dates to valid dates
func (g *GregorianDate) squash() {
	d := g.Time()
	g.y, g.m, g.d = d.Year(), d.Month(), d.Day()
}
