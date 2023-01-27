package pipe

import (
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

// Добавь метод Is для типа *PipelineError.
func (p *PipelineError) Is(target error) bool {
	if t, ok := target.(*PipelineError); ok &&
		t.Name == p.Name &&
		p.User == t.User {
		return true
	}
	return false
}
