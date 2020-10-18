package conway

import "fmt"

type quality int

const (
	abundant  = 1
	regular   = 0
	deficient = -1
)

type hebrewYear struct {
	y        int // the year
	s        quality
	leapYear bool
}

// length() isn't used anywhere; perhaps expose a Hebrew Year .Length() API?
func (h *hebrewYear) length() int {
	// ref: p. 4
	days := 354 + int(h.s)
	if h.leapYear {
		days += 30
	}
	return days
}

// String implements stringer.String.
func (h hebrewYear) String() string {
	return fmt.Sprintf("%d", h.y)
}

func (h *hebrewYear) monthLength(m HebrewMonth) int {
	switch m {
	case Nissan, Sivan, Av, Tishrei, Shevat, Adar_I:
		return 30
	case Iyar, Tamuz, Elul, Tevet, Adar_II, Adar:
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
	Nissan HebrewMonth = 3 + iota
	Iyar
	Sivan
	Tamuz
	Av
	Elul
	Tishrei
	Marcheshvan
	Kislev
	Tevet
	Shevat
	Adar_I
	Adar_II
	Adar
)

func (m HebrewMonth) num() int {
	if m < Tishrei {
		return int(m)
	} else if m == Adar {
		return int(Adar_I) - 1
	}
	return int(m) - 1 // Tishrei and Elul are both height 8
}

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
	case Adar:
		return "Adar"
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
	Y int
	D int
	M HebrewMonth
}

type hebrewDate struct {
	y hebrewYear
	d int
	m HebrewMonth
}

// String implements stringer.String.
func (h HebrewDate) String() string {
	return fmt.Sprintf("%d %s %d", h.D, h.M, h.Y)
}

// Equal compares two HebrewDates for equality
func (h HebrewDate) Equal(d HebrewDate) bool {
	return h.D == d.D && h.M == d.M && h.Y == d.Y
}
