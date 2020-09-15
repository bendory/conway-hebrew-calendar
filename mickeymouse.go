package conway

import (
	"fmt"
	"math"
	"time"
)

// TODO: This is actually a Gregorian mickeymouse; we also need to be able to make a
// Hebrew mickeymouse. See p. 2.
type mickeymouse struct {
	he, she, it int
	rh          *GregorianDate // Gregorian date of Rosh Hashannah
	hebrewYears [2]int
}

func newMickeyMouse(year int) mickeymouse {
	// compute all the needed values for calendar conversions.
	mm := mickeymouse{}
	mm.hebrewYears[0], mm.hebrewYears[1] = year+3760, year+3761

	// First compute the Roman date of RH; ref: p. 5.
	// Note that roshHashnnah computes an un-squashed Gregorian date, thereby
	// considering RH as a September date, which is what is needed to compute
	// IT.
	var b float64 // "bissextile" time; earliest possible RH
	switch {
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
	g := y%19 + 1
	f := float64((12 * g) % 19)

	a := 1.5 * float64(f) // "acrobatic" term jumps from 0-27; how far RH falls from earliest possible RH
	c := f + 1.0
	d := (2.0*float64(y) - 1.0) / 35.0
	e := (f + 1.0) / 760.0 // can be ignored for 1762-2168
	day := int(math.Round(a + b + (c-d-e)/18.0))

	// Now mark leap years.
	isLeapYear := f <= 6
	priorWasLeapYear := 12 <= f && f <= 18
	_ = priorWasLeapYear // TODO: What is this for?
	gregorianLeapYear := year%4 == 0 && (year%100 != 0 || year%400 == 0)

	// IT for the given date; ref: p. 3
	mm.it = day + 9

	// HE; ref: p. 3
	mm.he = mm.it + 29

	// SHE; ref: p. 3
	switch {
	case isLeapYear && !gregorianLeapYear:
		mm.she = mm.it + 10
	case isLeapYear && gregorianLeapYear:
		mm.she = mm.it + 11
	case isLeapYear && !gregorianLeapYear:
		mm.she = mm.it + 40
	case isLeapYear && gregorianLeapYear:
		mm.she = mm.it + 41
	}
	mm.rh = &GregorianDate{time.Date(year, time.September, day, 12, 0, 0, 0, time.Local)}
	return mm
}

type partners struct {
	h HebrewMonth
	m time.Month
	n int
}

type partnerList []partners

var allPartners = partnerList{
	{Tishrei, time.August, 8},
	{Marcheshvan, time.September, 9},
	{Kislev, time.October, 10},
	{Tevet, time.November, 11},
	{Shevat, time.December, 12},
	{Adar_I, time.January, 13},
	{Adar_II, time.February, 14},
	{Nissan, time.March, 3},
	{Iyar, time.April, 4},
	{Sivan, time.May, 5},
	{Tamuz, time.June, 6},
	{Av, time.July, 7},
	{Elul, time.August, 8},
}

var he = allPartners[0:2]
var she = allPartners[2:6]
var it = allPartners[6:]
