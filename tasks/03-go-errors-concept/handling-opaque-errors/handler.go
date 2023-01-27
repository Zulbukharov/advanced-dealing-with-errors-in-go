package queue

import (
	"io"
	"time"
)

const defaultPostpone = time.Second

type Job struct {
	ID int
}

type Handler struct{}

func (h *Handler) Handle(job Job) (postpone time.Duration, err error) {
	err = h.process(job)
	if err != nil {
		if isTemporary(err) {
			return time.Second * 1, nil
		} else if shouldBeSkipped(err) {
			return 0, nil
		}
	}
	return 0, err
}

func isTemporary(err error) bool {
	type temporary interface {
		Temporary() bool
	}
	if t, ok := err.(temporary); ok {
		return t.Temporary()
	}
	return false
}

func shouldBeSkipped(err error) bool {
	// Реализуй меня.
	type skip interface {
		Skip() bool
	}
	if s, ok := err.(skip); ok {
		return s.Skip()
	}
	return false
}

func (h *Handler) process(job Job) error {
	switch job.ID {
	case 1:
		return &InconsistentDataError{}
	case 2:
		return &NotReadyError{}
	case 3:
		return &NotFoundError{}
	case 4:
		return &AlreadyDoneError{}
	case 5:
		return &InvalidIDError{}
	case 6:
		return io.EOF
	}
	return nil
}
