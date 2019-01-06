package main

import (
	"net/http"

	"github.com/labstack/echo"
)

/*
	This is the Icecast yp v1 protocol implementation
	As documented under legacy in https://wiki.xiph.org/Icecast_Server/YP-protocol-v2
	the v2 protocol is not yet published
*/

// IcecastYPFields contains all possible fields in Icecast YP calls
type IcecastYPFields struct {
	Action               string `form:"action"`        // The YP protocol request type
	StreamName           string `form:"sn"`            // The name of the stream
	ContentType          string `form:"type"`          // content type
	Genre                string `form:"genre"`         // genre, space delimited
	Bitrate              string `form:"b"`             // The expected average bitrate for the stream
	ListenURL            string `form:"listenurl"`     // The URL of the actual stream, as used by player clients
	Description          string `form:"desc"`          // Server Description
	URL                  string `form:"url"`           // Stream URL
	SessionID            string `form:"sid"`           // Session ID
	SongTitle            string `form:"st"`            // Song Title
	listeners            string `form:"st"`            // Current number of listeners
	MaxListeners         string `form:"max_listeners"` // max listener limit for this stream
	AverageListeningtime string `form:"alt"`           // average listening time
	Hits                 string `form:"ht"`            // hits / tune ins
	AverageTuneIns       string `form:"cm"`            // 5min average tune ins
}

func icecastHandle(c echo.Context) error {
	fields := IcecastYPFields{}
	c.Bind(&fields)

	if fields.Action == "add" {
		return icecastAdd(fields, c)
	}

	if fields.Action == "touch" {
		return icecastTouch(fields, c)
	}

	if fields.Action == "remove" {
		return icecastRemove(fields, c)
	}

	return c.String(http.StatusOK, "")
}

func icecastAdd(fields IcecastYPFields, c echo.Context) error {
	// TODO: handle input

	if fields.StreamName == "" {
		return icecastFail("Stream name missing", c)
	}
	if fields.ListenURL == "" {
		return icecastFail("Listen URL name missing", c)
	}
	if fields.Genre == "" {
		return icecastFail("Genre missing", c)
	}

	c.Response().Header().Set("YPResponse", "1")
	c.Response().Header().Set("YPMessage", "")
	c.Response().Header().Set("SID", "TODO")
	c.Response().Header().Set("TouchFreq", "30")
	return c.String(http.StatusOK, "")
}

func icecastTouch(fields IcecastYPFields, c echo.Context) error {
	// TODO: handle input

	if fields.SessionID == "" {
		return icecastFail("SID missing", c)
	}

	c.Response().Header().Set("YPResponse", "1")
	c.Response().Header().Set("YPMessage", "")
	return c.String(http.StatusOK, "")
}

func icecastRemove(fields IcecastYPFields, c echo.Context) error {
	// TODO: handle input

	if fields.SessionID == "" {
		return icecastFail("SID missing", c)
	}

	c.Response().Header().Set("YPResponse", "1")
	c.Response().Header().Set("YPMessage", "")
	return c.String(http.StatusOK, "")
}

func icecastFail(err string, c echo.Context) error {
	c.Response().Header().Set("YPResponse", "0")
	c.Response().Header().Set("YPMessage", err)
	return c.String(http.StatusOK, "") // i know
}
