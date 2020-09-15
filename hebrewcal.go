package conway

import (
	"fmt"
	"time"
)

type HebrewMonth int

const (
	Tishrei = 1 + iota
	Marcheshvan
	Kislev
	Tevet
	Shevat
	Adar_I
	Adar_II
	Nissan
	Iyar
	Sivan
	Tamuz
	Av
	Elul
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
	y, d int
	m    HebrewMonth
}

func NewHebrewDate(t time.Time) HebrewDate {
	// TODO: convert from t to GregorianDate and from there to HebrewDate.
	return HebrewDate{}
}

// String implements stringer.String.
func (h HebrewDate) String() string {
	return fmt.Sprintf("%d %s %4d", h.d, h.m, h.y)
}

// Format prints the date based on time.Time.Format layouts.
func (h HebrewDate) Format(layout string) string {
	// TODO: implement Format for Hebrew dates.
	return "unimplemented"
}

// Time returns a time.Time object corresponding to this Date.
func (h HebrewDate) Time() time.Time {
	// TODO: convert to GregorianDate return time.Time.
	return time.Time{}
}

// height gives the "height" of the date, per Conway.
func (h HebrewDate) height() int {
	ht := h.d + int(h.m)
	return ht
}
