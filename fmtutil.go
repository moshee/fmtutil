// Package fmtutil implements formatting for numbers as common meaningful
// values.
package fmtutil

import (
	"fmt"
	"strconv"
	"time"
)

// SI represents an integer which can format itself with SI prefixes.
type SI uint64

const (
	K = 1024 << (10 * iota)
	M
	G
	T
	P
	E
)

func (x SI) String() string {
	n := 0.0
	s := ""
	switch {
	case x < K:
		return strconv.FormatUint(uint64(x), 10)
	case x < M:
		s = "k"
		n = float64(x) / K
	case x < G:
		s = "M"
		n = float64(x) / M
	case x < T:
		s = "G"
		n = float64(x) / G
	case x < P:
		s = "T"
		n = float64(x) / T
	case x < E:
		s = "P"
		n = float64(x) / P
	default:
		s = "E"
		n = float64(x) / E
	}

	return strconv.FormatFloat(Round(n, 1), 'f', -1, 64) + s
}

type SIUnit struct {
	N uint64
	U string
}

func (x SIUnit) String() string {
	return SI(x.N).String() + x.U
}

// Round rounds a number to the given number of total digits.
func Round(n float64, prec int) float64 {
	n *= float64(prec) * 10
	x := float64(int64(n + 0.5))
	return x / (float64(prec) * 10)
}

// Bytes is a common use case of SI prefix formatting.
type Bytes uint64

func (b Bytes) String() string {
	return SI(b).String() + "B"
}

const (
	Sec   = time.Second
	Min   = Sec * 60
	Hr    = Min * 60
	Day   = Hr * 24
	Week  = Day * 7
	Month = Day * 30
	Year  = Day * 365
)

// LongDuration formats a duration that is most likely much longer than what
// package time will handle. It uses the units seconds, minutes, hours, days,
// weeks, months (30 days), and years (365 days).
func LongDuration(n time.Duration) string {
	p := func(n time.Duration, s string) string {
		return fmt.Sprintf("%d%s", n, s)
	}

	switch {
	case n < Sec:
		return n.String()
	case n < Min:
		return p(n/Sec, "s")
	case n < Hr:
		return p(n/Min, "m")
	case n < Day:
		return p(n/Hr, "h")
	case n < 2*Week:
		return p(n/Day, "d")
	case n < Month:
		return p(n/Week, "w")
	case n < Year:
		return p(n/Month, "mo")
	default:
		return p(n/Year, "y")
	}
}
