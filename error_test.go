package zerr_test

import (
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
