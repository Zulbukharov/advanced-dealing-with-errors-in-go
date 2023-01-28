package requests

import (
	// Доступные пакеты, _ для сохранения импортов.
	"errors"
	_ "errors"
	"fmt"
	_ "fmt"
	"regexp"
	_ "regexp"
	"strings"
	_ "strings"
)

const maxPageSize = 100

// Реализуй нас.
var (
	errIsNotRegexp     = errors.New("exp is not regexp")
	errInvalidPage     = errors.New("invalid page")
	errInvalidPageSize = errors.New("invalid page size")
)

// Реализуй мои методы.
type ValidationErrors []error

func (e ValidationErrors) Error() string {
	msg := strings.Builder{}
	for _, v := range e {
		msg.WriteString(v.Error())
		msg.WriteByte('\t')
	}
	return msg.String()
}

func (e ValidationErrors) Is(target error) bool {
	for i := range e {
		if errors.Is(e[i], target) {
			return true
		}
	}
	return false
}

type SearchRequest struct {
	Exp      string
	Page     int
	PageSize int
}

func ValidatePageSize(size int) error {
	if size > maxPageSize {
		return fmt.Errorf("%w: %d > %d", errInvalidPageSize, size, maxPageSize)
	}
	if size < 1 {
		return fmt.Errorf("%w: %d < 1", errInvalidPageSize, size)
	}
	return nil
}

func ValidatePage(page int) error {
	if page < 0 {
		return fmt.Errorf("%w: %d", errInvalidPage, page)
	}
	return nil
}

func ValidateExp(exp string) error {
	if _, err := regexp.Compile(exp); err != nil {
		return fmt.Errorf("%w: %v", errIsNotRegexp, err)
	}
	return nil
}

func (r SearchRequest) Validate() error {
	errs := make(ValidationErrors, 0)
	if err := ValidateExp(r.Exp); err != nil {
		errs = append(errs, err)
	}

	if err := ValidatePage(r.Page); err != nil {
		errs = append(errs, err)
	}

	if err := ValidatePageSize(r.PageSize); err != nil {
		errs = append(errs, err)
	}
	if len(errs) > 0 {
		return errs
	}
	return nil
}
