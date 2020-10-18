package conway

import (
	"fmt"
	"time"
)

func ToHebrewDate(t time.Time) HebrewDate {
	y, m, d := t.Date()
	gMM := gregorianMickeyMouse(y)

	var hMM hmm
	var preferedAugustPartner HebrewMonth
	switch {
	// Date is RH or later.
	case m > gMM.rh.Month(), m == gMM.rh.Month() && d >= gMM.rh.Day():
		hMM = hebrewMickeyMouse(gMM.hebrewYears[1].y)
		preferedAugustPartner = Tishrei
	// Date is before RH.
	default:
		hMM = hebrewMickeyMouse(gMM.hebrewYears[0].y)
		preferedAugustPartner = Elul
	}

	ht := gMM.height(d, m)
	hm, heSheIt := hMM.partner(m, preferedAugustPartner)

	// If height < heSheIt, then stretch...
	for heSheIt >= ht { // ref: p. 3
		m--
		if m < time.January {
			m = time.December
		}
		d += gMM.monthLength(m)
		ht = gMM.height(d, m)
		hm, heSheIt = hMM.partner(m, preferedAugustPartner)
	}
	hd := ht - heSheIt

	// Date extends into next month -- shrink...
	for hd > hMM.y.monthLength(hm) {
		hd -= hMM.y.monthLength(hm)
		hm++ // This won't work for Adar or Elul -- but we won't hit this code path in those months!
	}
	return HebrewDate{D: hd, M: hm, Y: hMM.y.y}
}

func FromHebrewDate(h HebrewDate) time.Time {
	mm := hebrewMickeyMouse(h.Y)
	heSheIt := mm.heSheIt(h.M)
	ht := h.D + heSheIt
	gm := time.Month(h.M.num())
	gd := ht - int(gm)
	if gm > time.December {
		gm -= 12
	}
	gy := mm.rh.Year()
	if h.M <= Elul || h.M > Shevat || gm == time.January {
		gy++
	}
	return time.Date(gy, gm, gd, 12, 0, 0, 0, time.Local)
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

// String implements stringer.String.
func (h HebrewDate) String() string {
	return fmt.Sprintf("%d %s %d", h.D, h.M, h.Y)
}

// Equal compares two HebrewDates for equality
func (h HebrewDate) Equal(d HebrewDate) bool {
	return h.D == d.D && h.M == d.M && h.Y == d.Y
}
