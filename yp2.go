package main

import (
	"encoding/xml"
	"net/http"
	"regexp"
	"strings"

	"github.com/labstack/echo"
)

/*
	This is the SHOUTcast yp2 protocol implementation
*/

// YP2Request contains all possible fields in YP2 requests
type YP2Request struct {
	XMLName  xml.Name `xml:"yp"`
	Version  string   `xml:"version,attr"`
	Commands []YP2Cmd `xml:"cmd"`
}

// YP2Cmd contains a given command of a YP2Request
type YP2Cmd struct {
	Operation  string      `xml:"op,attr"`
	Sequence   string      `xml:"seq,attr"`
	Timestamp  string      `xml:"tstamp"`
	URL        YP2URL      `xml:"url"`
	ID         string      `xml:"id"`
	Peak       string      `xml:"peak"`
	MaxClients string      `xml:"maxclients"`
	Authhash   string      `xml:"authhash"`
	Bitrate    string      `xml:"bitrate"`
	Samplerate string      `xml:"samplerate"`
	VBR        string      `xml:"vbr"`
	SA         string      `xml:"sa"`
	PA         string      `xml:"pa"`
	DJ         string      `xml:"dj"`
	Source     string      `xml:"source"`
	DNAS       string      `xml:"dnas"`
	CPU        string      `xml:"cpu"`
	Stats      YP2Stats    `xml:"stats"`
	Metadata   YP2Metadata `xml:"metadata"`
}

// YP2URL is an URL defenition in YP2
type YP2URL struct {
	Port string `xml:"port"`
	Path string `xml:"path"`
	SID  string `xml:"sid"`
	Alt  string `xml:"alt"`
}

// YP2Stats is a statistics defenition in YP2
type YP2Stats struct {
	Listeners            string `xml:"listeners"`
	UniqueListeners      string `xml:"uniquelisteners"`
	AverageListeningTime string `xml:"avglistentime"`
	NewSessions          string `xml:"newsessions"`
	Connects             string `xml:"connects"`
}

// YP2Metadata is a metadata defenition in YP2
type YP2Metadata struct {
	Title string `xml:"title"`
	// probably more un Ultravox mode
}

// YP2Response contains all possible fields in YP2 requests
type YP2Response struct {
	XMLName  xml.Name  `xml:"yp"`
	Version  string    `xml:"version,attr"`
	Response []YP2Resp `xml:"resp,allowempty"`
}

// YP2Resp is the response part inside YP2Response
type YP2Resp struct {
	XMLName         xml.Name    `xml:"resp,allowempty"`
	Sequence        string      `xml:"seq,attr"`
	ID              string      `xml:"id,omitempty"`
	StationID       string      `xml:"stnid,omitempty"`
	UpdateFrequency string      `xml:"updatefreq,omitempty"`
	PublicIP        string      `xml:"publicip,omitempty"`
	Sation          *YP2Station `xml:"station,omitempty"`
	Peak            string      `xml:"peak,omitempty"`
}

// YP2Station is a station defenition in YP2
type YP2Station struct {
	Name   string `xml:"name,omitempty"`
	Genre  string `xml:"genre,omitempty"`
	URL    string `xml:"url,omitempty"`
	AdMode string `xml:"admode,omitempty"`
}

// NewYP2Response gives a new YP2Response with reguired data set
func NewYP2Response() YP2Response {
	return YP2Response{
		Version: "2",
		Response: []YP2Resp{
			YP2Resp{
				Sequence: "1",
			},
		},
	}
}

func yp2Handle(c echo.Context) error {
	fields := YP2Request{}
	c.Bind(&fields)

	if len(fields.Commands) == 0 {
		return sendXML(NewYP2Response(), c)
	}

	if fields.Commands[0].Operation == "add" {
		return handleYP2Add(fields, c)
	}
	if fields.Commands[0].Operation == "version" {
		return handleYP2Version(fields, c)
	}
	if fields.Commands[0].Operation == "update" {
		return handleYP2Update(fields, c)
	}
	if fields.Commands[0].Operation == "remove" {
		return handleYP2Remove(fields, c)
	}

	return c.String(http.StatusOK, "")
}

func handleYP2Add(fields YP2Request, c echo.Context) error {
	resp := NewYP2Response()
	resp.Version = "0" // we do not know why!

	resp.Response[0].ID = "2013197"        // no idea what this is :/
	resp.Response[0].StationID = "1852379" // dummuy data for now
	resp.Response[0].UpdateFrequency = "600"
	resp.Response[0].PublicIP = strings.Split(c.RealIP(), ",")[0]
	resp.Response[0].Sation = &YP2Station{
		Name:   "Discover.fm YP OK",
		Genre:  "Classical",
		URL:    "https://discover.fm",
		AdMode: "0", //hmmm, gives ideas
	}
	resp.Response[0].Peak = "0"

	return sendXML(resp, c)
}

func handleYP2Version(fields YP2Request, c echo.Context) error {
	resp := NewYP2Response()
	// nice but we are not the least interested
	return sendXML(resp, c)
}

func handleYP2Update(fields YP2Request, c echo.Context) error {
	resp := NewYP2Response()

	resp.Response[0].ID = "2013197"        // no idea what this is :/
	resp.Response[0].StationID = "1852379" // dummuy data for now
	resp.Response[0].UpdateFrequency = "600"
	resp.Response[0].PublicIP = strings.Split(c.RealIP(), ",")[0]
	resp.Response[0].Sation = &YP2Station{
		Name:   "Discover.fm YP OK",
		Genre:  "Classical",
		URL:    "https://discover.fm",
		AdMode: "0", //hmmm, gives ideas
	}

	return sendXML(resp, c)
}

func handleYP2Remove(fields YP2Request, c echo.Context) error {
	resp := NewYP2Response()

	return sendXML(resp, c)
}

func sendXML(in interface{}, c echo.Context) error {
	xmlbytes, _ := xml.Marshal(in)
	xmlstring := string(xmlbytes)

	// This is a hack to make Go write self closing tags in order to have the exact same bytes as the official API
	r, err := regexp.Compile(`<([a-zA-Z0-9]*) ([a-zA-Z0-9]*=\"[a-zA-Z0-9]*\")*><(\\|\/)([a-zA-Z0-9]*)>`)
	if err != nil {
		return c.XML(http.StatusOK, in) // fallback to default marshaling
	}
	matches := r.FindAllString(xmlstring, -1)

	if len(matches) > 0 {
		r, err = regexp.Compile("<([a-zA-Z0-9]* ([a-zA-Z0-9]*=\"[a-zA-Z0-9]*\")*)>")
		for i := 0; i < len(matches); i++ {
			xmlTag := r.FindString(matches[i])
			xmlTag = strings.Replace(xmlTag, "<", "", -1)
			xmlTag = strings.Replace(xmlTag, ">", "", -1)
			xmlstring = strings.Replace(xmlstring, matches[i], "<"+xmlTag+" />", -1)

		}
	}

	return c.XMLBlob(http.StatusOK, []byte(xmlstring))
}
