package pipe

import (
	"errors"
	"fmt"
)

type PipelineError struct {
	User        string
	Name        string
	FailedSteps []string
}

func (p *PipelineError) Error() string {
	return fmt.Sprintf("pipeline %q error", p.Name)
}

func IsPipelineError(err error, user, pipelineName string) bool {
	p := &PipelineError{}
	if ok := errors.As(err, &p); ok && p.User == user && p.Name == pipelineName {
		return true
	}
	return false
}
