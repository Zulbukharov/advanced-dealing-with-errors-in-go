package errors

var (
	ErrAlreadyDone      error = &AlreadyDoneError{Err{"job is already done"}}
	ErrInconsistentData error = &InconsistentDataError{Err{"job payload is corrupted"}}
	ErrInvalidID        error = &InvalidIDError{Err{"invalid job id"}}
	ErrNotReady         error = &NotReadyError{Err{"job is not ready to be performed"}}
	ErrNotFound         error = &NotFoundError{Err{"job wasn't found"}}
)

// Реализуй тип Err и типы для ошибок выше, используя его.

type AlreadyDoneError struct {
	Err
}

func (e *AlreadyDoneError) Error() string {
	return e.Err.Error()
}

type InconsistentDataError struct {
	Err
}

func (e *InconsistentDataError) Error() string {
	return e.Err.Error()
}

type InvalidIDError struct {
	Err
}

func (e *InvalidIDError) Error() string {
	return e.Err.Error()
}

type NotReadyError struct {
	Err
}

func (e *NotReadyError) Error() string {
	return e.Err.Error()
}

type NotFoundError struct {
	Err
}

func (e *NotFoundError) Error() string {
	return e.Err.Error()
}

type Err struct {
	msg string
}

func (e *Err) Error() string {
	return e.msg
}
