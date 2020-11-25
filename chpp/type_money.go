package chpp

import "github.com/dustin/go-humanize"

// Money represents the amount of money in SEK (swedish krona)
type Money int

// String returns a string representation of the type.
func (m Money) String() string {
	return humanize.Comma(int64(m)) + " kr"
}
