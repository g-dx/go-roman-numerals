package roman

import "testing"

func TestEncode(t *testing.T) {

	var data = []struct {
		in  uint
		out string
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
		{1998, "MCMXCVIII"},
		{1999, "MCMXCIX"},
		{2000, "MM"},
		{2001, "MMI"},
		{1333, "MCCCXXXIII"},
	}

	for _, tt := range data {
		v := Encode(tt.in)
		if v != tt.out {
			t.Errorf("Encode(%d) => got %s, want %s", tt.in, v, tt.out)
		}
	}
}