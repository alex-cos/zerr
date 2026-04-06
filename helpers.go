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
