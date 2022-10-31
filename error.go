package sitemapcrawl

func newUnexpectedTypeError(name string) UnexpectedTypeError {
	return UnexpectedTypeError{
		Name: name,
	}
}

type UnexpectedTypeError struct {
	Name string
}

func (u UnexpectedTypeError) Error() string {
	return "unexpected type: " + u.Name
}

var _ error = (*UnexpectedTypeError)(nil)
