package zerr_test

import (
	"errors"
	"testing"

	"github.com/alex-cos/zerr"
	"github.com/stretchr/testify/assert"
)

func TestError(t *testing.T) {
	t.Parallel()

	zerror := new(zerr.ZError)

	err := zerr.New("response returned an empty body")
	assert.Error(t, err)
	assert.Equal(t, "Error - response returned an empty body", err.Error())
	assert.ErrorAs(t, err, &zerror)
	assert.Equal(t, zerr.Error, zerror.Severity())
	assert.Equal(t, int64(0), zerror.Code())
	assert.Equal(t, "response returned an empty body", zerror.Message())

	err = zerr.NewS(zerr.Warning, "response returned an empty body")
	assert.Error(t, err)
	assert.Equal(t, "Warning - response returned an empty body", err.Error())
	assert.ErrorAs(t, err, &zerror)
	assert.Equal(t, zerr.Warning, zerror.Severity())
	assert.Equal(t, int64(0), zerror.Code())
	assert.Equal(t, "response returned an empty body", zerror.Message())

	err = zerr.NewC(10012, "response returned an empty body")
	assert.Error(t, err)
	assert.Equal(t, "Error[10012] - response returned an empty body", err.Error())
	assert.ErrorAs(t, err, &zerror)
	assert.Equal(t, zerr.Error, zerror.Severity())
	assert.Equal(t, int64(10012), zerror.Code())
	assert.Equal(t, "response returned an empty body", zerror.Message())

	err = zerr.NewSC(zerr.Warning, 10012, "response returned an empty body")
	assert.Error(t, err)
	assert.Equal(t, "Warning[10012] - response returned an empty body", err.Error())
	assert.ErrorAs(t, err, &zerror)
	assert.Equal(t, zerr.Warning, zerror.Severity())
	assert.Equal(t, int64(10012), zerror.Code())
	assert.Equal(t, "response returned an empty body", zerror.Message())

	err = zerr.Errorf("failed to write file '%s'", "filename")
	assert.Error(t, err)
	assert.Equal(t, "Error - failed to write file 'filename'", err.Error())
	assert.ErrorAs(t, err, &zerror)
	assert.Equal(t, zerr.Error, zerror.Severity())
	assert.Equal(t, int64(0), zerror.Code())
	assert.Equal(t, "failed to write file 'filename'", zerror.Message())

	err = zerr.ErrorSf(zerr.Info, "failed to write file '%s'", "filename")
	assert.Error(t, err)
	assert.Equal(t, "Info - failed to write file 'filename'", err.Error())
	assert.ErrorAs(t, err, &zerror)
	assert.Equal(t, zerr.Info, zerror.Severity())
	assert.Equal(t, int64(0), zerror.Code())
	assert.Equal(t, "failed to write file 'filename'", zerror.Message())

	err = zerr.ErrorCf(10032, "failed to write file '%s'", "filename")
	assert.Error(t, err)
	assert.Equal(t, "Error[10032] - failed to write file 'filename'", err.Error())
	assert.ErrorAs(t, err, &zerror)
	assert.Equal(t, zerr.Error, zerror.Severity())
	assert.Equal(t, int64(10032), zerror.Code())
	assert.Equal(t, "failed to write file 'filename'", zerror.Message())

	err = zerr.ErrorSCf(zerr.Info, 10032, "failed to write file '%s'", "filename")
	assert.Error(t, err)
	assert.Equal(t, "Info[10032] - failed to write file 'filename'", err.Error())
	assert.ErrorAs(t, err, &zerror)
	assert.Equal(t, zerr.Info, zerror.Severity())
	assert.Equal(t, int64(10032), zerror.Code())
	assert.Equal(t, "failed to write file 'filename'", zerror.Message())
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
