package error

import "github.com/pkg/errors"

var (
	ErrUnknown = errors.Errorf("unknown error %d", 10001)
)
