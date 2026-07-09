package chpp

import "github.com/dustin/go-humanize"

// Money represents an amount of money in SEK (Swedish krona), as used
// throughout Hattrick (team finances, transfers, bids, etc).
type Money int

// String returns the amount humanized with thousands separators and a
// "kr" suffix, e.g. "1,234,567 kr".
func (m Money) String() string {
	return humanize.Comma(int64(m)) + " kr"
}
