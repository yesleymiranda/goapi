package decoder_test

import (
	"goapi/pkg/web/decoder"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func Test_DecodeIDInt64_IsInvalid_Happy(t *testing.T) {
	req := httptest.NewRequest("GET", "/test/aaaa", nil)
	vars := map[string]string{"id": "aaaa"}
	req = mux.SetURLVars(req, vars)

	id, err := decoder.DecodeIDInt64(req)

	assert.Zero(t, id)
	assert.NotNil(t, err)
	assert.Equal(t, "id is invalid", err.Error())
}

func Test_DecodeIDInt64_IsRequired_Happy(t *testing.T) {
	req := httptest.NewRequest("GET", "/test/1111", nil)

	id, err := decoder.DecodeIDInt64(req)

	assert.Zero(t, id)
	assert.NotNil(t, err)
	assert.Equal(t, "id is required", err.Error())
}

func Test_DecodeIDInt64_Happy(t *testing.T) {
	req := httptest.NewRequest("GET", "/test/1111", nil)
	vars := map[string]string{"id": "1111"}
	req = mux.SetURLVars(req, vars)

	id, err := decoder.DecodeIDInt64(req)

	assert.Nil(t, err)
	assert.NotNil(t, id)
	assert.Equal(t, int64(1111), id)
}
