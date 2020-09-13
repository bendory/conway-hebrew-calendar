package conway

import (
	"fmt"
	"math"
	"time"
)

type conway struct {
	he, she, it int
	rh          GregorianDate // Gregorian date of Rosh Hashannah
}

func newConway(year int) conway {
	c := conway{rh: GregorianDate{m: time.September, y: year}}
	c.compute()
	return c
}

// compute all the needed values for calendar conversions.
func (cwy *conway) compute() {
	// First compute the Roman date of RH; ref: p. 5.
	// Note that roshHashnnah computes an un-squashed Gregorian date, thereby
	// considering RH as a September date, which is what is needed to compute
	// IT.
	var b float64 // "bissextile" time; earliest possible RH
	switch {
	case cwy.rh.y >= 1500 && cwy.rh.y < 1700:
		b = 3.0 // Earliest possible RH ~Sept 3
	case cwy.rh.y >= 1700 && cwy.rh.y < 1800:
		b = 4.0 // ~Sept 4
	case cwy.rh.y >= 1800 && cwy.rh.y < 1900:
		b = 5.0
	case cwy.rh.y >= 1900 && cwy.rh.y < 2100:
		b = 6.0
	case cwy.rh.y >= 2100 && cwy.rh.y < 2200:
		b = 7.0
	case cwy.rh.y >= 2200 && cwy.rh.y < 2300:
		b = 8.0
	case cwy.rh.y >= 2300 && cwy.rh.y < 2400:
		b = 9.0
	default:
		// TODO: expand valid years.
		panic(fmt.Sprintf("Rosh Hashannah can only be calculated for 1500-2400, not %d.", cwy.rh.y))
	}
	b += float64(cwy.rh.y%4) / 4.0 // adjust "bissextile" time for Roman leap year

	y := cwy.rh.y - 1900
	g := y%19 + 1
	f := float64((12 * g) % 19)

	a := 1.5 * float64(f) // "acrobatic" term jumps from 0-27; how far RH falls from earliest possible RH
	c := f + 1.0
	d := (2.0*float64(y) - 1.0) / 35.0
	e := (f + 1.0) / 760.0 // can be ignored for 1762-2168
	cwy.rh.d = int(math.Round(a + b + (c-d-e)/18.0))

	// Now mark leap years.
	isLeapYear := f <= 6
	priorWasLeapYear := 12 <= f && f <= 18
	_ = priorWasLeapYear // TODO: What is this for?
	gregorianLeapYear := cwy.rh.y%4 == 0 && (cwy.rh.y%100 != 0 || cwy.rh.y%400 == 0)

	// IT for the given date; ref: p. 3
	cwy.it = cwy.rh.d + 9

	// HE; ref: p. 3
	cwy.he = cwy.it + 29

	// SHE; ref: p. 3
	switch {
	case isLeapYear && !gregorianLeapYear:
		cwy.she = cwy.it + 10
	case isLeapYear && gregorianLeapYear:
		cwy.she = cwy.it + 11
	case isLeapYear && !gregorianLeapYear:
		cwy.she = cwy.it + 40
	case isLeapYear && gregorianLeapYear:
		cwy.she = cwy.it + 41
	}
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
