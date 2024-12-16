package error

import "github.com/pkg/errors"

var (
	ErrUnknown         = errors.Errorf("Unknown error %d", 10001)
	ErrTimeout         = errors.Errorf("Timeout %d", 10002)
	ErrInvalidPlayerID = errors.Errorf("Invalid player id %d", 10003)
	ErrInvalidParams   = errors.Errorf("Invalid request params %d", 10004)
	ErrInvalidHeader   = errors.Errorf("Invalid header %d", 10005)
	ErrInvalidUser     = errors.Errorf("Invalid user %d", 10006)
	ErrInvalidScore    = errors.Errorf("Invalid score %d", 10006)
)
