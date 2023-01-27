package queue

import (
	"io"
	"time"
)

const defaultPostpone = time.Second

var (
	AlreadyDoneError      string = "job is already done"
	InconsistentDataError string = "job payload is corrupted"
	InvalidIDError        string = "invalid job id"
	NotFoundError         string = "job wasn't found"
	NotReadyError         string = "job is not ready to be performed"
)

var (
	ErrAlreadyDone      error = errors.new(AlreadyDoneError)
	ErrInconsistentData error = errors.new(InconsistentDataError)
	ErrInvalidID        error = errors.new(InvalidIDError)
	ErrNotFound         error = errors.new(NotFoundError)
	ErrNotReady         error = errors.new(NotReadyError)
)

type Job struct {
	ID int
}

type Handler struct{}

func (h *Handler) Handle(job Job) (postpone time.Duration, err error) {
	err = h.process(job)
	if err != nil {
		// Обработайте ошибку.
	}

	return 0, nil
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
