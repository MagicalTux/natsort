package natsort

import (
	"reflect"
	"sort"
	"strings"
	"testing"

	"github.com/maruel/natural"
)

var testList = []string{
	"1000X Radonius Maximus",
	"000050X Radonius",
	"10X Radonius",
	"200X Radonius",
	"20X Radonius",
	"20X Radonius Prime",
	"30X Radonius",
	"40X Radonius",
	"Allegia 50 Clasteron",
	"Allegia 500 Clasteron",
	"Allegia 50B Clasteron",
	"Allegia 51 Clasteron",
	"Allegia 6R Clasteron",
	"Alpha 100",
	"Alpha 2",
	"Alpha 200",
	"Alpha 2A",
	"Alpha 2A-8000",
	"Alpha 2A-900",
	"Callisto Morphamax",
	"Callisto Morphamax 500",
	"Callisto Morphamax 5000",
	"Callisto Morphamax 600",
	"Callisto Morphamax 6000 SE",
	"Callisto Morphamax 6000 SE2",
	"Callisto Morphamax 700",
	"Callisto Morphamax 7000",
	"Xiph Xlater 10000",
	"Xiph Xlater 2000",
	"Xiph Xlater 300",
	"Xiph Xlater 40",
	"Xiph Xlater 5",
	"Xiph Xlater 50",
	"Xiph Xlater 500",
	"Xiph Xlater 5000",
	"Xiph Xlater 58",
}

func Test_Sort1(t *testing.T) {
	testListSortedOK := []string{
		"10X Radonius",
		"20X Radonius",
		"20X Radonius Prime",
		"30X Radonius",
		"40X Radonius",
		"000050X Radonius",
		"200X Radonius",
		"1000X Radonius Maximus",
		"Allegia 6R Clasteron",
		"Allegia 50 Clasteron",
		"Allegia 50B Clasteron",
		"Allegia 51 Clasteron",
		"Allegia 500 Clasteron",
		"Alpha 2",
		"Alpha 2A",
		"Alpha 2A-900",
		"Alpha 2A-8000",
		"Alpha 100",
		"Alpha 200",
		"Callisto Morphamax",
		"Callisto Morphamax 500",
		"Callisto Morphamax 600",
		"Callisto Morphamax 700",
		"Callisto Morphamax 5000",
		"Callisto Morphamax 6000 SE",
		"Callisto Morphamax 6000 SE2",
		"Callisto Morphamax 7000",
		"Xiph Xlater 5",
		"Xiph Xlater 40",
		"Xiph Xlater 50",
		"Xiph Xlater 58",
		"Xiph Xlater 300",
		"Xiph Xlater 500",
		"Xiph Xlater 2000",
		"Xiph Xlater 5000",
		"Xiph Xlater 10000",
	}
	testListSorted := make([]string, len(testList))
	copy(testListSorted, testList)
	Sort(testListSorted)

	if !reflect.DeepEqual(testListSortedOK, testListSorted) {
		t.Fatalf(`ERROR: sorted list different from expected results:
	Expected results:
%v

	Got:
%v`, strings.Join(testListSortedOK, "\n"), strings.Join(testListSorted, "\n"))
	}
}

func Test_Sort2(t *testing.T) {
	testList := []string{
		"z1.doc",
		"z10.doc",
		"z100.doc",
		"z101.doc",
		"z102.doc",
		"z11.doc",
		"z12.doc",
		"z13.doc",
		"z14.doc",
		"z15.doc",
		"z16.doc",
		"z17.doc",
		"z18.doc",
		"z19.doc",
		"z2.doc",
		"z20.doc",
		"z3.doc",
		"z4.doc",
		"z5.doc",
		"z6.doc",
		"z7.doc",
		"z8.doc",
		"z9.doc",
	}

	testListSortedOK := []string{
		"z1.doc",
		"z2.doc",
		"z3.doc",
		"z4.doc",
		"z5.doc",
		"z6.doc",
		"z7.doc",
		"z8.doc",
		"z9.doc",
		"z10.doc",
		"z11.doc",
		"z12.doc",
		"z13.doc",
		"z14.doc",
		"z15.doc",
		"z16.doc",
		"z17.doc",
		"z18.doc",
		"z19.doc",
		"z20.doc",
		"z100.doc",
		"z101.doc",
		"z102.doc",
	}

	testListSorted := testList[:]
	Sort(testListSorted)

	if !reflect.DeepEqual(testListSortedOK, testListSorted) {
		t.Fatalf(`ERROR: sorted list different from expected results:
	Expected results:
%v

	Got:
%v`, strings.Join(testListSortedOK, "\n"), strings.Join(testListSorted, "\n"))
	}
}

func BenchmarkSort1(b *testing.B) {
	for n := 0; n < b.N; n++ {
		b.StopTimer()
		list := make([]string, len(testList))
		copy(list, testList)
		b.StartTimer()
		Sort(list)
	}
}

func BenchmarkSortMaruelNatural(b *testing.B) {
	for n := 0; n < b.N; n++ {
		b.StopTimer()
		list := make([]string, len(testList))
		copy(list, testList)
		b.StartTimer()
		sort.Sort(natural.StringSlice(list))
	}
}

func Test_Compare(t *testing.T) {
	tests := []struct {
		a, b     string
		expected bool
		desc     string
	}{
		// Basic comparisons
		{"a", "b", true, "a < b"},
		{"b", "a", false, "b > a"},
		{"a", "a", false, "a == a"},
		{"", "", false, "empty == empty"},
		{"", "a", true, "empty < a"},
		{"a", "", false, "a > empty"},

		// Numeric comparisons
		{"1", "2", true, "1 < 2"},
		{"2", "1", false, "2 > 1"},
		{"9", "10", true, "9 < 10"},
		{"10", "9", false, "10 > 9"},
		{"1", "10", true, "1 < 10"},
		{"10", "1", false, "10 > 1"},

		// Leading zeros
		{"0", "1", true, "0 < 1"},
		{"00", "1", true, "00 < 1"},
		{"01", "1", false, "01 == 1"},
		{"001", "1", false, "001 == 1"},
		{"09", "10", true, "09 < 10"},
		{"009", "10", true, "009 < 10"},

		// Leading zeros followed by non-digit (the bug we fixed)
		{"0a", "1", true, "0a < 1 (0 < 1, then compare suffix)"},
		{"00a", "1", true, "00a < 1 (0 < 1, then compare suffix)"},
		{"0a", "0b", true, "0a < 0b (equal numbers, a < b)"},
		{"0z", "1a", true, "0z < 1a (0 < 1)"},

		// Mixed alphanumeric
		{"a1", "a2", true, "a1 < a2"},
		{"a2", "a10", true, "a2 < a10"},
		{"a10", "a2", false, "a10 > a2"},
		{"z1.doc", "z2.doc", true, "z1.doc < z2.doc"},
		{"z2.doc", "z10.doc", true, "z2.doc < z10.doc"},
		{"z10.doc", "z2.doc", false, "z10.doc > z2.doc"},

		// Equal strings
		{"abc", "abc", false, "abc == abc"},
		{"123", "123", false, "123 == 123"},
		{"a1b2", "a1b2", false, "a1b2 == a1b2"},
	}

	for _, tt := range tests {
		result := Compare(tt.a, tt.b)
		if result != tt.expected {
			t.Errorf("Compare(%q, %q) = %v, expected %v (%s)", tt.a, tt.b, result, tt.expected, tt.desc)
		}
	}
}

func Test_SortEdgeCases(t *testing.T) {
	// Test sorting with leading zeros followed by non-digits
	input := []string{"1a", "0b", "00c", "2a", "0a"}
	expected := []string{"0a", "0b", "00c", "1a", "2a"}

	result := make([]string, len(input))
	copy(result, input)
	Sort(result)

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("Sort edge case failed:\nExpected: %v\nGot: %v", expected, result)
	}
}
