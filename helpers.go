package zerr

import "errors"

func IsSeverity(err error, severity Severity) bool {
	var zerr *ZError
	if errors.As(err, &zerr) {
		return zerr.severity == severity
	}
	return false
}

func GetCode(err error) (int64, bool) {
	var zerr *ZError
	if errors.As(err, &zerr) {
		return zerr.code, true
	}
	return 0, false
}

func GetMessage(err error) string {
	var zerr *ZError
	if err == nil {
		return ""
	}
	if errors.As(err, &zerr) {
		return zerr.message
	}
	return err.Error()
}

func Chain(err error) []error {
	if err == nil {
		return nil
	}
	chain := []error{err}
	for {
		u, ok := err.(interface{ Unwrap() error })
		if !ok {
			break
		}
		next := u.Unwrap()
		if next == nil {
			break
		}
		chain = append(chain, next)
		err = next
	}
	return chain
}
