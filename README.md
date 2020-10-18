# conway-hebrew-calendar
An implementation of Professor John H. Conway's Hebrew Calendar algorithm

## Background

Professor Conway was able to convert between Hebrew and Gregorian calendar dates
by way of mental math in seconds. After his death, I learned that a
[paper](https://slusky.ku.edu/wp-content/uploads/2020/08/CONWAY-AGUS-SLUSKY-PDF.pdf)
explaining his algorthim had been published in the January 2014 edition of The
College Mathematics Journal. (For archival purposes, I've copied the PDF
[here](pdf/conway-agus-slusky.pdf).)

This repo implements Conway's Hebrew Calendar algorithm in Go.

## Ambiguities

Along the implementation path, I found a few notable items that are not explained in
the paper.

1. On p. 3, when calculating the date of Rosh Hashannah, _truncate_, don't
   round.  This seems to contradict the paper which states that Sept 29.5 -->
   Sept 30, but author David Slusky told me to truncate and empirical evidence
   is that truncating works.
1. On p. 3, when calculating `SHE`, use `4` in the tens digit if the _outgoing_
   year was a leap year.
