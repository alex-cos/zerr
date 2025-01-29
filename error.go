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
}

// Error the extended Error function.
func (z *ZError) Error() string {
	var buf strings.Builder

	buf.WriteString(z.severity.String())
	if z.code > 0 {
		buf.WriteString(fmt.Sprintf("[%d]", z.code))
	}
	if buf.Len() > 0 {
		buf.WriteString(" - ")
	}
	buf.WriteString(z.message)
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
