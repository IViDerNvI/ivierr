package ivierr

import (
	"fmt"
	"testing"
)

func TestWrap(t *testing.T) {
	for _, tc := range []struct {
		name     string
		errs     []error
		expected string
	}{
		{"basic test", []error{New("foo"), New("bar")}, fmt.Sprintf("%s\n%s", "foo", "bar")},
		{"nil error", []error{New("foo"), nil, New("bar")}, fmt.Sprintf("%s\n%s", "foo", "bar")},
	} {
		if got := Wrap(tc.errs...); got.Error() != tc.expected {
			t.Errorf("Wrap(%q) = %q; want %q", tc.errs, got, tc.expected)
		}
	}
}

func TestUnwrap(t *testing.T) {
	for _, tc := range []struct {
		name     string
		errs     []error
		expected string
	}{
		{"basic test", []error{New("foo"), New("bar")}, "foo"},
		{"nil error", []error{New("foo"), nil, New("bar")}, "foo"},
	} {
		if got := Unwrap(Wrap(tc.errs...)); got.Error() != tc.expected {
			t.Errorf("Unwrap(%q) = %q; want %q", tc.errs, got, tc.expected)
		}
	}
}

func BenchmarkWrap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Wrap(New("foo"), New("bar"))
	}
}

func BenchmarkUnwrap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Unwrap(Wrap(New("foo"), New("bar")))
	}
}
