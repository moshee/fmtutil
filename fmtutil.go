package fmtutil

import "strconv"

// SI represents an integer which can format itself with SI letters.
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
