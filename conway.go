package conway

import (
	"math"
	"time"
)

// roshHashnnah computes the Roman date of RH; ref: p. 5.
func roshHashannah(y int) GregorianDate {
	y -= 1900
	g := y%19 + 1
	f := float32((12 * g) % 19)

	isHebrewLeapYear := f <= 6
	previousWasLeapYear := 12 <= f && f <= 18

	a := 1.5 * float32(f)       // "acrobatic" term jumps from 0-27; how far RH falls from earliest possible RH
	b := 6.0 + float32(y%4)/4.0 // bissextile time; earliest possible RH (~Sept 6) + adjusts for Roman leap year
	c := f + 1.0
	d := (2.0*float32(y) - 1.0) / 35.0
	e := (f + 1.0) / 760.0 // can be ignored for 1762-2168
	dayOfMonth := int(math.Round(float64(a + b + (c-d-e)/18.0)))
	rh := GregorianDate{m: time.September, d: dayOfMonth, y: y + 1900}
	rh.squash()
	_, _ = isHebrewLeapYear, previousWasLeapYear
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
