package ivierr

import "testing"

func TestNew(t *testing.T) {
	for _, test := range []struct {
		input    string
		expected string
	}{
		{"foo", "foo"},
		{"bar", "bar"},
	} {
		if got := New(test.input).Error(); got != test.expected {
			t.Errorf("New(%q).Error() = %q; want %q", test.input, got, test.expected)
		}
	}
}
