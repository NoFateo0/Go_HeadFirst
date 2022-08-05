package ordinals

import "testing"

type forTest struct {
	num  int
	want string
}

func TestOrdinal(t *testing.T) {
	testers := []forTest{
		{num: 1, want: "1st"},
		{num: 2, want: "2nd"},
		{num: 3, want: "3rd"},
		{num: 4, want: "4th"},
		{num: 11, want: "11th"},
		{num: 21, want: "21st"},
	}

	for _, test := range testers {
		got := Ordinal(test.num)
		if got != test.want {
			t.Errorf("Ordinal(3) = \"%s\", want \"%s\"", got, test.want)
		}
	}
}

func TestOne(t *testing.T) {
	got := Ordinal(1)
	if got != "1st" {
		t.Errorf("didn't match expected value")
	}
}

func TestTwo(t *testing.T) {
	got := Ordinal(2)
	want := "2nd"
	if got != want {
		t.Errorf("Ordinal(3) = \"%s\", want \"%s\"", got, want)
	}
}
