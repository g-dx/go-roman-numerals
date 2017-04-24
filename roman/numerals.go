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
		numerals += base(hundreds, "C", "D", "M")
		n -= hundreds * 100
	}

	// Tens
	if (n > 10 && n <= 100) {
		tens := n / 10
		numerals += base(tens, "X", "L", "C")
		n -= tens * 10
	}

	// Ones
	numerals += base(n, "I", "V", "X")
	return numerals
}

func base(n uint, one string, five string, ten string) string {
	switch(n) {
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
		panic(fmt.Sprintf("Expected: value between 1 - 10, Actual: %s", n))
	}
}

func decode(n string) uint {
	return 0
}