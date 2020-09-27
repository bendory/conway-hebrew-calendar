package conway

import (
	"fmt"
	"time"
)

type quality int

const (
	abundant  = 1
	regular   = 0
	deficient = -1
)

type HebrewYear struct {
	y        int // the year
	s        quality
	leapYear bool
}

func (h *HebrewYear) Length() int {
	// ref: p. 4
	days := 354 + int(h.s)
	if h.leapYear {
		days += 30
	}
	return days
}

// String implements stringer.String.
func (h HebrewYear) String() string {
	return fmt.Sprintf("%d", h.y)
}

func (h *HebrewYear) monthLength(m HebrewMonth) int {
	switch m {
	case Nissan, Sivan, Av, Tishrei, Shevat, Adar_I:
		return 30
	case Iyar, Tamuz, Elul, Tevet, Adar_II:
		return 29
	case Marcheshvan:
		if h.s == 1 { // ref: p. 4
			return 30
		}
		return 29
	case Kislev:
		if h.s == -1 { // ref: p. 4
			return 29
		}
		return 30
	default:
		panic(fmt.Sprint("Invalid month:", m))
	}
}

type HebrewMonth int

const (
	Nissan = 3 + iota
	Iyar
	Sivan
	Tamuz
	Av
	Elul
	Tishrei = 8 + iota
	Marcheshvan
	Kislev
	Tevet
	Shevat
	Adar_I
	Adar_II
)

func (m HebrewMonth) String() string {
	switch m {
	case Tishrei:
		return "Tishrei"
	case Marcheshvan:
		return "Marcheshvan"
	case Kislev:
		return "Kislev"
	case Tevet:
		return "Tevet"
	case Shevat:
		return "Shevat"
	case Adar_I:
		return "Adar_I"
	case Adar_II:
		return "Adar_II"
	case Nissan:
		return "Nissan"
	case Iyar:
		return "Iyar"
	case Sivan:
		return "Sivan"
	case Tamuz:
		return "Tamuz"
	case Av:
		return "Av"
	case Elul:
		return "Elul"
	}
	panic(fmt.Sprintf("No known Hebrew month %d", m))
}

type HebrewDate struct {
	y HebrewYear
	d int
	m HebrewMonth
}

func NewHebrewDate(t time.Time) HebrewDate {
	return ToHebrewDate(t)
}

// String implements stringer.String.
func (h HebrewDate) String() string {
	return fmt.Sprintf("%d %s %s", h.d, h.m, h.y)
}

// Time returns a time.Time object corresponding to this Date.
func (h HebrewDate) Time() time.Time {
	// TODO: convert to GregorianDate return time.Time.
	return time.Time{}
}

// height gives the "height" of the date, per Conway.
func (h HebrewDate) height() int {
	return h.d + int(h.m)
}

// equals compares two HebrewDates for equality
func (h HebrewDate) Equals(d HebrewDate) bool {
	return h.d == d.d && h.m == d.m && h.y.y == d.y.y
}
