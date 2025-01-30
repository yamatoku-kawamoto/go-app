package errors

type extendedError struct {
	err  error
	kind Kind
}

func Error(err error, kind Kind) error {
	return &extendedError{
		err:  err,
		kind: kind,
	}
}

func (e *extendedError) Error() string {
	return e.err.Error()
}

func (e *extendedError) Kind() Kind {
	return e.kind
}
