package ally_api

import "encoding/xml"

type Response struct {
	XMLName xml.Name `xml:"response"`
	ResponseId  string  `xml:"id,attr"`
	ElapsedTime float64 `xml:"elapsedtime"`
	Error       string  `xml:"error"`
}
