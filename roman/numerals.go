package roman

import (
	"strings"
	"fmt"
)

func Encode(n uint) string {
	numerals := ""

	// Thousands
	if (n > 1000) {
		thousands := n / 1000
		numerals += strings.Repeat("M", int(thousands))
		n -= thousands * 1000
	}

	// Hundreds
	if (n > 100 && n <= 1000) {
		hundreds := n / 100
		numerals += toNumerals(hundreds, "C", "D", "M")
		n -= hundreds * 100
	}

	// Tens
	if (n > 10 && n <= 100) {
		tens := n / 10
		numerals += toNumerals(tens, "X", "L", "C")
		n -= tens * 10
	}

	// Ones
	numerals += toNumerals(n, "I", "V", "X")
	return numerals
}

func toNumerals(n uint, one string, five string, ten string) string {
	switch n {
	case 0 : return ""
	case 1 : return one
	case 2 : return one + one
	case 3 : return one + one + one
	case 4 : return one + five
	case 5 : return five
	case 6 : return five + one
	case 7 : return five + one + one
	case 8 : return five + one + one + one
	case 9 : return one + ten
	case 10: return ten
	default:
		panic(fmt.Sprintf("Unexpected value: %s", n))
	}
}

func Decode(s string) int {

	total, start, pos := 0, 0, 0

	// Thousands
	for ; pos < len(s) && s[pos] == 'M'; pos++ {}
	total += (pos - start) * 1000

	// Hundreds
	for start = pos; pos < len(s) && (s[pos] == 'C' ||  s[pos] == 'D' || s[pos] == 'M'); pos++ {}
	total += fromNumerals(s[start:pos], "C", "D", "M") * 100

	// Tens
	for start = pos; pos < len(s) && (s[pos] == 'X' ||  s[pos] == 'L' || s[pos] == 'C'); pos++ {}
	total += fromNumerals(s[start:pos], "X", "L", "C") * 10

	// Ones
	if (pos < len(s)) {
		total += fromNumerals(s[pos:], "I", "V", "X")
	}
	return total
}

func fromNumerals(s string, one string, five string, ten string) int {
	switch s {
	case "" : return 0
	case one: return 1
	case one + one: return 2
	case one + one + one: return 3
	case one + five: return 4
	case five: return 5
	case five + one: return 6
	case five + one + one: return 7
	case five + one + one + one: return 8
	case one + ten: return 9
	case ten: return 10
	default:
		panic(fmt.Sprintf("Unexpected value: %s", s))
	}
}