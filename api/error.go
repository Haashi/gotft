package api

type error struct {
	code    int
	message string
}

func (err error) Error() string {
	return err.message
}

const (
	ErrorNetwork      = 1
	ErrorNotFound     = 2
	ErrorUnauthorized = 3
	ErrorForbidden    = 4
	ErrorBadRequest   = 5
	ErrorDecode       = 6
)
