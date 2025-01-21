package ivierr

import (
	"reflect"
	"testing"
)

func TestKmpSearch(t *testing.T) {
	for _, test := range []struct {
		label    string
		text     string
		pattern  string
		expected []int
	}{
		{"basic test 1", "ababababca", "abababca", []int{2}},
		{"basic test 2", "ababababca", "abababac", []int{}},
		{"basic test 3", "ababababca", "ababababca", []int{0}},
	} {
		if got := kmpSearch(test.text, test.pattern); !reflect.DeepEqual(got, test.expected) {
			t.Errorf("[%s]: kmpSearch(%q, %q) = %#+v; want %#+v", test.label, test.text, test.pattern, got, test.expected)
		}
	}
}
