package strings_test

import (
	"testing"

	"github.com/cdvelop/strings"
)

func TestStringSplit(t *testing.T) {
	testCases := []struct {
		data      string
		separator string
		expected  []string
	}{
		{"texto1,texto2", ",", []string{"texto1", "texto2"}},
		{"apple,banana,cherry", ",", []string{"apple", "banana", "cherry"}},
		{"one.two.three.four", ".", []string{"one", "two", "three", "four"}},
		{"hello world", " ", []string{"hello", "world"}},
		{"hello. world", ".", []string{"hello", " world"}},
		{"h.", ".", []string{"h."}},
	}

	for _, tc := range testCases {
		result := strings.Split(tc.data, tc.separator)

		if !areStringSlicesEqual(result, tc.expected) {
			t.Errorf("StringSplit(%s, %s) = %v; expected %v", tc.data, tc.separator, result, tc.expected)
		}
	}
}

func areStringSlicesEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
