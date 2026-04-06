package zerr

import (
	"fmt"
	"strings"
)

// ZError the enhanced error structure.
type ZError struct {
	severity Severity
	code     int64
	message  string
	err      error
}

// Error the extended Error function.
func (z *ZError) Error() string {
	var buf strings.Builder

	buf.WriteString(z.severity.String())
	if z.code > 0 {
		buf.WriteString(fmt.Sprintf("[%d]", z.code))
	}
	buf.WriteString(" - ")
	buf.WriteString(z.message)
	if z.err != nil {
		buf.WriteString(": ")
		buf.WriteString(z.err.Error())
	}
	return buf.String()
}

func (z *ZError) Severity() Severity {
	return z.severity
}

func (z *ZError) Code() int64 {
	return z.code
}

func (z *ZError) Message() string {
	return z.message
}

func (z *ZError) Unwrap() error {
	return z.err
}

func New(message string) error {
	return &ZError{
		severity: Error,
		code:     0,
		message:  message,
	}
}

func NewS(severity Severity, message string) error {
	return &ZError{
		severity: severity,
		code:     0,
		message:  message,
	}
}

func NewC(code int64, message string) error {
	return &ZError{
		severity: Error,
		code:     code,
		message:  message,
	}
}

func NewSC(severity Severity, code int64, message string) error {
	return &ZError{
		severity: severity,
		code:     code,
		message:  message,
	}
}

func Errorf(format string, v ...interface{}) error {
	message := fmt.Sprintf(format, v...)
	return &ZError{
		severity: Error,
		code:     0,
		message:  message,
	}
}

func ErrorSf(severity Severity, format string, v ...interface{}) error {
	message := fmt.Sprintf(format, v...)
	return &ZError{
		severity: severity,
		code:     0,
		message:  message,
	}
}

func ErrorCf(code int64, format string, v ...interface{}) error {
	message := fmt.Sprintf(format, v...)
	return &ZError{
		severity: Error,
		code:     code,
		message:  message,
	}
}

func ErrorSCf(severity Severity, code int64, format string, v ...interface{}) error {
	message := fmt.Sprintf(format, v...)
	return &ZError{
		severity: severity,
		code:     code,
		message:  message,
	}
}

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
