package errs

type Unwrapper interface {
	Unwrap() error
}

func Unwrap(err error) error {
	for {
		nErr, ok := err.(Unwrapper)
		if !ok {
			return err
		}
		err = nErr.Unwrap()
	}
}
