package queue

import (
	"io"
	"time"
)

const defaultPostpone = time.Second

var ()

type AlreadyDoneError struct{}

func (e *AlreadyDoneError) Error() string {
	return "job is already done"
}

type InconsistentDataError struct{}

func (e *InconsistentDataError) Error() string {
	return "job payload is corrupted"
}

type InvalidIDError struct{}

func (e *InvalidIDError) Error() string {
	return "invalid job id"
}

type NotFoundError struct{}

func (e *NotFoundError) Error() string {
	return "job wasn't found"
}

type NotReadyError struct{}

func (e *NotReadyError) Error() string {
	return "job is not ready to be performed"
}

var (
	ErrAlreadyDone      error = new(AlreadyDoneError)
	ErrInconsistentData error = new(InconsistentDataError)
	ErrInvalidID        error = new(InvalidIDError)
	ErrNotFound         error = new(NotFoundError)
	ErrNotReady         error = new(NotReadyError)
)

type Job struct {
	ID int
}

type Handler struct{}

func (h *Handler) Handle(job Job) (postpone time.Duration, err error) {
	err = h.process(job)
	switch err {
	case ErrAlreadyDone, ErrInconsistentData, ErrInvalidID, ErrNotFound:
		return 0, nil
	case ErrNotReady:
		return time.Second * 1, nil
	default:
		return 0, err
	}
}

func (h *Handler) process(job Job) error {
	switch job.ID {
	case 1:
		return ErrInconsistentData
	case 2:
		return ErrNotReady
	case 3:
		return ErrNotFound
	case 4:
		return ErrAlreadyDone
	case 5:
		return ErrInvalidID
	case 6:
		return io.EOF
	}
	return nil
}
