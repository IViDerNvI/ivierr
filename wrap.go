package ivierr

import "fmt"

type errStack struct {
	err []error
}

type errUnwrapable interface {
	Unwrap() error
}

func (e *errStack) Error() string {
	var s string

	for i, err := range e.err {
		if err == nil {
			continue
		}

		if i == 0 {
			s += err.Error()
			continue
		}

		s += fmt.Sprintf("\n%s", err.Error())
	}

	return s
}

func (e *errStack) wrap(err ...error) {
	e.err = append(e.err, err...)
}

func Wrap(err ...error) error {
	var n int
	for _, e := range err {
		if e != nil {
			n++
		}
	}

	if n == 0 {
		return nil
	}

	e := &errStack{err: make([]error, 0, n)}

	e.wrap(err...)

	return e
}

func Unwrap(err error) error {
	e, ok := err.(errUnwrapable)
	if !ok {
		return nil
	}
	return e.Unwrap()
}

func (e *errStack) Unwrap() error {
	if len(e.err) == 0 {
		return nil
	}

	e.err = e.err[:len(e.err)-1]

	return e
}
