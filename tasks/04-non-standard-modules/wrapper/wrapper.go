package errors

// Wrapper требует от типа быть ошибкой, поддерживающей API
// как стандартной библиотеки, так и github/pkg/errors.
type Wrapper interface { // Добавь интерфейсу методов.
	Cause() error
	Unwrap() error
	Error() string
}
