package bdt

import "testing"

func TestRomanNumerals(t *testing.T) {
	tests := []struct {
		name  string
		input int
		want  string
	}{
		{"1", 1, "I"},
		{"2", 2, "II"},
		{"4", 4, "IV"},
		{"9", 9, "IX"},
		{"88", 88, "LXXXVIII"},
		{"642", 642, "DCXLII"},
		{"2299", 2299, "MMCCXCIX"},
		{"3999", 3999, "MMMCMXCIX"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assertRoman(t, test.input, test.want)
		})
	}
}

func assertRoman(t testing.TB, input int, want string) {
	t.Helper()

	got := ConvertToRoman(input)
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
