package zerr_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/alex-cos/zerr"
	"github.com/stretchr/testify/assert"
)

func TestError(t *testing.T) {
	t.Parallel()

	t.Run("New", func(t *testing.T) {
		t.Parallel()
		err := zerr.New("response returned an empty body")
		zerror := new(zerr.ZError)
		assert.Error(t, err)
		assert.Equal(t, "Error - response returned an empty body", err.Error())
		assert.ErrorAs(t, err, &zerror)
		assert.Equal(t, zerr.Error, zerror.Severity())
		assert.Equal(t, int64(0), zerror.Code())
		assert.Equal(t, "response returned an empty body", zerror.Message())
	})

	t.Run("NewS", func(t *testing.T) {
		t.Parallel()
		err := zerr.NewS(zerr.Warning, "response returned an empty body")
		zerror := new(zerr.ZError)
		assert.Error(t, err)
		assert.Equal(t, "Warning - response returned an empty body", err.Error())
		assert.ErrorAs(t, err, &zerror)
		assert.Equal(t, zerr.Warning, zerror.Severity())
		assert.Equal(t, int64(0), zerror.Code())
		assert.Equal(t, "response returned an empty body", zerror.Message())
	})

	t.Run("NewC", func(t *testing.T) {
		t.Parallel()
		err := zerr.NewC(10012, "response returned an empty body")
		zerror := new(zerr.ZError)
		assert.Error(t, err)
		assert.Equal(t, "Error[10012] - response returned an empty body", err.Error())
		assert.ErrorAs(t, err, &zerror)
		assert.Equal(t, zerr.Error, zerror.Severity())
		assert.Equal(t, int64(10012), zerror.Code())
		assert.Equal(t, "response returned an empty body", zerror.Message())
	})

	t.Run("NewSC", func(t *testing.T) {
		t.Parallel()
		err := zerr.NewSC(zerr.Warning, 10012, "response returned an empty body")
		zerror := new(zerr.ZError)
		assert.Error(t, err)
		assert.Equal(t, "Warning[10012] - response returned an empty body", err.Error())
		assert.ErrorAs(t, err, &zerror)
		assert.Equal(t, zerr.Warning, zerror.Severity())
		assert.Equal(t, int64(10012), zerror.Code())
		assert.Equal(t, "response returned an empty body", zerror.Message())
	})

	t.Run("Errorf", func(t *testing.T) {
		t.Parallel()
		err := zerr.Errorf("failed to write file '%s'", "filename")
		zerror := new(zerr.ZError)
		assert.Error(t, err)
		assert.Equal(t, "Error - failed to write file 'filename'", err.Error())
		assert.ErrorAs(t, err, &zerror)
		assert.Equal(t, zerr.Error, zerror.Severity())
		assert.Equal(t, int64(0), zerror.Code())
		assert.Equal(t, "failed to write file 'filename'", zerror.Message())
	})

	t.Run("ErrorSf", func(t *testing.T) {
		t.Parallel()
		err := zerr.ErrorSf(zerr.Info, "failed to write file '%s'", "filename")
		zerror := new(zerr.ZError)
		assert.Error(t, err)
		assert.Equal(t, "Info - failed to write file 'filename'", err.Error())
		assert.ErrorAs(t, err, &zerror)
		assert.Equal(t, zerr.Info, zerror.Severity())
		assert.Equal(t, int64(0), zerror.Code())
		assert.Equal(t, "failed to write file 'filename'", zerror.Message())
	})

	t.Run("ErrorCf", func(t *testing.T) {
		t.Parallel()
		err := zerr.ErrorCf(10032, "failed to write file '%s'", "filename")
		zerror := new(zerr.ZError)
		assert.Error(t, err)
		assert.Equal(t, "Error[10032] - failed to write file 'filename'", err.Error())
		assert.ErrorAs(t, err, &zerror)
		assert.Equal(t, zerr.Error, zerror.Severity())
		assert.Equal(t, int64(10032), zerror.Code())
		assert.Equal(t, "failed to write file 'filename'", zerror.Message())
	})

	t.Run("ErrorSCf", func(t *testing.T) {
		t.Parallel()
		err := zerr.ErrorSCf(zerr.Info, 10032, "failed to write file '%s'", "filename")
		zerror := new(zerr.ZError)
		assert.Error(t, err)
		assert.Equal(t, "Info[10032] - failed to write file 'filename'", err.Error())
		assert.ErrorAs(t, err, &zerror)
		assert.Equal(t, zerr.Info, zerror.Severity())
		assert.Equal(t, int64(10032), zerror.Code())
		assert.Equal(t, "failed to write file 'filename'", zerror.Message())
	})
}

func TestWrap(t *testing.T) {
	t.Parallel()

	origErr := errors.New("original error")

	t.Run("Wrap", func(t *testing.T) {
		t.Parallel()
		err := zerr.Wrap(origErr, "wrapped message")
		zerror := new(zerr.ZError)
		assert.Error(t, err)
		assert.Equal(t, "Error - wrapped message: original error", err.Error())
		assert.ErrorAs(t, err, &zerror)
		assert.Equal(t, zerr.Error, zerror.Severity())
		assert.Equal(t, int64(0), zerror.Code())
		assert.Equal(t, "wrapped message", zerror.Message())
		assert.ErrorIs(t, err, origErr)
	})

	t.Run("WrapS", func(t *testing.T) {
		t.Parallel()
		err := zerr.WrapS(zerr.Critical, origErr, "wrapped message")
		zerror := new(zerr.ZError)
		assert.Error(t, err)
		assert.Equal(t, "Critical - wrapped message: original error", err.Error())
		assert.ErrorAs(t, err, &zerror)
		assert.Equal(t, zerr.Critical, zerror.Severity())
		assert.ErrorIs(t, err, origErr)
	})

	t.Run("WrapC", func(t *testing.T) {
		t.Parallel()
		err := zerr.WrapC(5001, origErr, "wrapped message")
		zerror := new(zerr.ZError)
		assert.Error(t, err)
		assert.Equal(t, "Error[5001] - wrapped message: original error", err.Error())
		assert.ErrorAs(t, err, &zerror)
		assert.Equal(t, int64(5001), zerror.Code())
		assert.ErrorIs(t, err, origErr)
	})

	t.Run("WrapSC", func(t *testing.T) {
		t.Parallel()
		err := zerr.WrapSC(zerr.Warning, 5002, origErr, "wrapped message")
		zerror := new(zerr.ZError)
		assert.Error(t, err)
		assert.Equal(t, "Warning[5002] - wrapped message: original error", err.Error())
		assert.ErrorAs(t, err, &zerror)
		assert.Equal(t, zerr.Warning, zerror.Severity())
		assert.Equal(t, int64(5002), zerror.Code())
		assert.ErrorIs(t, err, origErr)
	})

	t.Run("Unwrap returns nil when no wrapped error", func(t *testing.T) {
		t.Parallel()
		err := zerr.New("simple error")
		zerror := new(zerr.ZError)
		assert.ErrorAs(t, err, &zerror)
		assert.NoError(t, zerror.Unwrap())
	})

	t.Run("errors.As unwraps to underlying error", func(t *testing.T) {
		t.Parallel()
		underlying := zerr.NewC(999, "underlying zerr")
		err := zerr.Wrap(underlying, "outer wrapper")
		var outer *zerr.ZError
		assert.ErrorAs(t, err, &outer)
		assert.Equal(t, int64(0), outer.Code())
		assert.Equal(t, "outer wrapper", outer.Message())
		assert.ErrorIs(t, err, underlying)
	})
}

func TestHelpers(t *testing.T) {
	t.Parallel()

	t.Run("IsSeverity", func(t *testing.T) {
		t.Parallel()

		t.Run("matches severity", func(t *testing.T) {
			t.Parallel()
			err := zerr.NewS(zerr.Warning, "test")
			assert.True(t, zerr.IsSeverity(err, zerr.Warning))
		})

		t.Run("does not match different severity", func(t *testing.T) {
			t.Parallel()
			err := zerr.NewS(zerr.Warning, "test")
			assert.False(t, zerr.IsSeverity(err, zerr.Error))
		})

		t.Run("finds severity in wrapped chain", func(t *testing.T) {
			t.Parallel()
			inner := zerr.NewS(zerr.Critical, "inner")
			err := zerr.Wrap(inner, "outer")
			assert.False(t, zerr.IsSeverity(err, zerr.Critical))
			assert.True(t, zerr.IsSeverity(err, zerr.Error))
			assert.True(t, zerr.IsSeverity(inner, zerr.Critical))
		})

		t.Run("returns false for non-zerr error", func(t *testing.T) {
			t.Parallel()
			err := errors.New("standard error")
			assert.False(t, zerr.IsSeverity(err, zerr.Error))
		})

		t.Run("returns false for nil error", func(t *testing.T) {
			t.Parallel()
			assert.False(t, zerr.IsSeverity(nil, zerr.Error))
		})
	})

	t.Run("GetCode", func(t *testing.T) {
		t.Parallel()

		t.Run("returns code when present", func(t *testing.T) {
			t.Parallel()
			err := zerr.NewC(404, "not found")
			code, ok := zerr.GetCode(err)
			assert.True(t, ok)
			assert.Equal(t, int64(404), code)
		})

		t.Run("returns zero code with ok=true", func(t *testing.T) {
			t.Parallel()
			err := zerr.New("no code")
			code, ok := zerr.GetCode(err)
			assert.True(t, ok)
			assert.Equal(t, int64(0), code)
		})

		t.Run("finds code in wrapped chain", func(t *testing.T) {
			t.Parallel()
			inner := zerr.NewC(500, "inner")
			err := zerr.Wrap(inner, "outer")
			code, ok := zerr.GetCode(err)
			assert.True(t, ok)
			assert.Equal(t, int64(0), code)
		})

		t.Run("returns false for non-zerr error", func(t *testing.T) {
			t.Parallel()
			err := errors.New("standard error")
			_, ok := zerr.GetCode(err)
			assert.False(t, ok)
		})

		t.Run("returns false for nil error", func(t *testing.T) {
			t.Parallel()
			_, ok := zerr.GetCode(nil)
			assert.False(t, ok)
		})
	})

	t.Run("GetMessage", func(t *testing.T) {
		t.Parallel()

		t.Run("returns message when present", func(t *testing.T) {
			t.Parallel()
			err := zerr.New("hello world")
			msg := zerr.GetMessage(err)
			assert.Equal(t, "hello world", msg)
		})

		t.Run("returns outer message in wrapped chain", func(t *testing.T) {
			t.Parallel()
			inner := zerr.New("inner message")
			err := zerr.Wrap(inner, "outer message")
			msg := zerr.GetMessage(err)
			assert.Equal(t, "outer message", msg)
		})

		t.Run("returns standard error message for non-zerr error", func(t *testing.T) {
			t.Parallel()
			err := errors.New("standard error")
			msg := zerr.GetMessage(err)
			assert.Equal(t, "standard error", msg)
		})

		t.Run("returns empty string for nil error", func(t *testing.T) {
			t.Parallel()
			msg := zerr.GetMessage(nil)
			assert.Empty(t, msg)
		})
	})

	t.Run("Chain", func(t *testing.T) {
		t.Parallel()

		t.Run("returns nil for nil error", func(t *testing.T) {
			t.Parallel()
			assert.Nil(t, zerr.Chain(nil))
		})

		t.Run("returns single error when no wrapping", func(t *testing.T) {
			t.Parallel()
			err := zerr.New("single error")
			chain := zerr.Chain(err)
			assert.Len(t, chain, 1)
			assert.Equal(t, "Error - single error", chain[0].Error())
		})

		t.Run("returns full chain for wrapped zerr", func(t *testing.T) {
			t.Parallel()
			inner := zerr.NewC(500, "inner")
			outer := zerr.Wrap(inner, "outer")
			chain := zerr.Chain(outer)
			assert.Len(t, chain, 2)
			assert.Equal(t, "Error - outer: Error[500] - inner", outer.Error())
		})

		t.Run("handles multiple wrapping levels", func(t *testing.T) {
			t.Parallel()
			inner := zerr.NewC(500, "root cause")
			mid := zerr.WrapS(zerr.Warning, inner, "middle layer")
			outer := zerr.Wrap(mid, "top layer")
			chain := zerr.Chain(outer)
			assert.Len(t, chain, 3)
		})

		t.Run("works with standard errors", func(t *testing.T) {
			t.Parallel()
			stdErr := errors.New("standard error")
			wrapped := zerr.Wrap(stdErr, "wrapped")
			chain := zerr.Chain(wrapped)
			assert.Len(t, chain, 2)
			assert.Equal(t, "standard error", chain[1].Error())
		})

		t.Run("handles fmt.Errorf wrapping", func(t *testing.T) {
			t.Parallel()
			inner := zerr.NewC(404, "not found")
			outer := fmt.Errorf("context: %w", inner)
			chain := zerr.Chain(outer)
			assert.Len(t, chain, 2)
		})
	})
}
