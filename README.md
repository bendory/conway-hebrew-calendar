# conway-hebrew-calendar
An implementation of
[Professor John H. Conway](https://www.princeton.edu/news/2020/04/14/mathematician-john-horton-conway-magical-genius-known-inventing-game-life-dies-age)'s
Hebrew Calendar algorithm.

## Background

Professor Conway was able to convert between Hebrew and Gregorian calendar dates
by way of mental math in seconds. After his death, I learned that a
[paper](https://slusky.ku.edu/wp-content/uploads/2020/08/CONWAY-AGUS-SLUSKY-PDF.pdf)
explaining his algorthim had been published in the January 2014 edition of The
College Mathematics Journal. (For archival purposes, I've copied the PDF
[here](pdf/conway-agus-slusky.pdf).)

This repo implements Conway's Hebrew Calendar algorithm in Go.

## Accuracy

I've tested this implementation for accuracy using dates from 1 CE - 3000 CE by
comparing to the Hebrew calendar behind http://www.hebcal.com/. Note that:

*  Results prior to the adoption of the Gregorian calendar (~1752, depending on
   location) use the [proleptic Gregorian
   calendar](https://en.wikipedia.org/wiki/Proleptic_Gregorian_calendar). This
   implementation is in Golang, and my understanding of Go's
   [time.Time](https://golang.org/src/time/time.go?s=6278:7279#L117) is that it
   similarly uses a proleptic Gregorian calendar.
*  Some time between 70 and 1178 CE, the observation-based Hebrew calendar was
   replaced by [the calculated calendar developed by Hillel
   HaNasi](https://en.wikipedia.org/wiki/Hillel_II#Fixing_of_the_calendar). So
   this implementation projects a proleptic calculated Hillel Hebrew calendar
   (a term that I just made up) onto a proleptic Gregorian calendar.

## Ambiguities

Along the implementation path, I found a few notable items that are not explained in
the paper.

1. On p. 3, when calculating the date of Rosh Hashannah, _truncate_, don't
   round.  This seems to contradict the paper which states that Sept 29.5 -->
   Sept 30, but author David Slusky told me to truncate and empirical evidence
   is that truncating works.
1. On p. 3, when calculating `SHE`, use `4` in the tens digit if the _outgoing_
   year was a leap year.
