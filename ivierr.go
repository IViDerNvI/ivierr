package ivierr

type innerError struct {
	text string
}

func (e *innerError) Error() string {
	return e.text
}

func New(text string) error {
	return &innerError{text}
}
