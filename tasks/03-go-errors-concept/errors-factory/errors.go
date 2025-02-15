package errors

// NewError возвращает новое значение-ошибку, текст которой является msg.
// Две ошибки с одинаковым текстом, созданные через NewError, не равны между собой:
//
//	NewError("end of file") != NewError("end of file")
func NewError(msg string) error {
	// Реализуй меня.
	return &errorStruct{msg}
}

type errorStruct struct {
	s string
}

func (e *errorStruct) Error() string {
	return e.s
}
