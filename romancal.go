package conway

import (
	"time"
)

type GregorianDate struct {
	time.Time
}

// NewGregorianDate is a convenience function to convert from y-m-d to date.
func NewGregorianDate(y int, m time.Month, d int) GregorianDate {
	return GregorianDate{time.Date(y, m, d, 12, 0, 0, 0, time.Local)}
}

// String implements stringer.String to print in default format "2 January 2006".
func (g GregorianDate) String() string {
	return g.Format("2 January 2006")
}

// height gives the "height" of the date, per Conway.
func (g GregorianDate) height() int {
	_, m, d := g.Date()
	h := d + int(m)
	if m <= time.February {
		h += 12
	}
	return h
}
