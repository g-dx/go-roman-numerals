package roman

import (
	"testing"
)

var data = []struct {
	val      uint
	numerals string
}{
	{1, "I"},
	{2, "II"},
	{3, "III"},
	{4, "IV"},
	{5, "V"},
	{6, "VI"},
	{7, "VII"},
	{8, "VIII"},
	{9, "IX"},
	{10, "X"},
	{11, "XI"},
	{49, "XLIX"},
	{99, "XCIX"},
	{100, "C"},
	{101, "CI"},
	{1929, "MCMXXIX"},
	{1998, "MCMXCVIII"},
	{1999, "MCMXCIX"},
	{2000, "MM"},
	{2001, "MMI"},
	{1333, "MCCCXXXIII"},
}

func TestEncode(t *testing.T) {

	for _, tt := range data {
		v := Encode(tt.val)
		if v != tt.numerals {
			t.Errorf("Encode(%d) => got %s, want %s", tt.val, v, tt.numerals)
		}
	}
}

func TestDecode(t *testing.T) {

	for _, tt := range data {
		v := Decode(tt.numerals)
		if v != int(tt.val) {
			t.Errorf("Decode(%s) => got %d, want %d", tt.numerals, v, tt.val)
		}
	}
}

func BenchmarkEncode(b *testing.B) {

	for i := 0; i < b.N; i++ {
		for _, test := range data {
			Encode(test.val)
		}
	}
}

func BenchmarkDecode(b *testing.B) {

	for i := 0; i < b.N; i++ {
		for _, test := range data {
			Decode(test.numerals)
		}
	}
}