package errors

type Kind int

const (
	GeneralError Kind = 1000 + iota
	ValidationError
)

const (
	DevelopmentError Kind = 8000 + iota
	UnsupportedQueryError
)

const (
	IllegalError Kind = 9000 + iota
)

func (e Kind) String() string {
	switch e {
	case ValidationError:
		return "ValidationError"
	case IllegalError:
		return "IllegalError"
	}
	return ""
}
