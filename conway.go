package conway

import (
	"fmt"
	"math"
	"time"
)

// roshHashnnah computes the Roman date of RH; ref: p. 5.
func roshHashannah(y int) GregorianDate {
	var b float64 // "bissextile" time; earliest possible RH
	switch {
	case y >= 1500 && y < 1700:
		b = 3.0 // Earliest possible RH ~Sept 3
	case y >= 1700 && y < 1800:
		b = 4.0 // ~Sept 4
	case y >= 1800 && y < 1900:
		b = 5.0
	case y >= 1900 && y < 2100:
		b = 6.0
	case y >= 2100 && y < 2200:
		b = 7.0
	case y >= 2200 && y < 2300:
		b = 8.0
	case y >= 2300 && y < 2400:
		b = 9.0
	default:
		// TODO: expand valid years.
		panic(fmt.Sprintf("Rosh Hashannah can only be calculated for 1500-2400, not %d.", y))
	}
	b += float64(y%4) / 4.0 // adjust "bissextile" time for Roman leap year

	y -= 1900
	g := y%19 + 1
	f := float64((12 * g) % 19)

	isHebrewLeapYear := f <= 6
	previousWasLeapYear := 12 <= f && f <= 18
	_, _ = isHebrewLeapYear, previousWasLeapYear // Do I need to use these anywhere?

	a := 1.5 * float64(f) // "acrobatic" term jumps from 0-27; how far RH falls from earliest possible RH
	c := f + 1.0
	d := (2.0*float64(y) - 1.0) / 35.0
	e := (f + 1.0) / 760.0 // can be ignored for 1762-2168
	dayOfMonth := int(math.Round(a + b + (c-d-e)/18.0))
	rh := GregorianDate{m: time.September, d: dayOfMonth, y: y + 1900}
	rh.squash()
	return rh
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
