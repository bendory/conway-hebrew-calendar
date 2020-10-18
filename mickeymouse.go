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
	hebrewYears       [2]hebrewYear
	gregorianLeapYear bool
}

func gregorianMickeyMouse(year int) gmm {
	if year < 1 {
		panic("Can't go prior to 1 BCE.")
	}
	mm := gmm{
		hebrewYears: [2]hebrewYear{hebrewYear{y: year + 3760}, hebrewYear{y: year + 3761}},
	}

	var (
		dayFloat float64
		day      int
	)
	{
		// First compute the Roman date of RH; ref: p. 5.
		// Note that roshHashnnah computes a Gregorian September RH date, which may
		// be a squashed or stretched real date. The September date is needed to
		// compute IT.
		var b float64 // "bissextile" factor is earliest possible RH as a September date
		{
			y := year/100 - 11 // year is an int, so this is a floor
			mod := y%4 - 1
			if mod < 0 {
				mod = 0
			}
			b = float64(y/4*3 + mod) // y is an int, so y/4 is a floor
		}

		b += float64(year%4) / 4.0 // adjust "bissextile" time for Roman leap year
		f := (12 * (year%19 + 1)) % 19

		a := 1.5 * float64(f) // "acrobatic" term jumps from 0-27; how far RH falls from earliest possible RH
		d := float64(2*(year-1900)-1) / 35.0
		e := float64(f+1) / 760.0 // optionally ignore for 1762-2168
		dayFloat = a + b + (float64(f+1)-d)/18.0 - e
		day = int(dayFloat) // truncate, don't round! per david.slusky@ku.edu via email.

		// Now mark leap years.
		mm.gregorianLeapYear = year%4 == 0 && (year%100 != 0 || year%400 == 0)
		mm.hebrewYears[0].leapYear = 12 <= f && f <= 18
		mm.hebrewYears[1].leapYear = f <= 6
	}

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

	mm.it = day + 9     // IT is the day of RH as a September date + 9; ref: p. 3
	mm.he = mm.it + 29  // HE; ref: p. 3
	mm.she = mm.it + 10 // SHE; ref: p. 3
	if mm.gregorianLeapYear {
		mm.she++
	}
	// NOTE: It isn't clear from p. 3, but SHE depends on the outgoing year.
	if !mm.hebrewYears[0].leapYear {
		mm.she += 30
	}

	return mm
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

func (mm *gmm) height(d int, m time.Month) int {
	ht := d + int(m)
	if m < time.March {
		ht += 12
	}
	return ht
}

// hmm is a Hebrew mickeymouse; ref: p. 2.
type hmm struct {
	he, she, it int
	rh          time.Time // Gregorian date of Rosh Hashannah
	y           hebrewYear
}

func hebrewMickeyMouse(year int) hmm {
	gregorianRHyear := year - 3761
	thisGmm := gregorianMickeyMouse(gregorianRHyear)
	nextGmm := gregorianMickeyMouse(gregorianRHyear + 1)
	mm := hmm{
		he:  thisGmm.he,
		she: nextGmm.she,
		it:  nextGmm.it,
		rh:  thisGmm.rh,
		y: hebrewYear{
			y:        year,
			leapYear: thisGmm.hebrewYears[1].leapYear,
			s:        quality(nextGmm.she - thisGmm.he),
		},
	}
	return mm
}

func (mm *hmm) heSheIt(m HebrewMonth) int {
	switch m {
	case Tishrei, Marcheshvan:
		return mm.he
	case Kislev:
		return int(math.Max(float64(mm.he), float64(mm.she)))
	case Tevet, Shevat, Adar_I, Adar_II, Adar:
		return mm.she
	case Nissan, Iyar, Sivan, Tamuz, Av, Elul:
		return mm.it
	default:
		panic(fmt.Sprintf("Unknown month: %v", m))
	}
}

// partner converts a time.Month to its partner HebrewMonth and heSheIt value;
// ref: p. 2.
func (mm *hmm) partner(m time.Month, preferedAugustPartner HebrewMonth) (HebrewMonth, int) {
	switch m {
	case time.August:
		// Elul, mm.it and Tishrei, mm.he are both partners for August. While
		// it+29 = he, it is for this year and he is for next year -- so our
		// caller, who knows which year we are in (by knowing if we are before
		// or after RH) tells us which is prefered.
		if preferedAugustPartner == Tishrei {
			return Tishrei, mm.he
		}
		return Elul, mm.it
	case time.September:
		return Marcheshvan, mm.he
	case time.October:
		return Kislev, int(math.Max(float64(mm.he), float64(mm.she)))
	case time.November:
		return Tevet, mm.she
	case time.December:
		return Shevat, mm.she
	case time.January:
		if mm.y.leapYear {
			return Adar_I, mm.she
		}
		return Adar, mm.she
	case time.February:
		if mm.y.leapYear {
			return Adar_II, mm.she
		}
		return Adar, mm.she
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

type hebrewDate struct {
	y hebrewYear
	d int
	m HebrewMonth
}

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
