package errs

import (
	"errors"
	"fmt"
	"time"
)

type WithTimeError struct {
	creationDate time.Time
	err          error
}

func NewWithTimeError(err error) error {
	return newWithTimeError(err, time.Now)
}

func (e *WithTimeError) Error() string {
	return fmt.Sprintf("%v, occurred at: %s", e.err, e.creationDate)
}

func (e *WithTimeError) Is(target error) bool {
	return errors.Is(e.err, target)
}

func (e *WithTimeError) Time() time.Time {
	return e.creationDate
}

func (e *WithTimeError) As(target any) bool {
	return errors.As(e.err, target)
}

func newWithTimeError(err error, timeFunc func() time.Time) error {
	return &WithTimeError{
		creationDate: timeFunc(),
		err:          err,
	}
}
