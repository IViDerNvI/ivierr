package ivierr

import "testing"

func TestIs(t *testing.T) {
	for _, test := range []struct {
		test_type string
		label     int
		err       error
		target    error
		expected  bool
	}{
		{"basic test", 1, Wrap(New("ErrA"), New("ErrB")), New("ErrA"), true},
		{"basic test", 2, Wrap(New("ErrA")), New("ErrA"), true},
	} {
		if got := Is(test.err, test.target); got != test.expected {
			t.Errorf("[%s #%d]: Is(%q, %q) = %t; want %t", test.test_type, test.label, test.err, test.target, got, test.expected)
		}
	}
}

func BenchmarkIs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Is(Wrap(New("ErrA"), New("ErrB")), New("ErrA"))
	}
}
