package conway

import (
	"fmt"
	"math"
	"time"
)

// gmm is a Gregorian mickeymouse; ref: p. 2.
type gmm struct {
	he, she, it       int
	rh                time.Time // Gregorian date of Rosh Hashannah
	hebrewYears       [2]HebrewYear
	gregorianLeapYear bool
}

func gregorianMickeyMouse(year int) gmm {
	mm := gmm{
		hebrewYears: [2]HebrewYear{HebrewYear{y: year + 3760}, HebrewYear{y: year + 3761}},
	}

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
	dayFloat := a + b + (float64(c)-d)/18.0 - e
	day := int(dayFloat) // truncate, don't round! per david.slusky@ku.edu via email.

	// Now mark leap years.
	mm.gregorianLeapYear = year%4 == 0 && (year%100 != 0 || year%400 == 0)
	mm.hebrewYears[0].leapYear = 12 <= f && f <= 18
	mm.hebrewYears[1].leapYear = f <= 6

	// We now know rh... unless rh must be postponed...
	mm.rh = time.Date(year, time.September, day, 12, 0, 0, 0, time.Local)
	switch mm.rh.Weekday() {
	case time.Sunday, time.Wednesday, time.Friday: // ref: p. 6 (לא אד״ו ראש)
		day++
		mm.rh = time.Date(year, time.September, day, 12, 0, 0, 0, time.Local)
	case time.Tuesday: // Third דחיה; ref: p. 7
		if !mm.hebrewYears[1].leapYear && dayFloat-float64(day) > .633 {
			// day+2 shortens this year from 356 --> 354 days and implies that
			// the prior year was longer.
			day += 2
			mm.rh = time.Date(year, time.September, day, 12, 0, 0, 0, time.Local)
		}
	case time.Monday: // Fourth דחיה; ref: p. 7
		if mm.hebrewYears[0].leapYear && dayFloat-float64(day) > .898 {
			// day+1 lengthens the prior year from 382 --> 383 days
			day++
			mm.rh = time.Date(year, time.September, day, 12, 0, 0, 0, time.Local)
		}
	}

	// IT is the day of RH as a September date + 9; ref: p. 3
	mm.it = day + 9

	// HE; ref: p. 3
	mm.he = mm.it + 29

	// SHE; ref: p. 3
	mm.she = mm.it + 10
	if mm.gregorianLeapYear {
		mm.she++
	}
	// NOTE: It isn't clear from p. 3, but SHE depends on the outgoing year.
	if !mm.hebrewYears[0].leapYear {
		mm.she += 30
	}

	// TODO: the right way to calculate
	// s = (nextHebrewYear.she - thisHebrewYear.he)
	mm.validate()
	return mm
}

// partner converts a time.Month to its partner HebrewMonth and heSheIt value;
// ref: p. 2.
func (mm *gmm) partner(m time.Month) (HebrewMonth, int) {
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

func (mm *gmm) monthLength(m time.Month) int {
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

func (mm *gmm) validate() {
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

func (mm *gmm) height(d int, m time.Month) int {
	ht := d + int(m)
	if m < time.March {
		ht += 12
	}
	return ht
}

func ToHebrewDate(t time.Time) HebrewDate {
	y, m, d := t.Date()
	mm := gregorianMickeyMouse(y)
	ht := mm.height(d, m)
	hm, heSheIt := mm.partner(m)

	// If height < heSheIt, then stretch...
	for heSheIt > ht { // ref: p. 3
		m--
		if m < time.January {
			m = time.December
		}
		d += mm.monthLength(m)
		ht = mm.height(d, m)
		hm, heSheIt = mm.partner(m)
	}
	hd := ht - heSheIt
	var hy HebrewYear
	if mm.rh.After(t) {
		hy = mm.hebrewYears[0] // before rh we're in the prior year
	} else {
		hy = mm.hebrewYears[1] // after rh we're in the next year
	}

	return HebrewDate{d: hd, m: hm, y: hy}
}

func FromHebrewDate(h HebrewDate) time.Time {
	mm := hebrewMickeyMouse(h.y.y)
	heSheIt := mm.heSheIt(h.m)
	ht := h.d + heSheIt
	gm := time.Month(h.m.num())
	gd := ht - int(gm)
	if gm > time.December {
		gm -= 12
	}
	gy := mm.rh.Year()
	if h.m <= Elul || gm == time.January {
		gy++
	}
	return time.Date(gy, gm, gd, 12, 0, 0, 0, time.Local)
}

// hmm is a Hebrew mickeymouse; ref: p. 2.
type hmm struct {
	he, she, it int
	rh          time.Time // Gregorian date of Rosh Hashannah
	y           HebrewYear
}

func hebrewMickeyMouse(year int) hmm {
	gregorianRHyear := year - 3761
	thisGmm := gregorianMickeyMouse(gregorianRHyear)
	nextGmm := gregorianMickeyMouse(gregorianRHyear + 1)
	if thisGmm.hebrewYears[1].y != nextGmm.hebrewYears[0].y {
		panic(fmt.Sprintf("Hebrew year mismatch: %d != %d", thisGmm.hebrewYears[1].y, nextGmm.hebrewYears[0].y))
	}
	if thisGmm.hebrewYears[1].leapYear != nextGmm.hebrewYears[0].leapYear {
		panic(fmt.Sprintf("Hebrew leapYear mismatch: %t != %t", thisGmm.hebrewYears[1].leapYear, nextGmm.hebrewYears[0].leapYear))
	}
	mm := hmm{
		he:  thisGmm.he,
		she: nextGmm.she,
		it:  nextGmm.it,
		rh:  thisGmm.rh,
		y: HebrewYear{
			y:        year,
			leapYear: thisGmm.hebrewYears[1].leapYear,
			s:        quality(nextGmm.she - thisGmm.he),
		},
	}
	mm.validate()
	return mm
}

func (mm *hmm) heSheIt(m HebrewMonth) int {
	switch m {
	case Tishrei, Marcheshvan:
		return mm.he
	case Kislev:
		return int(math.Max(float64(mm.he), float64(mm.she)))
	case Tevet, Shevat, Adar_I, Adar_II:
		return mm.she
	case Nissan, Iyar, Sivan, Tamuz, Av, Elul:
		return mm.it
	default:
		panic(fmt.Sprintf("Unknown month: %v", m))
	}
}

func (mm *hmm) validate() {
	switch mm.y.s {
	case abundant, regular, deficient:
	default:
		panic(fmt.Sprintf("s = %d!?", mm.y.s))
	}
}
