package api

type Error struct {
	Code    int
	Message string
}

func (err Error) Error() string {
	return err.Message
}

const (
	ErrorNetwork      = 1
	ErrorNotFound     = 2
	ErrorUnauthorized = 3
	ErrorForbidden    = 4
	ErrorBadRequest   = 5
	ErrorDecode       = 6
)
