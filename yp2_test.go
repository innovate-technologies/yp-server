package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func Test_yp2_add(t *testing.T) {
	e := echo.New()

	xml := `<?xml version="1.0" encoding="utf-8"?><yp version="2"><cmd op="add" seq="1"><tstamp>1546856225</tstamp><url><port>8000</port><path>/stream</path><sid>1</sid><alt></alt></url><bitrate>128000</bitrate><samplerate>44100</samplerate><vbr>0</vbr><mimetype>audio/mpeg</mimetype><sa>0</sa><pa>0</pa><peak>0</peak><maxclients>9999</maxclients><authhash>test</authhash><dj></dj><source>Legacy / Unknown</source><dnas>2.5.5.733/posix(linux x64)</dnas><cpu>1/1</cpu></cmd></yp>`
	want := `<?xml version="1.0" encoding="UTF-8"?>
<yp version="0"><resp seq="1"><id>2013197</id><stnid>1852379</stnid><updatefreq>600</updatefreq><publicip>167.99.199.90</publicip><station><name>Discover.fm YP OK</name><genre>Classical</genre><url>https://discover.fm</url><admode>0</admode></station><peak>0</peak></resp></yp>`

	req := httptest.NewRequest(http.MethodPost, "/icecast", strings.NewReader(xml))
	req.Header.Set("Content-Type", "text/xml;charset=utf-8")
	req.Header.Set("X-Forwarded-For", "167.99.199.90")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, yp2Handle(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, want, rec.Body.String())
	}
}

func Test_yp2_update(t *testing.T) {
	e := echo.New()

	xml := `<?xml version="1.0" encoding="utf-8"?><yp version="2"><cmd op="update" seq="1"><tstamp>1546856336</tstamp><url><port>8000</port><path>/stream</path><sid>1</sid><alt></alt></url><bitrate>128000</bitrate><samplerate>44100</samplerate><vbr>0</vbr><mimetype>audio/mpeg</mimetype><id>2013197</id><sa>0</sa><pa>0</pa><peak>0</peak><maxclients>9999</maxclients><authhash>GukWIwT3nYePIq9GEMXU</authhash><dj></dj><source>ocaml-cry (Mozilla compatible)</source><dnas>2.5.5.733/posix(linux x64)</dnas><cpu>1/1</cpu><stats><listeners>3</listeners><uniquelisteners>2</uniquelisteners><avglistentime>8</avglistentime><newsessions>0</newsessions><connects>3</connects></stats><metadata><title seq="1">Epic Soul Factory - Extinction</title></metadata></cmd></yp>`
	want := `<?xml version="1.0" encoding="UTF-8"?>
<yp version="2"><resp seq="1"><id>2013197</id><stnid>1852379</stnid><updatefreq>600</updatefreq><publicip>167.99.199.90</publicip><station><name>Discover.fm YP OK</name><genre>Classical</genre><url>https://discover.fm</url><admode>0</admode></station></resp></yp>`

	req := httptest.NewRequest(http.MethodPost, "/icecast", strings.NewReader(xml))
	req.Header.Set("Content-Type", "text/xml;charset=utf-8")
	req.Header.Set("X-Forwarded-For", "167.99.199.90")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, yp2Handle(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, want, rec.Body.String())
	}
}

func Test_yp2_version(t *testing.T) {
	e := echo.New()

	xml := `<?xml version="1.0" encoding="utf-8"?><yp version="2"><cmd op="version" seq="1"><dnas>2.5.5.733/posix(linux x64)</dnas></cmd></yp>`
	want := `<?xml version="1.0" encoding="UTF-8"?>
<yp version="2"><resp seq="1" /></yp>`

	req := httptest.NewRequest(http.MethodPost, "/icecast", strings.NewReader(xml))
	req.Header.Set("Content-Type", "text/xml;charset=utf-8")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, yp2Handle(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, want, rec.Body.String())
	}
}

func Test_yp2_remove(t *testing.T) {
	e := echo.New()

	xml := `<?xml version="1.0" encoding="utf-8"?><yp version="2"><cmd op="remove" seq="1"><tstamp>1546856215</tstamp><url><port>8000</port><path>/stream</path><sid>1</sid><alt></alt></url><id>2013197</id><peak>0</peak><maxclients>9999</maxclients><authhash>test</authhash></cmd></yp>`
	want := `<?xml version="1.0" encoding="UTF-8"?>
<yp version="2"><resp seq="1" /></yp>`

	req := httptest.NewRequest(http.MethodPost, "/icecast", strings.NewReader(xml))
	req.Header.Set("Content-Type", "text/xml;charset=utf-8")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, yp2Handle(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, want, rec.Body.String())
	}
}
