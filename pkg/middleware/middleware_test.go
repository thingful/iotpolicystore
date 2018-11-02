package middleware_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DECODEproject/iotpolicystore/pkg/middleware"
	"github.com/stretchr/testify/assert"
)

func testHandler() http.HandlerFunc {
	fn := func(rw http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(rw, "ok")
	}
	return http.HandlerFunc(fn)
}

func TestRequestIDMiddleware(t *testing.T) {
	ts := httptest.NewServer(middleware.RequestIDMiddleware(testHandler()))
	defer ts.Close()

	client := ts.Client()

	req, err := http.NewRequest(http.MethodGet, ts.URL, nil)
	assert.Nil(t, err)

	resp, err := client.Do(req)
	assert.Nil(t, err)

	reqID := resp.Header.Get(middleware.RequestIDHeader)
	assert.NotEqual(t, "", reqID)
}

func TestRequestIDMiddlewareWithValue(t *testing.T) {
	ts := httptest.NewServer(middleware.RequestIDMiddleware(testHandler()))
	defer ts.Close()

	client := ts.Client()

	req, err := http.NewRequest(http.MethodGet, ts.URL, nil)
	assert.Nil(t, err)
	req.Header.Set(middleware.RequestIDHeader, "foobar")

	resp, err := client.Do(req)
	assert.Nil(t, err)

	reqID := resp.Header.Get(middleware.RequestIDHeader)
	assert.Equal(t, "foobar", reqID)
}
