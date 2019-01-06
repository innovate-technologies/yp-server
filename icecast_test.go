package main

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func Test_add_icecast(t *testing.T) {
	e := echo.New()

	vals := url.Values{}
	vals.Set("action", "add")
	vals.Set("sn", "OPENcast")
	vals.Set("type", "audio/mpeg")
	vals.Set("genre", "pop rock")
	vals.Set("b", "128")
	vals.Set("listenurl", "https://opencast.radioca.st/stream/128kbps")

	req := httptest.NewRequest(http.MethodPost, "/icecast", strings.NewReader(vals.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, icecastHandle(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "1", rec.Header().Get("YPResponse"))
		assert.Equal(t, "30", rec.Header().Get("TouchFreq"))
	}
}

func Test_add_icecast_missing_param(t *testing.T) {
	e := echo.New()

	vals := url.Values{}
	vals.Set("action", "add")
	vals.Set("sn", "OPENcast")
	vals.Set("type", "audio/mpeg")
	vals.Set("genre", "pop rock")
	vals.Set("b", "128")
	//vals.Set("listenurl", "https://opencast.radioca.st/stream/128kbps")

	req := httptest.NewRequest(http.MethodPost, "/icecast", strings.NewReader(vals.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, icecastHandle(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "0", rec.Header().Get("YPResponse"))
	}
}

func Test_touch_icecast(t *testing.T) {
	e := echo.New()

	vals := url.Values{}
	vals.Set("action", "touch")
	vals.Set("sid", "OPENcast")

	req := httptest.NewRequest(http.MethodPost, "/icecast", strings.NewReader(vals.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, icecastHandle(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "1", rec.Header().Get("YPResponse"))
	}
}

func Test_touch_icecast_no_sid(t *testing.T) {
	e := echo.New()

	vals := url.Values{}
	vals.Set("action", "touch")
	//vals.Set("sid", "OPENcast")

	req := httptest.NewRequest(http.MethodPost, "/icecast", strings.NewReader(vals.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, icecastHandle(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "0", rec.Header().Get("YPResponse"))
	}
}

func Test_remove_icecast(t *testing.T) {
	e := echo.New()

	vals := url.Values{}
	vals.Set("action", "remove")
	vals.Set("sid", "OPENcast")

	req := httptest.NewRequest(http.MethodPost, "/icecast", strings.NewReader(vals.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, icecastHandle(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "1", rec.Header().Get("YPResponse"))
	}
}

func Test_remove_icecast_no_sid(t *testing.T) {
	e := echo.New()

	vals := url.Values{}
	vals.Set("action", "remove")
	//vals.Set("sid", "OPENcast")

	req := httptest.NewRequest(http.MethodPost, "/icecast", strings.NewReader(vals.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, icecastHandle(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "0", rec.Header().Get("YPResponse"))
	}
}
