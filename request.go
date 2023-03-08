package rest

type Request interface {
	Execute() ([]byte, error)
}
