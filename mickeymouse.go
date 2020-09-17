package conway

import (
	"fmt"
	"math"
	"time"
)

// partner converts a time.Month to its partner HebrewMonth and heSheIt value;
// ref: p. 2.
func (mm *mickeymouse) partner(m time.Month) (HebrewMonth, int) {
	switch m {
	case time.August:
		return Tishrei, mm.he
		// Note: Elul, mm.it is an alternative partner for August. It doesn't
		// actually matter because it+29 = he, so that stretches the Elul date
		// into Tishrei or vice versa.
	case time.September:
		return Marcheshvan, mm.he
	case time.October:
		return Kislev, int(math.Max(float64(mm.he), float64(mm.she)))
	case time.November:
		return Tevet, mm.she
	case time.December:
		return Shevat, mm.she
	case time.January:
		return Adar_I, mm.she
	case time.February:
		return Adar_II, mm.she
	case time.March:
		return Nissan, mm.it
	case time.April:
		return Iyar, mm.it
	case time.May:
		return Sivan, mm.it
	case time.June:
		return Tamuz, mm.it
	case time.July:
		return Av, mm.it
	default:
		panic(fmt.Sprint("Unknown month:", m))
	}
}

func ToHebrewDate(date GregorianDate) HebrewDate {
	ht := date.height()
	y, m, d := date.Date()

	mm := gregorianMickeyMouse(y)
	hm, heSheIt := mm.partner(m)

	// If height < heSheIt, then stretch...
	if heSheIt > ht { // ref: p. 3
		d += mm.monthLength(m)
		m--
		if m < time.January {
			m = time.December
		}
		ht = d + int(m)
		if m <= time.February {
			ht += 12
		}
		hm, heSheIt = mm.partner(m)
	}
	hd := ht - heSheIt
	var hy HebrewYear
	if mm.rh.After(date.Time) {
		hy = mm.hebrewYears[0] // before rh we're in the prior year
	} else {
		hy = mm.hebrewYears[1] // after rh we're in the next year
	}

	return HebrewDate{d: hd, m: hm, y: hy}
}

type mousetype int8

const (
	gregorianMouse = 1
	hebrewMouse    = 2
)

// TODO: This is actually a Gregorian mickeymouse; we also need to be able to make a
// Hebrew mickeymouse. See p. 2.
type mickeymouse struct {
	he, she, it       int
	rh                *GregorianDate // Gregorian date of Rosh Hashannah
	hebrewYears       [2]HebrewYear
	mt                mousetype
	gregorianLeapYear bool
}

func gregorianMickeyMouse(gregorianYear int) mickeymouse {
	year := gregorianYear
	// compute all the needed values for calendar conversions.
	mm := mickeymouse{mt: gregorianMouse}
	mm.hebrewYears[0], mm.hebrewYears[1] = HebrewYear{y: year + 3760}, HebrewYear{y: year + 3761}

	// First compute the Roman date of RH; ref: p. 5.
	// Note that roshHashnnah computes an un-squashed Gregorian date, thereby
	// considering RH as a September date, which is what is needed to compute
	// IT.
	var b float64 // "bissextile" time; earliest possible RH
	switch {
	// b adjusts by centuries; ref: p. 8
	case year >= 1500 && year < 1700:
		b = 3.0 // Earliest possible RH ~Sept 3
	case year >= 1700 && year < 1800:
		b = 4.0 // ~Sept 4
	case year >= 1800 && year < 1900:
		b = 5.0
	case year >= 1900 && year < 2100:
		b = 6.0
	case year >= 2100 && year < 2200:
		b = 7.0
	case year >= 2200 && year < 2300:
		b = 8.0
	case year >= 2300 && year < 2400:
		b = 9.0
	default:
		// TODO: expand valid years.
		panic(fmt.Sprintf("Rosh Hashannah can only be calculated for 1500-2400, not %d.", year))
	}
	b += float64(year%4) / 4.0 // adjust "bissextile" time for Roman leap year

	y := year - 1900
	g := year%19 + 1
	f := (12 * g) % 19

	a := 1.5 * float64(f) // "acrobatic" term jumps from 0-27; how far RH falls from earliest possible RH
	c := f + 1
	d := float64(2*y-1) / 35.0
	e := float64(f+1) / 760.0 // can be ignored for 1762-2168
	dayFloat := a + b + (float64(c)-d-e)/18.0
	day := int(math.Round(dayFloat))

	// Now mark leap years.
	isLeapYear := f <= 6
	priorWasLeapYear := 12 <= f && f <= 18
	mm.gregorianLeapYear = year%4 == 0 && (year%100 != 0 || year%400 == 0)
	mm.hebrewYears[0].leapYear = priorWasLeapYear
	mm.hebrewYears[1].leapYear = isLeapYear

	// IT is the day of RH as a September date + 9; ref: p. 3
	mm.it = day + 9

	// HE; ref: p. 3
	mm.he = mm.it + 29

	// SHE; ref: p. 3
	mm.she = mm.it + 10
	if mm.gregorianLeapYear {
		mm.she++
	}
	// NOTE: It isn't clear from p. 3, but SHE depends on priorWasLeapYear, not
	// isLeapYear.
	if !priorWasLeapYear {
		mm.she += 30
	}

	// TODO: the right way to calculate
	// s = (nextHebrewYear.she - thisHebrewYear.he)

	// We now know rh... unless rh must be postponed...
	mm.rh = &GregorianDate{time.Date(year, time.September, day, 12, 0, 0, 0, time.Local)}
	switch mm.rh.Weekday() {
	case time.Sunday, time.Wednesday, time.Friday: // ref: p. 6 (לא אד״ו ראש)
		mm.rh = &GregorianDate{time.Date(year, time.September, day+1, 12, 0, 0, 0, time.Local)}
	case time.Tuesday: // Third דחיה; ref: p. 7
		if !isLeapYear && dayFloat-float64(day) > .633 {
			// day+2 shortens this year from 356 --> 354 days and implies that
			// the prior year was longer.
			mm.rh = &GregorianDate{time.Date(year, time.September, day+2, 12, 0, 0, 0, time.Local)}
		}
	case time.Monday: // Fourth דחיה; ref: p. 7
		if priorWasLeapYear && dayFloat-float64(day) > .898 {
			// day+1 lengthens the prior year from 382 --> 383 days
			mm.rh = &GregorianDate{time.Date(year, time.September, day+1, 12, 0, 0, 0, time.Local)}
		}
	}
	mm.validate()
	return mm
}

func (mm *mickeymouse) monthLength(m time.Month) int {
	switch m {
	case time.January, time.March, time.May, time.July, time.August, time.October, time.December:
		return 31
	case time.April, time.June, time.September, time.November:
		return 30
	case time.February:
		if mm.gregorianLeapYear {
			return 29
		}
		return 28
	default:
		panic(fmt.Sprint("Unknown month:", m))
	}
}

func (mm *mickeymouse) validate() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("%#v\n", mm)
		}
	}()
	// p. 2
	if mm.it < 12 || mm.it > 44 {
		panic(fmt.Sprintf("12<=IT<=44: IT==%d", mm.it))
	}
	if mm.he < 41 || mm.he > 73 {
		panic(fmt.Sprintf("41<=HE<=73: HE==%d", mm.he))
	}
	if mm.she < 41 || mm.she > 73 {
		panic(fmt.Sprintf("41<=SHE<=73: SHE==%d", mm.she))
	}
	if mm.it >= mm.she {
		panic(fmt.Sprintf("IT<SHE: IT==%d SHE==%d", mm.it, mm.she))
	}
	if mm.it >= mm.he {
		panic(fmt.Sprintf("IT<HE: IT==%d HE==%d", mm.it, mm.he))
	}
}
