package anagrams

import (
	"slices"
	"strings"
	"testing"
)

func compareGroupedAnagrams(result, expected [][]string) bool {
	for _, slice := range result {
		slices.Sort(slice)
	}

	for _, slice := range expected {
		slices.Sort(slice)
	}

	cmp := func(a, b []string) int {
		if len(a) == 0 {
			return -1
		}

		if len(b) == 0 {
			return 1
		}

		return strings.Compare(a[0], b[0])
	}

	slices.SortFunc(result, cmp)
	slices.SortFunc(expected, cmp)

	for i := 0; i < len(result); i++ {
		if slices.Compare(result[i], expected[i]) != 0 {
			return false
		}
	}

	return true
}

func TestCreateToken(t *testing.T) {
	tests := []struct {
		input    string
		expected token
	}{
		{
			"",
			token{"", ""},
		},
		{
			"a",
			token{"a", "a"},
		},
		{
			"za",
			token{"az", "za"},
		},
		{
			"aab",
			token{"aab", "aab"},
		},
		{
			"aba",
			token{"aab", "aba"},
		},
		{
			"baa",
			token{"aab", "baa"},
		},
		{
			"eat",
			token{"aet", "eat"},
		},
	}

	for i, test := range tests {
		result := createToken(test.input)
		if result.key != test.expected.key {
			t.Errorf("test[%d] key wrong. expected=%q got=%q",
				i, test.expected.key, result.key)
		}
		if result.value != test.expected.value {
			t.Errorf("test[%d] value wrong. expected=%q got=%q",
				i, test.expected.value, result.value)
		}
	}
}

func TestGroup(t *testing.T) {
	tests := []struct {
		input    []string
		expected [][]string
	}{
		{
			[]string{""},
			[][]string{{""}},
		},
		{
			[]string{"a"},
			[][]string{{"a"}},
		},
		{
			[]string{"a", "a"},
			[][]string{{"a", "a"}},
		},
		{
			[]string{"ab", "ba"},
			[][]string{{"ab", "ba"}},
		},
		{
			[]string{"a", "b", "c", "d"},
			[][]string{{"a"}, {"b"}, {"c"}, {"d"}},
		},
		{
			[]string{"ab", "b", "c", "d", "ab"},
			[][]string{{"ab", "ab"}, {"b"}, {"c"}, {"d"}},
		},
		{
			[]string{"eat", "tea", "tan", "ate", "nat", "bat"},
			[][]string{{"bat"}, {"nat", "tan"}, {"ate", "eat", "tea"}},
		},
	}

	for i, test := range tests {
		result := Group(test.input)
		if !compareGroupedAnagrams(result, test.expected) {
			t.Errorf("test[%d] result wrong. expected=%v got=%v",
				i, test.expected, result)
		}
	}
}
