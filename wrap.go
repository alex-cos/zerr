package zerr

func Wrap(err error, message string) error {
	return &ZError{
		severity: Error,
		code:     0,
		message:  message,
		err:      err,
	}
}

func WrapS(severity Severity, err error, message string) error {
	return &ZError{
		severity: severity,
		code:     0,
		message:  message,
		err:      err,
	}
}

func WrapC(code int64, err error, message string) error {
	return &ZError{
		severity: Error,
		code:     code,
		message:  message,
		err:      err,
	}
}

func WrapSC(severity Severity, code int64, err error, message string) error {
	return &ZError{
		severity: severity,
		code:     code,
		message:  message,
		err:      err,
	}
}
