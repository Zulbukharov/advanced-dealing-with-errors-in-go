package pipe

import (
	"fmt"
	"log"
	"reflect"
)

type UserError struct {
	Operation string
	User      string
}

func (u *UserError) Error() string {
	return fmt.Sprintf("user %s cannot do op %s", u.User, u.Operation)
}

type PipelineError struct {
	User        string
	Name        string
	FailedSteps []string
}

func (p *PipelineError) Error() string {
	return fmt.Sprintf("pipeline %q error", p.Name)
}

// Добавь метод As для типа *PipelineError.
func (p PipelineError) As(target any) bool {
	log.Println(reflect.TypeOf(target))
	switch tt := target.(type) {
	case **UserError:
		if *tt == nil {
			*tt = &UserError{}
		}
		(*tt).Operation = p.Name
		(*tt).User = p.User
		return true
	}
	return false
}
