package hootsuite

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetHTTPStatusEndpoint(t *testing.T) {
	t.Run("HTTP_STATUS_ENDPOINT_ERROR_EMPTY_NAME", func(t *testing.T) {
		name := ""
		slug := "test-service"
		url := "http://dev.bhinneka.com/test"
		_, err := GetHTTPStatusEndpoint(name, slug, url, httpURLType, false)
		assert.Error(t, err)
	})
	t.Run("HTTP_STATUS_ENDPOINT_ERROR_INVALID_SLUG", func(t *testing.T) {
		name := "Test Service"
		slug := "test service"
		url := "http://dev.bhinneka.com/test"
		_, err := GetHTTPStatusEndpoint(name, slug, url, httpURLType, false)
		assert.Error(t, err)
	})
	t.Run("HTTP_STATUS_ENDPOINT_ERROR_EMPTY_SLUG", func(t *testing.T) {
		name := "Test Service"
		slug := ""
		url := "http://dev.bhinneka.com/test"
		_, err := GetHTTPStatusEndpoint(name, slug, url, httpURLType, false)
		assert.Error(t, err)
	})
	t.Run("HTTP_STATUS_ENDPOINT_ERROR_EMPTY_URL", func(t *testing.T) {
		name := "Test Service"
		slug := "test-service"
		url := ""
		_, err := GetHTTPStatusEndpoint(name, slug, url, httpURLType, false)
		assert.Error(t, err)
	})
	t.Run("HTTP_STATUS_ENDPOINT_ERROR_INVALID_URL_TYPE", func(t *testing.T) {
		name := "Test Service"
		slug := "test-service"
		url := "http://dev.bhinneka.com/test"
		_, err := GetHTTPStatusEndpoint(name, slug, url, "", false)
		assert.Error(t, err)
	})
	t.Run("HTTP_STATUS_ENDPOINT_SUCCESS_NO_TRAVERSE", func(t *testing.T) {
		name := "Test Service"
		slug := "test-service"
		url := "http://dev.bhinneka.com/test"
		statusEndpoint, err := GetHTTPStatusEndpoint(name, slug, url, httpURLType, false)
		assert.NoError(t, err)
		assert.Equal(t, name, statusEndpoint.Name)
		assert.Equal(t, slug, statusEndpoint.Slug)
		assert.Equal(t, httpURLType, statusEndpoint.Type)
		assert.Equal(t, false, statusEndpoint.IsTraversable)
	})
	t.Run("HTTP_STATUS_ENDPOINT_SUCCESS_TRAVERSE", func(t *testing.T) {
		name := "Test Service"
		slug := "test-service"
		url := "http://dev.bhinneka.com/test"
		statusEndpoint, err := GetHTTPStatusEndpoint(name, slug, url, httpsURLType, true)
		assert.NoError(t, err)
		assert.Equal(t, name, statusEndpoint.Name)
		assert.Equal(t, slug, statusEndpoint.Slug)
		assert.Equal(t, httpsURLType, statusEndpoint.Type)
		assert.Equal(t, true, statusEndpoint.IsTraversable)
	})
}
